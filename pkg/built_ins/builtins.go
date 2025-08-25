package built_ins

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"go-shell/pkg/config"
)

func Cd(args []string, current_dir string){
	var dir string
   	if len(args) < 2 {
      	home,err:= os.UserHomeDir()
	  	if err!= nil {
			fmt.Println("cd: \n", err)
			return
	  	}
		dir = home
   	} else {
	dir = args[1]
   }

   err := os.Chdir(dir)
   if err!= nil {
	fmt.Println("cd: ",err)
   }
}

func Which(args []string){
	if(len(args)< 2){
		fmt.Fprintf(os.Stderr,"which: missing which arguments \n")
		return
	}

	cmd := args[1]

	if strings.Contains(cmd,"/"){
		if config.IsExecutable(cmd){
			fmt.Println(cmd);
		} else {
			fmt.Fprintf(os.Stderr, "which: %s not found\n", cmd)
		}
	}

	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv,string(os.PathListSeparator))
	 for _,dir := range paths{
		fullpath := filepath.Join(dir,cmd)
		if(config.IsExecutable(fullpath)){
			fmt.Println(fullpath)
			return
		} 
	 }
	fmt.Fprintf(os.Stderr,"which: %s not found\n",cmd)
		

}


func Echo(args []string){

	if len(args)<2{
		fmt.Printf("\n");
	}
	i:= 1
    
	if args[i][0] == '$'{
		val := args[1]
		val = strings.TrimPrefix(val,"$")
	  	value := os.Getenv(val);
	  	fmt.Println(value)
	} else {
    	for ;i<len(args);i++{
			fmt.Printf("%s ",args[i])
	   	}
	   	fmt.Printf("\n")
	}
}

func Pwd(){
	dir,err:= os.Getwd()
	if err!=nil{
		fmt.Println("Error getting current directory : \n", err)
		return
	}
	fmt.Printf("%s\n",dir);
}

func Export(args []string){
   if len(args)==1 {
      for _, env := range os.Environ() {
		fmt.Println(env)
		
	  }	
	  return
   }

   for _,arg := range args[1:]{
		if strings.Contains(arg,"="){
			parts:= strings.SplitN(arg,"=",2)
			name:= parts[0]
			value:= parts[1]
			os.Setenv(name,value)
		} else {
			val, exists := os.LookupEnv(arg)
			if exists {
				fmt.Printf("%s=%s\n", arg, val)
			} else {
				fmt.Fprintf(os.Stderr, "export: %s: not set\n", arg)
			}
		}
   }
}