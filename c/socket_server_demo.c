
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/_endian.h>
#include <sys/socket.h>
#include <unistd.h>
int main(void) {
  int listen_fd, conn_fd;
  struct sockaddr_in server_addr;

  listen_fd = socket(AF_INET, SOCK_STREAM, 0);
  if (listen_fd < 0) {
    perror("socket");
    close(listen_fd);
    exit(EXIT_FAILURE);
  }

  memset(&server_addr, 0, sizeof(server_addr));
  server_addr.sin_family = AF_INET;
  server_addr.sin_addr.s_addr = INADDR_ANY;
  server_addr.sin_port = htons(8081);

  if (bind(listen_fd, (struct sockaddr *)&server_addr, sizeof(server_addr)) <
      0) {
    perror("bind");
    close(listen_fd);
    exit(EXIT_FAILURE);
  }
  listen(listen_fd, 5);

  printf("listening...\n");

  conn_fd = accept(listen_fd, NULL, NULL);
  printf("accepted\n");

  write(conn_fd, "hello client\n", 13);
  sleep(5);

  close(conn_fd);
  close(listen_fd);

  return EXIT_SUCCESS;
}
