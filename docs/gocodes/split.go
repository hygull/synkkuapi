package main

import ("fmt";"strings")

func main() {
    s:="Rishikesh.jpeg.jggghhg.fggfgfgf.fhfhhfhfh.gdfdtrdf#fhf.python"
    str:=""
    two:=strings.Split(s,".")
    for _,part:=range two[:len(two)-1]{
         str+=part  
    }
    fmt.Println(two)
    fmt.Println(str)
    fmt.Println(two[len(two)-1])
    fmt.Printf("hello, world\n")
}
