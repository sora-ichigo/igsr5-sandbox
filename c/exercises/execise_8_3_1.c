// grepコマンドの -i オプションと -v オプションに対応したバージョンを作成せよ。
// 組み合わせてつか用できるようにせよ。

#include <regex.h>
#include <stdio.h>
#include <stdlib.h>

void die(const char *msg) {
  perror(msg);
  exit(EXIT_FAILURE);
}

FILE *fopen_or_exit(const char *filename, const char *mode) {
  FILE *f = fopen(filename, mode);
  if (f == NULL)
    die("fopen");
  return f;
}

void fclose_or_exit(FILE *f) {
  if (fclose(f) != 0)
    die("fclose");
}

void fputs_or_exit(const char *s, FILE *stream) {
  if (fputs(s, stream) == EOF)
    die("fputs");
}

regex_t compile_regex_or_exit(const char *pattern, int cflags) {
  regex_t regex;
  if (regcomp(&regex, pattern, cflags) != 0)
    die("regcomp");
  return regex;
}

void do_grep(const char *filename, const char *pattern) {
  FILE *f = fopen_or_exit(filename, "r");

  regex_t regex = compile_regex_or_exit(pattern, REG_EXTENDED | REG_ICASE);
  regmatch_t match[1];

  while (1) {
    char *line = NULL;
    size_t linecap = 0;

    if (getline(&line, &linecap, f) == -1) {
      if (feof(f)) {
        break;
      }
      free(line);
      die("getline");
    }

    if (regexec(&regex, line, 1, match, 0) == 0) {
      fputs_or_exit(line, stdout);
    }

    free(line);
  }

  fclose_or_exit(f);

  return;
}

int main(int argc, char *argv[]) {
  if (argc != 3) {
    fprintf(stderr, "Usage: %s <arg1> <arg2>\n", argv[0]);
    return EXIT_FAILURE;
  }

  do_grep(argv[1], argv[2]);

  return EXIT_SUCCESS;
}
