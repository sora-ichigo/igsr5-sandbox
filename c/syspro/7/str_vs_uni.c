#include <stdio.h>
#include <stdlib.h>

struct parameter {
    int     width;
    float   height;
} data;

int w;
float h;

int main() {
    data.height = 10.0;
    data.width = 5;

    w = data.width;
    h = data.height;

    printf("width = %d\n", w);
    printf("height = %f\n", h);

    return 0;
}


