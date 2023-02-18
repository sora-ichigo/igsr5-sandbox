#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void show_stack(int *addr, int count)
{
	int i, j;
	char values[4];
	printf("  address |value(int)|value(char)\n");
	printf("---------------------------------\n");
	for (i = 0; i < count; i++) {
		printf(" %08x | %08x | ", addr + i, *(addr + i));
		memcpy(values, (addr + i), 4);
		for (j = 0; j < 4; j++) {
			if (values[j] == '\0') {
				values[j] = '*';
			} else if (values[j] < 32 || values[j] > 126) {
				values[j] = '.';
			}
			printf("%c ", values[j]);
		}
		printf("\n");
	}
}

void getprofile(int arg)
{
	char firstname[4];
	char lastname[4];

	int id;
	printf("[ADDRESS]\n");
	printf("&ID=%08x, &(firstname[0])=%08x, &(lastname[0])=%08x\n", &id, firstname, lastname);

	printf("Input your id >");
	scanf("%d", &id);
	printf("Input your first name >");
	scanf("%s", firstname);
	printf("[BEFORE inputting the last name]\n");
	show_stack(&id, 10);
	printf("Input your last name >");
	scanf("%s", lastname);
	printf("[AFTER inputting the last name]\n");
	show_stack(&id, 10);
	printf("ID=%d, firstname=%s, lastname=%s\n", id, firstname, lastname);
}

void malicious_func()
{
	printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n");
	printf("Mwa ha ha! Welcome to malicious code\n");
	printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n");
	exit(1);
}

int main()
{
	getprofile(0x11111111);
	return 0;
}
