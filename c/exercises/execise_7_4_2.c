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

static void *xmalloc(size_t size) {
  void *p = malloc(size);
  if (!p) {
    perror("malloc");
    exit(EXIT_FAILURE);
  }
  return p;
}

static char *line_read(FILE *f) {
  char *line = malloc(CHUNK_SIZE);
  if (fgets(line, CHUNK_SIZE, f) == NULL) {
    if (feof(f)) {
      free(line);
      return NULL;
    }
    perror("fgets");
    exit(EXIT_FAILURE);
  }
  return line;
}

static void ring_push(char **ring, int ring_index, char *line) {
  char *line_copy = xmalloc(strlen(line) + 1);
  strcpy(line_copy, line);
  if (ring[ring_index])
    free(ring[ring_index]);
  ring[ring_index] = line_copy;
}

static void ring_print(char **ring, int line_read) {
  for (int i = 0; i < TAIL_N; i++) {
    int print_pos = (line_read + i) % TAIL_N;
    if (ring[print_pos])
      fputs(ring[print_pos], stdout);
  }
}

static void ring_free(char **ring) {
  for (int i = 0; i < TAIL_N; i++) {
    if (ring[i])
      free(ring[i]);
  }
}

int main(int argc, char *argv[]) {
  for (int i = 1; i < argc; i++) {
    FILE *f = xfopen(argv[i], "rb");

    char *ring[TAIL_N] = {0};
    int line_read_count = 0;
    int ring_index = 0;

    for (;;) {
      char *line = line_read(f);
      if (feof(f)) {
        break;
      }

      ring_index = line_read_count % TAIL_N;
      ring_push(ring, ring_index, line);

      free(line);
      line_read_count++;
    }

    ring_print(ring, line_read_count);
    ring_free(ring);

    xfclose(f);
  }

  return EXIT_SUCCESS;
}
