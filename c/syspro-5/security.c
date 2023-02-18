#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void success()
{
	printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n");
	printf("	    Welcome\n");
	printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n");
}

void judge(char *pw)
{
	char answer[10] = "Yasoshima";
	if(strcmp(pw,answer)== 0){
		success();
	}
	else{
		printf("Authentication Failed\n");
	}
}

void setting()
{
	char password[10];

	printf("Password : ");
	scanf("%s", password);
	judge(password);
}

int main()
{
	setting();
	return 0;
}
