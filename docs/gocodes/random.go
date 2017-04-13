package main

import "fmt"
import "crypto/rand"
func main() {
   fmt.Printf("hello, world\n")
   fmt.Println(rand_str(10))
}
func rand_str(str_size int) string {
    alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes = make([]byte, str_size)
    fmt.Println("len(alphanum),  byte(len(alphanum)) :",len(alphanum),byte(len(alphanum)))
    
    fmt.Println("bytes:",bytes);
    rand.Read(bytes)
    fmt.Println("bytes:",string(bytes));
    for i, b := range bytes {
        fmt.Println("b : ",b)
        bytes[i] = alphanum[b%byte(len(alphanum))]
    }
    return string(bytes)
}
