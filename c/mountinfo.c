#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <sys/param.h>
#include <sys/ucred.h>
#include <sys/mount.h>
#include <util.h>

int main(){
	long mntsize;
	struct statfs *mntbuf;
	int i;
	
	mntsize = getmntinfo(&mntbuf, MNT_WAIT);
	printf("\n");
	printf("                Filesystem     512-blocks\n");
	for(i = 0; i < mntsize ; i++)
		printf("%27s\t%ld\n", mntbuf[i].f_mntfromname, mntbuf[i].f_blocks);
	printf("\n");
}
