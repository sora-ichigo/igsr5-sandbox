#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/signal.h>
#include <unistd.h>

static volatile sig_atomic_t stop = 0;

void usr1_handler(int sig) { stop = 1; }

int main(int argc, char *argv[]) {

  struct sigaction sa = {0};
  sa.sa_handler = usr1_handler;
  sigemptyset(&sa.sa_mask);
  sa.sa_flags = SA_RESTART;
  sigaction(SIGINT, &sa, NULL);
  sigaction(SIGTERM, &sa, NULL);

  sigset_t block, old;
  sigemptyset(&block);
  sigaddset(&block, SIGINT);
  sigaddset(&block, SIGTERM);
  sigprocmask(SIG_BLOCK, &block, &old);

  printf("pid=%d, waiting SIGINT/SIGTERM...\n", getpid());

  while (!stop) {
    sigset_t waitmask;
    sigemptyset(&waitmask);
    sigsuspend(&waitmask);
  }

  sigprocmask(SIG_SETMASK, &old, NULL);

  printf("cleanup and exit.\n");

  return EXIT_SUCCESS;
}
