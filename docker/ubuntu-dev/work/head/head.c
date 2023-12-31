#include <stdio.h>
#include <stdlib.h>

static void do_head(FILE *f, long n);

int main(int argc, char **argv)
{
  if (argc < 2)
  {
    perror("Usage: head FILE");
    exit(1);
  }

  long n;
  n = atol(argv[1]);

  do_head(stdin, n);

  return 0;
}

static void do_head(FILE *f, long n)
{
  if (n <= 0)
  {
    return;
  }

  int c;
  while ((c = getc(f)) != EOF)
  {
    if (putchar(c) < 0)
      exit(1);

    if (c == '\n')
    {
      n--;
      if (n == 0)
      {
        return;
      }
    }
  }
}
