#include <stdio.h>
#include <stdlib.h>
#include <sys/_types/_timeval.h>
#include <sys/time.h>
#include <time.h>

int main(int argc, char *argv[]) {

  time_t t = time(NULL);
  if (t == (time_t)-1) {
    perror("time");
    exit(EXIT_FAILURE);
  }
  printf("time: %ld\n", t);

  struct timeval tv;
  struct timezone tz;
  if (gettimeofday(&tv, &tz) != 0) {
    perror("gettimeofday");
    exit(EXIT_FAILURE);
  }

  printf("timeval sec: %ld\n", tv.tv_sec);
  printf("timeval usec: %d\n", tv.tv_usec);
  printf("timezone dsttime: %d\n", tz.tz_dsttime);
  printf("timeval minuteswest: %d\n", tz.tz_minuteswest);

  return EXIT_SUCCESS;
}
