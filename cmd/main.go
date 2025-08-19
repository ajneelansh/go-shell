package main

import (
	"bufio"
	"fmt"
	"os"
	"go-shell/pkg/input_parser"
)

func main(){

	reader := bufio.NewReader(os.Stdin)

  for {
	  
	  fmt.Printf("[my_shell]:~$")
    
	 input,err := reader.ReadString('\n')

	 if err!=nil{
		fmt.Fprintln(os.Stderr, "Error reading input : ",err)
		continue
	 }

	 fmt.Println(input);
     tokens := input_parser.TokenizeInput(input)
	 fmt.Println(tokens)


  }
}