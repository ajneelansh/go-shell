package config

import "os"

func IsExecutable(path string) bool {
   info, err := os.Stat(path)
   if err!= nil{
	return false
   }
   mode := info.Mode()
   return mode.IsRegular() && mode.Perm() & 0111 != 0
}