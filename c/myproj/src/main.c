#include "math.h"
#include "util.h"

int main(void) {
  int a = add(3, 4);
  int b = mul(2, 5);

  print_result("a", a);
  print_result("b", b);

  return 0;
}
