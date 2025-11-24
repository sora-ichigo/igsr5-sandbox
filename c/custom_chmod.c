#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>

int main(int argc, char *argv[]) {
  if (argc < 2) {
    fprintf(stderr, "%s: wrong arguments\n", argv[0]);
    exit(0);
  }

  int mode;
  mode = strtol(argv[1], NULL, 8);
  int i;
  for (i = 2; i < argc; i++) {
    if (chmod(argv[i], mode) < 0) {
      perror(argv[i]);
    }
  }

  return 0;
}
