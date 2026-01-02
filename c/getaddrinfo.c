#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>

int main(int argc, char *argv[]) {
  struct addrinfo hints;
  struct addrinfo *result, *rp;
  int ret;

  memset(&hints, 0, sizeof(hints));
  hints.ai_family = AF_UNSPEC;
  hints.ai_socktype = SOCK_STREAM;
  hints.ai_flags = AI_ADDRCONFIG;

  ret = getaddrinfo("example.com", "80", &hints, &result);
  if (ret != 0) {
    fprintf(stderr, "getaddrinfo error: %s\n", gai_strerror(ret));
    return 1;
  }

  for (rp = result; rp != NULL; rp = rp->ai_next) {
    char host[NI_MAXHOST];
    char service[NI_MAXSERV];

    ret = getnameinfo(rp->ai_addr, rp->ai_addrlen, host, sizeof(host), service,
                      sizeof(service), NI_NUMERICHOST | NI_NUMERICSERV);
    if (ret != 0) {
      fprintf(stderr, "getnameinfo error: %s\n", gai_strerror(ret));
      continue;
    }

    printf("family=%s, address=%s, port=%s\n",
           rp->ai_family == AF_INET    ? "IPv4"
           : rp->ai_family == AF_INET6 ? "IPv6"
                                       : "Other",
           host, service);
  }

  freeaddrinfo(result);
  return EXIT_SUCCESS;
}
