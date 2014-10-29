package main

import "log"
import "github.com/bartmeuris/dkinit"
import "os"
import "os/signal"
import "os/exec"
import "syscall"

type DkProc struct {
	Cmd *exec.Cmd
}

func DkProcess(name string, arg ...string) *DkProc {
	ret := &DkProc{}
	ret.Cmd = exec.Command(name, arg...)
	if ret.Cmd == nil {
		return nil
	}
	// Pass the environment variables
	ret.Cmd.Env = os.Environ()
	return ret
}

func (dk *DkProc) Signal(sig os.Signal) error {
	return dk.Cmd.Process.Signal(sig)
}

func main() {
	c := make(chan os.Signal, 5)
	rc := make(chan int)
	//dkinit.Regpid()
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGCHLD, syscall.SIGTRAP, syscall.SIGQUIT, syscall.SIGHUP)
	go func(ch chan os.Signal, retch chan int) {
		defer close(retch)

		for s := range ch {
			log.Printf(">> GO: Got signal: %n", s.String())
			switch(s) {
			case os.Kill: fallthrough;
			case syscall.SIGQUIT: fallthrough;
			case os.Interrupt:
				log.Printf("Signal: exit!\n")
				retch <- 1
				return
			case syscall.SIGCHLD:
				for {
					r := dkinit.Waitanypid()
					if (r <= 0) {
						break;
					}
					log.Printf("Signal: Waitanypid caught signal from: %d\n", r)
				}
			}
		}
	}(c, rc)
	dkinit.Regpid()
	
	/*
	procattr := &os.ProcAttr {
		Dir: ".",
		Env: []string{ "TEST=BLA" },
		Files: []*os.File{nil, nil, nil},
	}
	*/
	go func(retch chan int){
		defer close(retch)

		cmd := exec.Command("/bin/sleep", "30")
		err := cmd.Start()
		if err != nil {
			log.Printf("Error occured: %s\n", err)
		} else {
			log.Printf("Process launched successfully\n")
			cmd.Wait()
			log.Printf("Process exited\n")
		}
	}(rc)

	<- rc
	log.Printf("<< Exit!\n")
}
