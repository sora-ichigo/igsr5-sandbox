#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>

static void do_cat(const char *path);
static void do_cat_without_path();

int main(int argc, char *argv[])
{
  if (argc < 2)
  {
    do_cat_without_path();
  };

  int i;
  for (i = 1; i < argc; i++)
  {
    do_cat(argv[i]);
  }
  return 0;
}

static void do_cat_without_path()
{
  unsigned char buf[1024];
  int n;

  for (;;)
  {
    n = read(STDIN_FILENO, buf, sizeof buf);
    if (n < 0)
    {
      perror("stdin");
      exit(1);
    }

    if (n == 0)
    {
      break;
    }

    if (write(STDOUT_FILENO, buf, n) < 0)
    {
      perror("stdout");
      exit(1);
    }
  }
}

static void do_cat(const char *path)
{
  int fd;
  unsigned char buf[1024];
  int n;

  fd = open(path, O_RDONLY);
  if (fd < 0)
  {
    perror(path);
    exit(1);
  }

  for (;;)
  {
    n = read(fd, buf, sizeof buf);
    if (n < 0)
    {
      perror(path);
      exit(1);
    }
    if (n == 0)
    {
      break;
    }

    if (write(STDOUT_FILENO, buf, n) < 0)
    {
      perror(path);
      exit(1);
    }
  }
  if (close(fd) < 0)
  {
    perror(path);
    exit(1);
  }
}
