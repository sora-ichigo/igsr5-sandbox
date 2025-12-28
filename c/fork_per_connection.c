#include <arpa/inet.h>
#include <netinet/in.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/_endian.h>
#include <sys/_types/_pid_t.h>
#include <sys/_types/_ssize_t.h>
#include <sys/errno.h>
#include <sys/signal.h>
#include <sys/socket.h>
#include <unistd.h>
int main(int argc, char *argv[]) {
  int listen_fd;
  listen_fd = socket(AF_INET, SOCK_STREAM, 0);
  if (listen_fd < 0) {
    perror("socket");
    exit(1);
  }

  signal(SIGCHLD, SIG_IGN);

  struct sockaddr_in addr;

  int opt = 1;
  if (setsockopt(listen_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt)) < 0) {
    perror("setsockopt");
    close(listen_fd);
    exit(1);
  }

  memset(&addr, 0, sizeof(addr));
  addr.sin_family = AF_INET;
  addr.sin_addr.s_addr = htonl(INADDR_ANY);
  addr.sin_port = htons(8081);

  if (bind(listen_fd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
    perror("bind");
    close(listen_fd);
    exit(1);
  }

  if (listen(listen_fd, 16) < 0) {
    perror("listen");
    close(listen_fd);
    exit(1);
  }

  printf("listening on 0.0.0.0:8081\n");

  while (1) {
    struct sockaddr_in peer;
    socklen_t peerlen = sizeof(peer);

    int conn_fd = accept(listen_fd, (struct sockaddr *)&peer, &peerlen);
    if (conn_fd < 0) {
      if (errno == EINTR)
        continue;
      perror("accept");
      close(listen_fd);
      exit(1);
    }

    pid_t pid = fork();
    if (pid < 0) {
      perror("fork");
      close(conn_fd);
      continue;
    }

    if (pid == 0) {
      close(listen_fd);

      char ip[INET_ADDRSTRLEN];
      inet_ntop(AF_INET, &peer.sin_addr, ip, sizeof(ip));
      printf("[child %d] accepted %s:%d\n", getpid(), ip, ntohs(peer.sin_port));

      const char *msg = "hello from forked server\n";
      (void)write(conn_fd, msg, strlen(msg));

      char buf[256];
      ssize_t n = read(conn_fd, buf, sizeof(buf) - 1);
      if (n > 0) {
        buf[n] = '\0';
        printf("[child %d] received: %s\n", getpid(), buf);
      }

      close(conn_fd);
      _exit(0);
    }

    close(conn_fd);
  }

  return EXIT_SUCCESS;
}
