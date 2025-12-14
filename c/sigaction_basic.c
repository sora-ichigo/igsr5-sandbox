#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/signal.h>
#include <unistd.h>

static volatile sig_atomic_t stop = 0;

void on_sigint(int signum) { stop = 1; }

int main(int argc, char *argv[]) {
  struct sigaction sa;
  sa.sa_handler = on_sigint;
  sigemptyset(&sa.sa_mask);
  sa.sa_flags = SA_RESTART;

  if (sigaction(SIGINT, &sa, NULL) == -1) {
    perror("sigaction");
    exit(EXIT_FAILURE);
  }

  while (!stop) {
    printf("working...\n");
    sleep(1);
  }

  printf("cleanup and exit.\n");

  return EXIT_SUCCESS;
}
