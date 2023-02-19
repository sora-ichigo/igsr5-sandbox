#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <fcntl.h>
#include <sys/wait.h>

int main()
{
	pid_t pid;
        int pipefd[2];

	printf("Exec ps command...\n");

        if (pipe(pipefd) < 0) {
          perror("failed to create pipe");
          exit(-1);
        }

	pid = fork();
	if (pid == 0) {
		// 子プロセスの場合
                close(pipefd[0]);
		dup2(pipefd[1], STDOUT_FILENO);
                execlp("ps", "ps", NULL);
	} else {
		// 親プロセスの場合
                close(pipefd[1]);
                dup2(pipefd[0], STDIN_FILENO);
                execlp("grep", "grep", "bash", NULL);
        }

        return 0;
}
