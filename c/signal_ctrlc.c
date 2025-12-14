#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

void on_sigint(int signum) {
  printf("caught SIGINT (%d). but I will not exit. \n", signum);
}

int main(int argc, char *argv[]) {
  if (signal(SIGINT, on_sigint) == SIG_ERR) {
    perror("signal");
    exit(EXIT_FAILURE);
  }

  printf("pid = %d\n", getpid());
  printf("try pressing Ctrl+C.\n");

  while (1) {
    printf("working...\n");
    sleep(1);
    raise(SIGINT);
  }

  return EXIT_SUCCESS;
}
