#include <assert.h>
#include <ctype.h>
#include <dirent.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/resource.h>
#include <sys/signal.h>
#include <sys/stat.h>
#include <sys/syscall.h>
#include <sys/time.h>
#include <sys/types.h>
#include <sys/user.h>
#include <sys/wait.h>
#include <time.h>
#include <unistd.h>

#include "config.h"

int main(int argc, char *argv[]) {
  int pid;

  int lang = atoi(argv[1]);
  char *workdir = argv[2];

  chdir(workdir);

  const char *CP_C[] = {
      "gcc", "Main.c",   "-o",       "Main",           "-Wall",
      "-lm", "--static", "-std=c99", "-DONLINE_JUDGE", NULL};
  const char *CP_X[] = {
      "g++", "Main.cc",  "-o",         "Main",           "-Wall",
      "-lm", "--static", "-std=c++0x", "-DONLINE_JUDGE", NULL};
  const char *CP_J[] = {"javac", "-J-Xms32m", "-J-Xmx256m", "Main.java", NULL};

  pid = fork();
  if (pid == 0) {
    struct rlimit LIM;
    LIM.rlim_max = 30;
    LIM.rlim_cur = 30;
    setrlimit(RLIMIT_CPU, &LIM);  // longest compilation time is 30s.
    alarm(0);
    alarm(30);

    LIM.rlim_max = 100 * STD_MB;
    LIM.rlim_cur = 100 * STD_MB;
    setrlimit(RLIMIT_FSIZE, &LIM);

    LIM.rlim_max = STD_MB << 11;
    LIM.rlim_cur = STD_MB << 11;
    setrlimit(RLIMIT_AS, &LIM);

    freopen("ce.txt", "w", stderr);  // record copmile error

    switch (lang) {
      case LangC:
        execvp(CP_C[0], (char *const *)CP_C);
        break;
      case LangCC:
        execvp(CP_X[0], (char *const *)CP_X);
        break;
      case LangJava:
        execvp(CP_J[0], (char *const *)CP_J);
        break;
      default:
        exit(-1);
    }
    exit(0);
  } else {
    int status = 0;
    waitpid(pid, &status, 0);
    if (status == 0) {
      exit(0);
    } else {
      exit(1);
    }
  }
}