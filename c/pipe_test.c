#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/_types/_pid_t.h>
#include <sys/_types/_ssize_t.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
  int fds[2];
  pid_t pid;

  if (pipe(fds) == -1) {
    perror("pipe");
    exit(1);
  }

  pid = fork();
  if (pid < 0) {
    perror("fork");
    exit(1);
  }

  if (pid == 0) {
    // 子プロセス
    close(fds[1]);

    char buf[128];
    ssize_t n;
    while ((n = read(fds[0], buf, sizeof(buf))) > 0) {
      fwrite(buf, 1, n, stdout);
    }

    if (n < 0) {
      perror("read");
      exit(1);
    }
    buf[n] = '\0';

    printf("Child received EOF\n");

    close(fds[0]);

  } else {
    // 親プロセス
    close(fds[0]);

    const char *msg = "Hello from parent\n";
    if (write(fds[1], msg, strlen(msg)) < 0) {
      perror("write");
      exit(1);
    }

    close(fds[1]);

    wait(NULL);
  }

  return 0;
}
