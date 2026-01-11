// n=5固定のtailコマンドを実装しなさい
// この実装はファイル先頭から read_line を実行するため、
// 行数nのファイルに対して必ず全行を走査し、時間計算量は O(n) となる

#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define CHUNK_SIZE 1024
#define TAIL_N 5

static FILE *xfopen(const char *filename, const char *mode) {
  FILE *f = fopen(filename, mode);
  if (!f) {
    perror("fopen");
    exit(EXIT_FAILURE);
  }
  return f;
}

static void xfclose(FILE *f) {
  if (fclose(f) != 0) {
    perror("fclose");
    exit(EXIT_FAILURE);
  }
}

int main(int argc, char *argv[]) {
  for (int i = 1; i < argc; i++) {
    FILE *f = xfopen(argv[i], "rb");

    char *ring[TAIL_N] = {0};
    int idx = 0;
    int pos = 0;

    for (;;) {
      char *ret;
      char line[CHUNK_SIZE];
      ret = fgets(line, sizeof(line), f);
      if (ret == NULL) {
        if (feof(f))
          break;
        perror("fgets");
        exit(EXIT_FAILURE);
      }

      pos = idx % TAIL_N;
      char *p = malloc(strlen(line) + 1);
      if (ring[pos])
        free(ring[pos]);
      strcpy(p, line);
      ring[pos] = p;
      idx++;
    }

    int print_pos = idx % TAIL_N;
    for (int j = 0; j < TAIL_N; j++) {
      if (ring[print_pos]) {
        fputs(ring[print_pos], stdout);
        free(ring[print_pos]);
      }
      print_pos = (print_pos + 1) % TAIL_N;
    }

    xfclose(f);
  }

  return EXIT_SUCCESS;
}
