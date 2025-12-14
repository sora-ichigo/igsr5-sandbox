#include <stdio.h>
#include <stdlib.h>
#include <sys/resource.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
  struct rusage usage;
  if (getrusage(RUSAGE_SELF, &usage) != 0) {
    perror("getrusage");
    exit(EXIT_FAILURE);
  }

  printf("ru_utime: %ld\n", usage.ru_utime.tv_sec);
  printf("ru_stime: %ld\n", usage.ru_stime.tv_sec);

  return EXIT_SUCCESS;
}
