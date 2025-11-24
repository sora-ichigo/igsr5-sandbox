#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char *argv[])
{
  if (argc != 3) {
    fprintf(stderr, "%s: no arguments\n", argv[0]);
    exit(0);
  }

  if (symlink(argv[1], argv[2]) < 0) {
    perror(argv[1]);
    exit(1);
  }

  return 0;
}
