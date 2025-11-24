CC = gcc
CFLAGS = -Wall -Wextra -O2

OBJS = main.o math.o

myapp: $(OBJS)
	$(CC) $(OBJS) -o myapp

main.o: main.c math.h
	$(CC) $(CFLAGS) -c main.c

math.o: math.c math.h
	$(CC) $(CFLAGS) -c math.c

clean:
	rm -f *.o myapp
