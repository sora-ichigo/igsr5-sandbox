#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
  printf("ruid=%d euid=%d gid=%d egid=%d\n", getuid(), geteuid(), getgid(),
         getegid());

  return EXIT_SUCCESS;
}
