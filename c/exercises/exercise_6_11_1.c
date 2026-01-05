// タブ文字を "\t" に、改行文字を "$" + 改行文字に変換する cat
// コマンドを実装しなさい

#include <errno.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>

static void outc(int c) {
  if (putchar(c) == EOF) {
    perror("putchar");
    exit(EXIT_FAILURE);
  }
}

static void transform(FILE *in) {
  for (;;) {
    int c = fgetc(in);
    if (c == EOF) {
      if (ferror(in)) {
        perror("fgetc");
        exit(EXIT_FAILURE);
      }
      return; // EOF
    }

    if (c == '\n') {
      outc('$');
      outc('\n');
    } else if (c == '\t') {
      outc('\\');
      outc('t');
    } else {
      outc(c);
    }
  }
}

int main(int argc, char *argv[]) {
  if (argc == 1) {
    transform(stdin);
    if (fflush(stdout) == EOF) {
      perror("fflush");
      exit(EXIT_FAILURE);
    }
  }

  int status = EXIT_SUCCESS;

  for (int ai = 1; ai < argc; ai++) {
    FILE *f = fopen(argv[ai], "rb");
    if (!f) {
      fprintf(stderr, "%s: %s\n", argv[ai], strerror(errno));
      status = EXIT_FAILURE;
      continue;
    }
    transform(f);
    if (fclose(f) != 0) {
      fprintf(stderr, "%s: %s\n", argv[ai], strerror(errno));
      status = EXIT_FAILURE;
      continue;
    }
  }

  if (fflush(stdout) == EOF)
    status = EXIT_FAILURE;

  return status;
}
