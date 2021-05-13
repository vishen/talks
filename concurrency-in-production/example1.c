#include<stdio.h>

static volatile int counter = 0;

void inc() {
	counter = counter + 1;
}

int main(int argc, char *argv[]) {
	printf("main: counter=%d\n", counter);	
	inc();
	printf("main: counter=%d\n", counter);	
}
