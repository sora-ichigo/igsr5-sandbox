#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/signal.h>
#include <unistd.h>

void handler(int sig) { printf("handler: got %d\n", sig); }

int main(int argc, char *argv[]) {
  struct sigaction sa = {0};
  sa.sa_handler = handler;
  sigemptyset(&sa.sa_mask);
  sigaction(SIGUSR1, &sa, NULL);

  sigset_t block;
  sigemptyset(&block);
  sigaddset(&block, SIGUSR1);
  sigprocmask(SIG_BLOCK, &block, NULL);
  printf("SIGUSR1 is now BLOCKED. pid=%d\n", getpid());

  printf("send SIGUSR no. I will NOT handle it yet.\n");
  sleep(5);

  sigset_t pend;
  sigpending(&pend);
  if (sigismember(&pend, SIGUSR1)) {
    printf("SIGUSR1 is pending.\n");
  } else {
    printf("SIGUSR1 not pending.\n");
  }

  printf("unblocking SIGUSR1.\n");
  sigprocmask(SIG_UNBLOCK, &block, NULL);

  sleep(1);

  return EXIT_SUCCESS;
}
