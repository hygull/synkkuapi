package main

import (
    "fmt"
    "strconv"
    "time"
)

func main() {
    timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
    fmt.Printf("%T %v",timestamp,timestamp) // prints: 1436773875771421417
    // now:=time.Now();
    // nanos := now.Unix()
    // millis := nanos / 1000000
    // fmt.Println(millis)
}