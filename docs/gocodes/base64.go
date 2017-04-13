package main

import "fmt"
import "encoding/base64"
import "reflect"
func main() {
    var email string 
    
    fmt.Print("Enter your email : ")
    fmt.Scan(&email)
    fmt.Println(email)
    fmt.Println("email(encoded) : ",base64.StdEncoding.EncodeToString([]byte(email)));
    fmt.Print("Enter the generated base64 to get original one : ")
    fmt.Scan(&email);
    decoded,_:=base64.StdEncoding.DecodeString(email) //decode is a slice of bytes
    fmt.Println(reflect.TypeOf(decoded))
    fmt.Println("email(decoded) : ",string(decoded))
}
