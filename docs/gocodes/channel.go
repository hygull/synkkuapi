package main
import "fmt"
func main(){
	namesChan := make(chan string)
	namesChan <- "Rishikesh Agrawani"
	namesChan <- "Golang"

	fmt.Println(<-namesChan)
	fmt.Println(<-namesChan)
}
