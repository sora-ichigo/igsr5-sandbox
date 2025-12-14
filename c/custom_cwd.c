#include <dirent.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/syslimits.h>
#include <unistd.h>

int main(void) {
  char *buf;

  buf = getcwd(NULL, 0);
  if (!buf) {
    perror("getcwd");
    exit(EXIT_FAILURE);
  }

  printf("buf: %s\n", buf);

  DIR *dir = opendir(".");
  if (!dir) {
    perror("opendir");
    exit(EXIT_FAILURE);
  }

  struct dirent *ent;
  int i = 0;
  while ((ent = readdir(dir)) != NULL) {
    i++;
    if (i > 5) {
      break;
    }
    printf("d_name: %s\n", ent->d_name);
  }
  closedir(dir);

  if (chdir("..") < 0) {
    perror("chdir");
    exit(EXIT_FAILURE);
  }

  buf = getcwd(NULL, 0);
  printf("buf: %s\n", buf);

  dir = opendir(".");
  if (!dir) {
    perror("opendir");
    exit(EXIT_FAILURE);
  }

  i = 0;

  while ((ent = readdir(dir)) != NULL) {
    i++;
    if (i > 5) {
      break;
    }
    printf("d_name: %s\n", ent->d_name);
  }
  closedir(dir);

  free(buf);

  return EXIT_SUCCESS;
}
