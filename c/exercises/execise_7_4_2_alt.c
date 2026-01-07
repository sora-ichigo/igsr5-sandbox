#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>

#define TAIL_N 5
#define CHUNK_SIZE 4096

static void xfseek(FILE *f, long offset, int whence) {
  if (fseek(f, offset, whence) != 0) {
    perror("fseek");
    exit(EXIT_FAILURE);
  }
}

static long xftell(FILE *f) {
  long size = ftell(f);
  if (size < 0) {
    perror("ftell");
    exit(EXIT_FAILURE);
  }
  return size;
}

static size_t xfread(void *buf, size_t size, size_t n, FILE *f) {
  size_t ret = fread(buf, size, n, f);
  if (ferror(f)) {
    perror("fread");
    exit(EXIT_FAILURE);
  }
  return ret;
}

static size_t xfwrite(const void *buf, size_t size, size_t n, FILE *f) {
  size_t ret = fwrite(buf, size, n, f);
  if (ret < n) {
    perror("fwrite");
    exit(EXIT_FAILURE);
  }
  return ret;
}

static long find_tail_start(FILE *f, int n) {
  xfseek(f, 0, SEEK_END);
  long pos = xftell(f);

  int newlines_found = 0;
  char buf[CHUNK_SIZE];

  while (pos > 0 && newlines_found <= n) {
    long read_size = (pos < CHUNK_SIZE) ? pos : CHUNK_SIZE;
    pos -= read_size;
    xfseek(f, pos, SEEK_SET);
    xfread(buf, 1, read_size, f);

    for (long i = read_size - 1; i >= 0; i--) {
      if (buf[i] == '\n')
        newlines_found++;
      if (newlines_found > n) {
        pos += i + 1; // この時点で末尾からn個目の\nがどの位置にあるのか分かる
        break;
      }
    }
  }
  return pos;
}

static void write_from_position(FILE *f, long start) {
  xfseek(f, start, SEEK_SET);

  char buf[CHUNK_SIZE];
  size_t nread;
  while ((nread = xfread(buf, 1, CHUNK_SIZE, f)) > 0) {
    xfwrite(buf, 1, nread, stdout);
  }
}

int main(int argc, char *argv[]) {
  if (argc < 2) { // 標準入力
    // TODO: 後で実装する
  } else { // ファイル名
    for (int i = 1; i < argc; i++) {
      FILE *f = fopen(argv[i], "rb");
      if (!f) {
        perror("fopen");
        exit(EXIT_FAILURE);
      }

      long start = find_tail_start(f, TAIL_N);
      write_from_position(f, start);

      if (fclose(f) != 0) {
        perror("fclose");
        exit(EXIT_FAILURE);
      }
    }
  }
  return EXIT_SUCCESS;
}
