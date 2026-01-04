// cat コマンドを実装する
// - 引数で渡されたファイルを標準出力に出力する
// - 引数がなければ標準入力を標準出力にに出力する

#include <fcntl.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/_types/_ssize_t.h>
#include <sys/errno.h>
#include <sys/fcntl.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>

#define BUF_SIZE 1024

void *xmalloc(size_t size);
ssize_t write_all(int fd, const void *buf, size_t count);

int main(int argc, char *argv[]) {
  // 標準入力
  if (argc == 1) {
    char *buf = xmalloc(BUF_SIZE);
    ssize_t n;
    while ((n = read(STDIN_FILENO, buf, BUF_SIZE)) > 0) {
      if (write_all(STDOUT_FILENO, buf, n) != n) {
        perror("write");
        exit(EXIT_FAILURE);
      }
    }
    if (n < 0) {
      perror("read");
      exit(EXIT_FAILURE);
    }

    free(buf);
    exit(EXIT_SUCCESS);
  }

  // ファイル入力
  for (int i = 1; i < argc; i++) {
    int fd = open(argv[i], O_RDONLY);
    printf("filename: %s, fd: %d\n", argv[i], fd);

    char *buf = xmalloc(BUF_SIZE);

    ssize_t n;
    while ((n = read(fd, buf, BUF_SIZE)) > 0) {
      if (write_all(STDOUT_FILENO, buf, n) != n) {
        perror("write");
        exit(EXIT_FAILURE);
      }
    }

    if (n < 0) {
      perror("read");
      exit(EXIT_FAILURE);
    }

    free(buf);
    if (close(fd) < 0) {
      perror("close");
      exit(EXIT_FAILURE);
    }
  }

  return EXIT_SUCCESS;
}

void *xmalloc(size_t size) {
  if (size == 0)
    size = 1;
  void *p = malloc(size);
  if (!p) {
    perror("malloc");
    abort();
  }

  return p;
}

ssize_t write_all(int fd, const void *buf, size_t count) {
  size_t written = 0;

  while (written < count) {
    ssize_t n = write(fd, buf, count);
    if (n < 0) {
      if (errno == EINTR)
        continue;
      return -1;
    }
    written += n;
  }
  return written;
}
