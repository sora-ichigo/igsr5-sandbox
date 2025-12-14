#include <arpa/inet.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/_endian.h>
#include <sys/socket.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
  int sockfd;
  struct sockaddr_in server_addr;

  sockfd = socket(AF_INET, SOCK_STREAM, 0);
  if (sockfd < -1) {
    perror("socket");
    exit(EXIT_FAILURE);
  }

  memset(&server_addr, 0, sizeof(server_addr));
  server_addr.sin_family = AF_INET;
  server_addr.sin_port = htons(8081);
  inet_pton(AF_INET, "127.0.0.1", &server_addr.sin_addr);

  if (connect(sockfd, (struct sockaddr *)&server_addr, sizeof(server_addr)) <
      0) {
    perror("connect");
    close(sockfd);
    exit(EXIT_FAILURE);
  }

  printf("connected\n");
  if (write(sockfd, "hello\n", 6) < 0) {
    perror("write");
    close(sockfd);
    exit(EXIT_FAILURE);
  }

  close(sockfd);

  return EXIT_SUCCESS;
}
