package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"go-shell/pkg/built_ins"
	"go-shell/pkg/input_parser"
	"go-shell/pkg/core"

	"os"
)

func main(){

	reader := bufio.NewReader(os.Stdin)


  for {
	initial_directory, err := os.Getwd()

	if err!=nil {
		fmt.Fprintln(os.Stderr, "Error getting current directory : ",err)
	}

	fmt.Printf("[my_shell]@~%s: $ ",initial_directory)
    
	input,err := reader.ReadString('\n')

	if err!=nil{
		fmt.Fprintln(os.Stderr, "Error reading input : ",err)
		continue
	}
   
	var tokens []string
	if strings.Contains(input,"|"){
		core.ExecutePipe(input)
	} else {
		tokens = input_parser.TokenizeInput(input)
	}
	
	

	

   if(len(tokens)!=0) { 

		switch tokens[0] {
			case "cd": 
	    		built_ins.Cd(tokens,initial_directory)
		    case "echo":
				built_ins.Echo(tokens)
			case "which":
				built_ins.Which(tokens)
			case "export":
				built_ins.Export(tokens)
			case "pwd":
				built_ins.Pwd()
			case "exit":
				return
			default:
				cmd := exec.Command(tokens[0],tokens[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin
				cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Fprintln(os.Stderr,"Command failed: ", err)
			}
	 }
   }
  }
}