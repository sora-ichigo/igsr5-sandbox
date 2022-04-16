CFLAGS=-dtd=c11 -g -static

main: main.c
	gcc -O0 -o main main.c
test: main
	./test.sh
.PHONY: clean
clean:
	rm -f 9cc *.o *~ tmp*
