#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <fcntl.h>
#include <sys/wait.h>

int main()
{
	pid_t pid;
	int status;
	int fd;

	printf("Exec ps command...\n");
	pid = fork();
	if (pid == 0) {
		// 子プロセスの場合
		if ((fd = open("out.txt", O_WRONLY | O_CREAT | O_TRUNC, 0644)) == -1) {
			perror("open failed");
			exit(EXIT_FAILURE);
		}
		if (dup2(fd, STDOUT_FILENO) != STDOUT_FILENO) {
			perror("dup2 failed");
			close(fd);
			exit(EXIT_FAILURE);
		}
		execlp("ps", "ps", NULL);
		close(fd);
	} else {
		// 親プロセスの場合
		wait(&status);
		printf("done.\n");
	}

	return 0;
}
