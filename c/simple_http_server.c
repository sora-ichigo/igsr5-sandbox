#include <arpa/inet.h>
#include <ctype.h>
#include <errno.h>
#include <netinet/in.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/_endian.h>
#include <sys/_types/_socklen_t.h>
#include <sys/_types/_ssize_t.h>
#include <sys/socket.h>
#include <unistd.h>

#define PORT 8081
#define BUFFER_SIZE 1024

#define MAX_HEADER_BYTES (64 * 1024)
#define MAX_LINE_BYTES (8 * 1024)
#define READ_CHUNK 4096

typedef struct {
  char *raw;
  size_t raw_len;

  char *method;
  char *path;
  char *version;

  char **header_keys;
  char **header_vals;
  size_t header_count;
} http_request;

static void die(const char *msg) {
  perror(msg);
  exit(EXIT_FAILURE);
}

static void *xmalloc(size_t n) {
  void *p = malloc(n);
  if (!p)
    die("malloc");
  return p;
}

static void *xrealloc(void *p, size_t n) {
  void *q = realloc(p, n);
  if (!q)
    die("realloc");
  return q;
}

static int send_all(int fd, const void *buf, size_t len) {
  const char *p = (const char *)buf;
  size_t sent = 0;
  while (sent < len) {
    ssize_t n = write(fd, p + sent, len - sent);
    if (n < 0) {
      if (errno == EINTR)
        continue;
      return -1;
    }
    if (n == 0)
      return -1;
    sent += (size_t)n;
  }
  return 0;
}

// \r\n\r\n を探す (見つかったら raw のヘッダ終端位置を返す)
static ssize_t find_header_end(const char *buf, size_t len) {
  if (len < 4)
    return -1;
  for (size_t i = 0; i + 3 < len; i++) {
    if (buf[i] == '\r' && buf[i + 1] == '\n' && buf[i + 2] == '\r' &&
        buf[i + 3] == '\n') {
      return (ssize_t)(i + 4);
    }
  }
  return -1;
}

// ヘッダ終端まで読みこむ (ボディは読み込まない)
static int read_headers(int fd, char **out, size_t *out_len) {
  size_t cap = READ_CHUNK;
  size_t len = 0;
  char *buf = (char *)xmalloc(cap);

  while (1) {
    if (len >= MAX_HEADER_BYTES) {
      free(buf);
      return 413; // Payload Too Large
    }

    if (cap - len < READ_CHUNK) {
      cap = cap * 2;
      if (cap > MAX_HEADER_BYTES + 1)
        cap = MAX_HEADER_BYTES + 1;
      buf = (char *)xrealloc(buf, cap);
    }

    ssize_t n = read(fd, buf + len, cap - len);
    if (n < 0) {
      if (errno == EINTR)
        continue;
      free(buf);
      return 400;
    }
    if (n == 0) {
      free(buf);
      return 400;
    }
    len += (size_t)n;

    ssize_t end = find_header_end(buf, len);
    if (end >= 0) {
      size_t header_len = (size_t)end;
      if (header_len > MAX_HEADER_BYTES) {
        free(buf);
        return 413;
      }

      // raw を header_len ピッタリ + NUL に詰める
      char *raw = (char *)xmalloc(header_len + 1);
      memcpy(raw, buf, header_len);
      raw[header_len] = '\0';
      free(buf);

      *out = raw;
      *out_len = header_len;
      return 0;
    }
  }
}

static char *trim(char *s) {
  while (*s && (*s == ' ' || *s == '\t'))
    s++;
  char *end = s + strlen(s);
  while (end > s && (end[-1] == ' ' || end[-1] == '\t'))
    end--;
  *end = '\0';
  return s;
}

static int validate_line_len(const char *line) {
  size_t n = strlen(line);
  return (n <= MAX_LINE_BYTES) ? 0 : -1;
}

// p から end までの範囲で NUL をスキップし、次の非 NUL 位置を返す。なければ end
static char *skip_nuls(char *p, char *end) {
  while (p < end && *p == '\0')
    p++;
  return p;
}

// p から end までの範囲で現在行の末尾 (次の NUL 位置)
// まで進める。返り値は次位置 (末尾NULの次)
static char *next_after_cstring(char *p, char *end) {
  while (p < end && *p != '\0')
    p++;
  if (p < end)
    p++;
  return p;
}

// リクエストライン + ヘッダ行をパース
static int parse_request(http_request *req) {
  if (!req->raw || req->raw_len == 0)
    return 0;

  char *start = req->raw;
  char *end = req->raw + req->raw_len;

  // \r \n を \0 につぶす
  for (char *p = start; p < end; p++) {
    if (*p == '\r' || *p == '\n')
      *p = '\0';
  }

  // 1行目 (リクエストライン)
  char *p = skip_nuls(start, end);
  if (p >= end || *p == '\0')
    return 400;
  if (validate_line_len(p) < 0)
    return 414;

  char *line1 = p;

  // method path version を SP で分割
  char *sp1 = strchr(line1, ' ');
  if (!sp1)
    return 400;
  *sp1 = '\0';

  char *sp2 = strchr(sp1 + 1, ' ');
  if (!sp2)
    return 400;
  *sp2 = '\0';

  req->method = line1;
  req->path = sp1 + 1;
  req->version = sp2 + 1;

  if (*req->method == '\0' || *req->path == '\0' || *req->version == '\0')
    return 400;
  if (strncmp(req->version, "HTTP/1.1", 8) != 0 &&
      strncmp(req->version, "HTTP/1.0", 8) != 0)
    return 400;

  // 次の行へ
  // NOTE: ここがバグっていて、ヘッダ行じゃなくて "/" になってしまっている
  p = next_after_cstring(req->version, end);
  p = skip_nuls(p, end);

  // ヘッダ配列確保 (適当)
  size_t max_headers = req->raw_len / 2 + 4;
  req->header_keys = (char **)xmalloc(sizeof(char *) * max_headers);
  req->header_vals = (char **)xmalloc(sizeof(char *) * max_headers);
  req->header_count = 0;

  // ヘッダ行を列挙
  while (1) {
    p = skip_nuls(p, end);
    if (p >= end || *p == '\0')
      break;

    if (validate_line_len(p) < 0)
      return 431; // Request Header Field Too Long

    char *colon = strchr(p, ':');
    if (!colon) {
      printf("colon is not found\n");
      return 400;
    }

    *colon = '\0';
    char *key = trim(p);
    char *val = trim(colon + 1);
    printf("key: %s\n", key);
    printf("val: %s\n", val);

    if (*key == '\0') {
      printf("key is not found\n");
      return 400;
    }

    req->header_keys[req->header_count] = key;
    req->header_vals[req->header_count] = val;
    req->header_count++;

    if (req->header_count >= max_headers)
      return 431;

    p = next_after_cstring(val, end);
  }

  return 0;
}

