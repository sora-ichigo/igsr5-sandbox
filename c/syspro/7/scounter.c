#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/sem.h>
#include <sys/shm.h>

// dead lock 再現手順
// 1. process 1 でセグフォ1に対して減算を行う
// 2. process 2 でセグフォ2に対して減算を行う
// 3. process 1 でセグフォ2に対して減算を行う
// 4. process 2 でセグフォ1に対して減算を行う
//
// 以下にprocessごとの再現ログを示す

/*
 *
 * process 1
 *
 *  $ ./scounter
 *  shared memory ID 1 = 65560
 *  shared memory ID 2 = 65561
 *  select operation from following options.
 *  0. subtract to semaphore 1.
 *  1. add to semaphore 1.
 *  2. subtract to semaphore 2.
 *  3. add to semaphore 2.
 *  4. exit
 *
 *  select(0~4):0
 *
 *  subtract to semaphore 1.
 *
 *  select operation from following options.
 *  0. subtract to semaphore 1.
 *  1. add to semaphore 1.
 *  2. subtract to semaphore 2.
 *  3. add to semaphore 2.
 *  4. exit
 *
 *  select(0~4):2
 *  // dead lock!!
 */

/*
 * process 2
 *
 * $ ./scounter 65560 65561
 * shared memory ID 1 = 65560
 * shared memory ID 2 = 65561
 *
 * select operation from following options.
 * 0. subtract to semaphore 1.
 * 1. add to semaphore 1.
 * 2. subtract to semaphore 2.
 * 3. add to semaphore 2.
 * 4. exit
 *
 * select(0~4):2
 *
 * subtract to semaphore 2.
 *
 * select operation from following options.
 * 0. subtract to semaphore 1.
 * 1. add to semaphore 1.
 * 2. subtract to semaphore 2.
 * 3. add to semaphore 2.
 * 4. exit
 *
 * select(0~4):0
 *
 * // dead lock!!
 */

void find_or_initialize_sem(int* id, int* sid) {
  if (*id != -1) {
    // セマフォIDの取得
    if ((*sid = semget(*id, 1, 0)) == -1) {
      perror("semget");
      exit(EXIT_FAILURE);
    }

    return;
  }

  union semun inival;

  // 共有メモリ領域の確保
  if ((*id = shmget(IPC_PRIVATE, sizeof(int), (IPC_CREAT | 0666))) == -1) {
    perror("shmget");
    exit(EXIT_FAILURE);
  }

  // セマフォの確保
  if ((*sid = semget(*id, 1, (IPC_CREAT | 0700))) == -1) {
    perror("semget");
    exit(EXIT_FAILURE);
  }

  // セマフォの初期設定
  inival.val = 1;
  semctl(*sid, 0, SETVAL, inival);
  return;
}

int main(int argc, char *argv[])
{
  int id_1 = -1, id_2 = -1, sid_1, sid_2;
  int *c1, *c2;
  int i;

  union semun inival;
  struct sembuf semlock;
  struct shmid_ds buf;

  semlock.sem_num = 0;
  semlock.sem_flg = 0;

  if (argc == 3) {
    id_1 = atoi(argv[1]);
    id_2 = atoi(argv[2]);
  }

  find_or_initialize_sem(&id_1, &sid_1);
  find_or_initialize_sem(&id_2, &sid_2);

  printf("shared memory ID 1 = %d\n", id_1);
  printf("shared memory ID 2 = %d\n", id_2);

  int op; // 0 or 1 or 2 or 3 or 4
  while (1) {
    printf("\nselect operation from following options.\n");
    printf("0. subtract to semaphore 1.\n");
    printf("1. add to semaphore 1.\n");
    printf("2. subtract to semaphore 2.\n");
    printf("3. add to semaphore 2.\n");
    printf("4. exit\n");

    printf("\nselect(0~4):");
    scanf("%d", &op);

    if (op == 4) {
      printf("exit\n");
      break;
    }

    else if (op >= 0 && op <= 3) {
      switch (op) {
        case 0:
          // セマフォ1の減算
          semlock.sem_num = 0;
          semlock.sem_op = -1;
          semlock.sem_flg = 0;
          semop(sid_1, &semlock, 1);

          printf("\n");
          printf("subtract to semaphore 1.\n");
          break;
        case 1:
          // セマフォ1の加算
          semlock.sem_num = 0;
          semlock.sem_op = 1;
          semlock.sem_flg = 0;
          semop(sid_1, &semlock, 1);

          printf("\n");
          printf("add to semaphore 1.\n");
          break;
        case 2:
          // セマフォ2の減算
          semlock.sem_num = 0;
          semlock.sem_op = -1;
          semlock.sem_flg = 0;
          semop(sid_2, &semlock, 1);

          printf("\n");
          printf("subtract to semaphore 2.\n");
          break;
        case 3:
          // セマフォ2の加算
          semlock.sem_num = 0;
          semlock.sem_op = 1;
          semlock.sem_flg = 0;
          semop(sid_2, &semlock, 1);

          printf("\n");
          printf("add to semaphore 2.\n");
          break;
      }
    }
  }

  shmctl(id_1, IPC_STAT, &buf);
  shmctl(id_2, IPC_STAT, &buf);

  // 共有メモリ領域の解放
  if (buf.shm_nattch == 1) {
    if (shmctl(id_1, IPC_RMID, 0) == -1) {
      perror("shmctl");
    }
    if (shmctl(id_2, IPC_RMID, 0) == -1) {
      perror("shmctl");
    }
  }

  if(argc == 1) {
    // セマフォの解放
    if (semctl(sid_1, IPC_RMID, 0) == -1) {
      perror("semctl");
    }
    if (semctl(sid_2, IPC_RMID, 0) == -1) {
      perror("semctl");
    }
  }

  return 0;
}

