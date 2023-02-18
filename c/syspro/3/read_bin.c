#include <stdio.h>

int main () {
  FILE *finput = fopen("test.dat",  "rb");

  // read height from test.dat
  int age;
  char name[20];
  float height;
  fseek(finput, sizeof(age) + sizeof(name), SEEK_SET);
  fread(&height, sizeof(float), 1, finput);
  printf("height: %f\n", height);
  fclose(finput);

  // write height on a.out
  FILE *fout = fopen("a.out", "w");
  fprintf(fout, "%f\n", height);
  fclose(fout);

  return(0);
}
