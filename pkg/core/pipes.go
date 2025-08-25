package core

import (
	"os"
	"os/exec"
	"strings"
)

func ExecutePipe(input string) error{
  cmds := strings.Split(input, "|")

  for i := range cmds {
		cmds[i] = strings.TrimSpace(cmds[i])
	}


  var executors []*exec.Cmd 

  for _,c := range cmds {
	args := strings.Fields(c)
	cmd := exec.Command(args[0],args[1:]...)
	executors = append(executors, cmd)
  }

  
  for i:= 0;i<len(executors)-1;i++ {
	pipe,err := executors[i].StdoutPipe()
	if err!= nil {
		return err
	}

	executors[i+1].Stdin = pipe
  }

	executors[len(executors)-1].Stdout = os.Stdout
  	executors[0].Stdin = os.Stdin
  	executors[len(executors)-1].Stderr = os.Stderr

 for _, cmd := range executors {
		if err := cmd.Start(); err != nil {
			return err
		}
	}

	for _, cmd := range executors {
		if err := cmd.Wait(); err != nil {
			return err
		}
	}

	return nil
}


// executors are collection of commands

// second for loop connects the pipe i.e treat output of first command as input of its next

// last value of executors contains final result
