package main

import "fmt"
import "time"
import "strings"

func main() {
   s:=strings.Join(strings.Fields(time.Now().String()[0:19]),":")
   chars:=[]string{" ","-",":"}
   for _,char:=range chars{
   		s=strings.Replace(s,char,"_",-1)
   }
   fmt.Println("string : ",s);
}
