package main
import "fmt"
import "crypto/md5"

func main(){
	var str string
	fmt.Print("Enter a string whose md5 version you wanna see : ")
	fmt.Scan(&str)
	fmt.Println("\n")
	fmt.Printf("%T",md5.Sum([]byte(str))) //[]uint8
	fmt.Printf("%v %x",str+" (in md5 converted form)                  :",md5.Sum([]byte(str)));
	fmt.Println();
}