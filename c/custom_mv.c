#include <stdlib.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
  if (argc != 3) {
    fprintf(stderr, "%s: wrong argument\n", argv[0]);
    exit(1);
  }

  if (rename(argv[1], argv[2]) < 0) {
    perror(argv[1]);
    exit(1);
  }

  return 0;
}