static const char *get_header(const http_request *req, const char *key) {
  for (size_t i = 0; i < req->header_count; i++) {
    const char *k = req->header_keys[i];
    size_t a = strlen(k), b = strlen(key);
    if (a != b)
      continue;

    int eq = 1;
    for (size_t j = 0; j < a; j++) {
      if (tolower((unsigned char)k[j]) != tolower((unsigned char)key[j])) {
        eq = 0;
        break;
      }
    }
    if (eq)
      return req->header_vals[i];
  }
  return NULL;
}

static void free_request(http_request *req) {
  free(req->header_keys);
  free(req->header_vals);
  free(req->raw);
  memset(req, 0, sizeof(*req));
}

static int respond_text(int fd, int status, const char *status_text,
                        const char *body) {
  if (!body)
    body = "";
  int body_len = (int)strlen(body);

  char header[512];
  int n = snprintf(header, sizeof(header),
                   "HTTP/1.1 %d %s\r\n"
                   "Content-Type: text/plain; charset=utf-8\r\n"
                   "Content-Length: %d\r\n"
                   "Connection: close\r\n"
                   "\r\n",
                   status, status_text, body_len);
  if (n < 0 || n >= (int)sizeof(header))
    return -1;

  if (send_all(fd, header, (size_t)n) < 0)
    return -1;
  if (send_all(fd, body, (size_t)body_len) < 0)
    return -1;
  return 0;
}

int main(int argc, char *argv[]) {
  int server_fd = socket(AF_INET, SOCK_STREAM, 0);
  if (server_fd < 0)
    die("socket");

  int opt = 1;
  if (setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt)) < 0)
    die("setsockopt");

  struct sockaddr_in addr;
  memset(&addr, 0, sizeof(addr));
  addr.sin_family = AF_INET;
  addr.sin_addr.s_addr = INADDR_ANY;
  addr.sin_port = htons(PORT);

  if (bind(server_fd, (struct sockaddr *)&addr, sizeof(addr)) < 0)
    die("bind");
  if (listen(server_fd, 64) < 0)
    die("listen");

  printf("Listening on https://localhost:%d\n", PORT);

  while (1) {
    struct sockaddr_in client_addr;
    socklen_t client_len = sizeof(client_addr);
    int client_fd =
        accept(server_fd, (struct sockaddr *)&client_addr, &client_len);
    if (client_fd < 0) {
      perror("accept");
      continue;
    }

    http_request req;
    memset(&req, 0, sizeof(req));

    char *raw = NULL;
    size_t raw_len = 0;
    int rc = read_headers(client_fd, &raw, &raw_len);
    if (rc != 0) {
      (void)respond_text(client_fd, rc, "Bad Request", "bad request\n");
      close(client_fd);
      continue;
    }

    req.raw = raw;
    req.raw_len = raw_len;

    rc = parse_request(&req);
    if (rc != 0) {
      (void)respond_text(client_fd, rc, "Bad Request", "bad request\n");
      free_request(&req);
      close(client_fd);
      continue;
    }

    // GET
    if (strcmp(req.method, "GET") != 0) {
      (void)respond_text(client_fd, 405, "Method Not Allowed",
                         "only GET is supported\n");
      free_request(&req);
      close(client_fd);
      continue;
    }

    // HTTP/1.1 は Host 必須
    if (strncmp(req.version, "HTTP/1.1", 8) == 0) {
      if (!get_header(&req, "Host")) {
        (void)respond_text(client_fd, 400, "Bad Request",
                           "Host header is required in HTTP/1.1\n");
        free_request(&req);
        close(client_fd);
        continue;
      }
    }

    // ルーティング
    if (strcmp(req.path, "/") == 0) {
      (void)respond_text(client_fd, 200, "OK", "hello from C server\n");
    } else if (strcmp(req.path, "/health") == 0) {
      (void)respond_text(client_fd, 200, "OK", "ok\n");
    } else if (strcmp(req.path, "/echo") == 0) {
      char body[512];
      snprintf(body, sizeof(body), "method=%s\npath=%s\nversion=%s\n",
               req.method, req.path, req.version);
      (void)respond_text(client_fd, 200, "OK", body);
    } else {
      (void)respond_text(client_fd, 404, "Not Found", "not found\n");
    }

    free_request(&req);
    close(client_fd);
  }

  close(server_fd);
  return EXIT_SUCCESS;
}
