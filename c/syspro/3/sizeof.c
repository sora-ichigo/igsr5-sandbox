#include <stdio.h>

int main()
{
	char *str11  = "";
	char str12[] = "";
	char *str21  = "hoge";
	char str22[] = "hoge";

	printf("%ld %s\n", sizeof(""), "\"\"");
	printf("%ld %s\n", sizeof("hoge"), "hoge");
	printf("%ld %s\n", sizeof(str11), str11);
	printf("%ld %s\n", sizeof(str12), str12);
	printf("%ld %s\n", sizeof(str21), str21);
	printf("%ld %s\n", sizeof(str22), str22);

	return 0;
}

