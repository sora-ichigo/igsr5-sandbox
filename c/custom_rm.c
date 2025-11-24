#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

int main(int argc, char *argv[])
{
  if (argc < 2) {
    fprintf(stderr, "%s: no argument\n", argv[0]);
    exit(1);
  }

  if (unlink(argv[1]) < 0) {
    perror(argv[1]);
    exit(1);
  }

  return 0;
}
