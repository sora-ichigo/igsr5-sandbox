#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/_types/_pid_t.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
  int fds[2];
  if (pipe(fds) == -1) {
    perror("pipe");
    exit(EXIT_FAILURE);
  }

  pid_t pid = fork();
  if (pid < 0) {
    perror("fork");
    exit(EXIT_FAILURE);
  }

  if (pid == 0) {
    close(fds[1]);

    if (dup2(fds[0], STDIN_FILENO) == -1) {
      perror("dup2");
      exit(EXIT_FAILURE);
    }

    execlp("cat", "cat", (char *)NULL);

    perror("execlp");
    exit(EXIT_FAILURE);

  } else {
    close(fds[0]);

    const char *msg = "Hello child process\n"
                      "This is stdin redirected via pipe\n";
    if (write(fds[1], msg, strlen(msg)) < 0) {
      perror("write");
      exit(EXIT_FAILURE);
    }
    close(fds[1]);

    wait(&pid);
  }

  return EXIT_SUCCESS;
}
