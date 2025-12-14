#include <stdio.h>
#include <stdlib.h>

extern char **environ;

int main(int argc, char *argv[]) {
  char **p;

  for (p = environ; *p; p++) {
    printf("%s\n", *p);
  }

  char *value = getenv("USER");
  printf("USER: %s\n", value);

  if (setenv("USER", "hoge", 1) != 0) {
    perror("setenv");
    exit(EXIT_FAILURE);
  }

  value = getenv("USER");
  printf("USER: %s\n", value);

  return EXIT_SUCCESS;
}
