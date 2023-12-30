#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
  int i;

  for (i = 1; i < argc; i++)
  {
    FILE *f;
    f = fopen(argv[i], "r");
    if (!f)
    {
      perror(argv[i]);
      exit(1);
    }

    int c;
    while ((c = fgetc(f)) != EOF)
    {
      if (putchar(c) < 0)
      {
        perror(argv[i]);
        exit(1);
      }
    }
    fclose(f);
  }
  return 0;
}
