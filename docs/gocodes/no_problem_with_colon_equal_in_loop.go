package main

import "fmt"

func main() {
    a:=[]string{"abc","xyz","pqr","dfc"}
    for _,txt:=range a{
        b:=txt
        fmt.Println(b)
    }
}
