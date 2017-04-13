//https://www.socketloop.com/tutorials/golang-regular-expression-alphanumeric-underscore

package main

 import "fmt"
 import "regexp"

 func main() {

  a := "testing_123"

  re := regexp.MustCompile("^[a-zA-Z0-9_.]*$")
  fmt.Println(re.MatchString("123"))
  fmt.Println(re.MatchString("abc"))
  fmt.Println(re.MatchString(a))
  fmt.Println(re.MatchString("世界"))

 }
O