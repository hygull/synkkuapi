package main

import "fmt"
import "time"
import "strings"

func main() {
   s:=strings.Join(strings.Fields(time.Now().String()[0:19]),"")
   a:=strings.Replace(s,"-","_",-1)
   fmt.Printf("%s",a);
   fmt.Println()
   b:=strings.Replace(a,":","_",-1)
   fmt.Printf("%s",b);
   fmt.Println()
}
