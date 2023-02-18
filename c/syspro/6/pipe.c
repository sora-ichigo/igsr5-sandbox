#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <fcntl.h>
#include <sys/wait.h>

int main()
{
	pid_t pid;
	int status;
	int fds[2];
        fds[0]= open("out.txt", O_WRONLY | O_CREAT | O_TRUNC, 0644);
        fds[1]= open("out.txt", O_WRONLY | O_CREAT | O_TRUNC, 0644);

	printf("Exec ps command...\n");
	pid = fork();
	if (pid == 0) {
		// 子プロセスの場合
		if (dup2(fds[1], STDOUT_FILENO) != STDOUT_FILENO) {
			perror("dup2 failed");
			close(fds[1]);
			exit(EXIT_FAILURE);
		}
		execlp("ps", "ps", NULL);
                close(fds[1]); // error
	} else {
		// 親プロセスの場合
                if (dup2(fds[0], STDIN_FILENO) != STDIN_FILENO) {
                  perror("dup2 failed");
                  close(fds[0]);
                  exit(EXIT_FAILURE);
                }
                close(fds[1]);
		execlp("grep", "grep", "bash", NULL);
                close(fds[0]); // error
        }

        return 0;
}
