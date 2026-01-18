#include <regex.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
  if (argc != 3) {
    fprintf(stderr, "Usage: %s <arg1> <arg2>\n", argv[0]);
    return EXIT_FAILURE;
  }

  char *filename = argv[1];
  char *pattern = argv[2];

  FILE *f = fopen(filename, "r");
  if (!f) {
    perror("fopen");
    exit(EXIT_FAILURE);
  }
  regex_t regex;
  if (regcomp(&regex, pattern, REG_EXTENDED) != 0) {
    perror("regcomp");
    exit(EXIT_FAILURE);
  }
  regmatch_t match[1];

  while (1) {
    char *line = NULL;
    size_t linecap = 0;

    if (getline(&line, &linecap, f) == -1) {
      if (feof(f)) {
        break;
      }
      perror("getline");
      free(line);
      exit(EXIT_FAILURE);
    }

    if (regexec(&regex, line, 1, match, 0) == 0) {
      fputs(line, stdout);
    }

    free(line);
  }

  if (fclose(f) != 0) {
    perror("fclose");
    exit(EXIT_FAILURE);
  }

  return EXIT_SUCCESS;
}
