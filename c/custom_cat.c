#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>

static void do_cat(const char *path);
static void die(const char *s);

int main(int argc, char *argv[])
{
  if (argc < 2) {
    fprintf(stderr, "%s: file name not given\n", argv[0]);
    exit(1);
  }

  int i;
  for (i = 0; i < argc; i++) {
    do_cat(argv[i]);
  }

  return 0;
}

#define BUFFER_SIZE 2048

static void do_cat(const char *path) {
  int fd;
  unsigned char buf[BUFFER_SIZE];
  int n;
  int newline_count = 0;
  printf("%s\n", path);

  fd = open(path, O_RDONLY);
  if (fd < 0) die(path);

  for (;;) {
    n = read(fd, buf, sizeof buf);
    if (n < 0) die(path);
    if (n == 0) break;

    int i;
    for (i = 0; i < n; i++) {
      if (buf[i] == '\n') {
        newline_count++;
      }
    }
  }

  printf("%d\n", newline_count);

  if (close(fd) < 0) die(path);
}

static void die(const char *s) {
  perror(s);
  exit(1);
}

