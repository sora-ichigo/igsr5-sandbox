#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>

int main(int argc, char *argv[])
{
  if (argc != 2) {
    fprintf(stderr, "%s: wrong arguments\n", argv[0]);
    exit(1);
  }

  struct stat st;

  if (lstat(argv[1], &st) < 0) {
    perror(argv[1]);
    exit(1);
  }

  printf("st_mode = %o\n", st.st_mode);
  printf("S_IFMT = %o\n", S_IFMT);
  printf("~S_IFMT = %o\n", ~S_IFMT);
  printf("mode = %o\n", st.st_mode & ~S_IFMT);

  return 0;
}
