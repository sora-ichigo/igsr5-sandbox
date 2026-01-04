// cat コマンドを実装する
// - 引数で渡されたファイルを標準出力に出力する
// - 引数がなければ標準入力を標準出力にに出力する

#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>

#define BUFSIZE 1024

int main(int argc, char *argv[]) {
  for (int i = 1; i < argc; i++) {
    FILE *f = fopen(argv[i], "r");
    if (!f) {
      perror("fopen");
      exit(EXIT_FAILURE);
    }

    char buf[BUFSIZE];

    size_t n;
    while ((n = fread(buf, 1, BUFSIZE, f)) > 0) {
      size_t written = 0;

      while (written < n) {
        size_t w = fwrite(buf + written, 1, n - written, stdout);

        if (w == 0) {
          perror("fwrite");
          exit(EXIT_FAILURE);
        }
        written += w;
      }
    }

    if (ferror(f)) {
      perror("fread");
      exit(EXIT_FAILURE);
    }

    fclose(f);
  }

  return EXIT_SUCCESS;
}
