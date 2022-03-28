#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
  if (argc != 2) {
    fprintf(stderr, "引数の個数が正しくありません\n");
    return 1;
  }

  char *p = argv[1];

  printf(".intel_syntax noprefix\n");
  printf(".globl main\n");
  printf("main:\n");
  // 最初の数字を取り出す
  printf("  mov rax, %ld\n", strtol(p, &p, 10));

  // 0個以上の項を解析
  while (*p) {
    if (*p == '+') { // + の場合
      p++;
      printf("  add rax, %ld\n", strtol(p, &p, 10));
      continue;
    }

    if (*p == '-') { // - の場合
      p++;
      printf("  sub rax, %ld\n", strtol(p, &p, 10));
      continue;
    }

    // それ以外 - エラー
    fprintf(stderr, "予期しない文字です: '%c'\n", *p);
    return 1;
  }

  printf("  ret\n");
  return 0;
}
