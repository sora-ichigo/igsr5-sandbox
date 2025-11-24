#include <stdio.h>

static void log_internal(const char *msg) {
  fprintf(stderr, "[internal] %s\n", msg);
}

void internal_example() { log_internal("internal_example called"); }
