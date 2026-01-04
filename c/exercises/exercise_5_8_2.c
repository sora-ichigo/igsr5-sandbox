// wc -l を実装する

#include <stdio.h>
#include <stdlib.h>
#include <sys/_types/_ssize_t.h>

#define BUFSIZE 1024

int main(int argc, char *argv[]) {
  char buf[BUFSIZE];

  // 標準入力
  if (argc == 1) {
    ssize_t n;
    int count = 0;
    while ((n = fread(buf, 1, BUFSIZE, stdin)) > 0) {
      int seen = 0;
      while (seen < n) {
        if (buf[seen] == '\n')
          count++;
        seen++;
      }
    }

    if (ferror(stdin)) {
      perror("fread");
      exit(EXIT_FAILURE);
    }

    if (feof(stdin)) {
      printf("eof\n");
      exit(EXIT_SUCCESS);
    }

    printf("%d\n", count);
    exit(EXIT_SUCCESS);
  }

  // ファイル入力
  for (int i = 1; i < argc; i++) {
    char *filename = argv[i];
    FILE *f = fopen(filename, "r");
    if (!f) {
      perror("fopen");
      exit(EXIT_FAILURE);
    }

    ssize_t n;
    int count = 0;
    while ((n = fread(buf, 1, BUFSIZE, f)) > 0) {
      int seen = 0;
      while (seen < n) {
        if (buf[seen] == '\n')
          count++;
        seen++;
      }
    }

    printf("%d %s\n", count, filename);

    fclose(f);
  }

  return EXIT_SUCCESS;
}
