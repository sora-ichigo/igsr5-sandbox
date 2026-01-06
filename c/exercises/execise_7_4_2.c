// n=5固定のtailコマンドを実装しなさい
// この実装はファイル先頭から read_line を実行するため、
// 行数nのファイルに対して必ず全行を走査し、時間計算量は O(n) となる

#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

enum { TAIL_N = 5, CHUNK = 1024 };

static char *read_line(FILE *f) {
  char buf[CHUNK];

  char *line = NULL;
  size_t cap = 0;
  size_t len = 0;

  while (fgets(buf, sizeof(buf), f) != NULL) {
    size_t blen = strlen(buf);
    size_t new_cap = cap ? cap * 2 : 128;
    while (new_cap < len + blen + 1)
      new_cap *= 2;
    char *p = realloc(line, new_cap);
    if (!p) {
      free(line);
      perror("realloc");
      exit(EXIT_FAILURE);
    }
    line = p;
    cap = new_cap;

    memcpy(line + len, buf, blen);
    len += blen;
    line[len] = '\0';

    if (blen > 0 && buf[blen - 1] == '\n') {
      return line; // 行末まで読めた
    }
    // 改行が来ていない = 長い行なので続けて読む
  }

  if (ferror(f)) {
    perror("fgets");
    free(line);
    exit(EXIT_FAILURE);
  }

  // EOF
  if (len == 0) {
    free(line);
    return NULL; // 何も読めていない
  }

  return line;
}

int main(int argc, char *argv[]) {
  for (int i = 1; i < argc; i++) {
    FILE *f = fopen(argv[i], "rb");
    if (!f) {
      perror("fopen");
      exit(EXIT_FAILURE);
    }

    char *ring[TAIL_N] = {0};
    size_t count = 0;

    for (;;) {
      char *line = read_line(f);
      if (!line)
        break;

      size_t idx = count % TAIL_N;
      free(ring[idx]);
      ring[idx] = line;
      count++;
    }

    size_t out = (count < TAIL_N) ? count : TAIL_N;
    size_t start = (count <= TAIL_N) ? 0 : (count % TAIL_N);

    for (size_t k = 0; k < out; k++) {
      size_t idx = (start + k) % TAIL_N;
      fputs(ring[idx], stdout);
    }

    for (int k = 0; k < TAIL_N; k++)
      free(ring[k]);

    if (fclose(f) != 0) {
      perror("fclose");
      exit(EXIT_FAILURE);
    }
  }
  return EXIT_SUCCESS;
}
