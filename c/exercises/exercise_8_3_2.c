// マッチした文字列のみ出力するgrepコマンドを実装せよ。

#include <_regex.h>
#include <regex.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

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

void do_grep(const char *filename, const char *pattern, bool match_icase,
             bool out_invert) {
  FILE *f = fopen_or_exit(filename, "r");

  int cflags = REG_EXTENDED;
  if (match_icase) {
    cflags |= REG_ICASE;
  }
  regex_t regex = compile_regex_or_exit(pattern, cflags);
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

    bool matched = (regexec(&regex, line, 1, match, 0) == 0);
    if (matched != out_invert) {
      regoff_t len = match[0].rm_eo - match[0].rm_so;
      fwrite(line + match[0].rm_so, 1, len, stdout);
      fputc('\n', stdout);
    }

    free(line);
  }

  regfree(&regex);
  fclose_or_exit(f);

  return;
}

int main(int argc, char *argv[]) {
  if (argc < 3) {
    fprintf(stderr, "Usage: %s <arg1> <arg2>\n", argv[0]);
    return EXIT_FAILURE;
  }

  int opt;
  bool match_icase = false;
  bool out_invert = false;
  while ((opt = getopt(argc, argv, "iv")) != -1) {
    switch (opt) {
    case 'i':
      // -i オプションの処理（大文字小文字を区別しない）
      match_icase = true;
      break;
    case 'v':
      // -v オプションの処理（マッチしない行を表示）
      out_invert = true;
      break;
    default:
      fprintf(stderr, "Usage: %s [-i] [-v] <arg1> <arg2>\n", argv[0]);
      return EXIT_FAILURE;
    }
  }

  do_grep(argv[optind], argv[optind + 1], match_icase, out_invert);

  return EXIT_SUCCESS;
}
