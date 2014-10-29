// +build linux
#include <unistd.h>
#include <sys/wait.h>
#include <stdio.h>

void sighandler(int sig) {
	printf(">> C Code: Signal received: %d\n", sig);
}

void setHandler(int sig, struct sigaction *new) {
	struct sigaction old;
	sigaction(sig, new, &old);
	if (old.sa_flags && SA_SIGINFO) {
		printf("WARNING: sa_sigaction used for signal %d!\n", sig);
	} 
	if (old.sa_handler != new->sa_handler) {
		printf("WARNING: sa_handler different for signal %d!\n", sig);
	}
}

void regpid() {
	struct sigaction new;
	
	new.sa_handler = sighandler;
	sigemptyset(&new.sa_mask);
	new.sa_flags = 0;
	setHandler(SIGCHLD, &new);
	setHandler(SIGHUP , &new);
}

int wait_any_pid() {
	pid_t ret;
	ret = waitpid(-1, NULL, WNOHANG);
	return (int)ret;
}

