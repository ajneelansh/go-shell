package main

import (
	"bufio"
	"fmt"
	// "go-shell/pkg/built-ins"
	"go-shell/pkg/built_ins"
	"go-shell/pkg/input_parser"
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

    tokens := input_parser.TokenizeInput(input)
	
	

	

   if(len(tokens)!=0) { 

		switch tokens[0] {
			case "cd": 
	    		built_ins.Cd(tokens,initial_directory)
		    case "echo":
				built_ins.Echo(tokens)
			case "which":
				// built_ins.Which(tokens,initial_directory)
			case "env":
				// built_ins.Env(tokens,initial_directory)
			case "pwd":
				built_ins.Pwd()
			case "exit":
				return
	  
	 }


   }


  }
}