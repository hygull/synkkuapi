package main
import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
)

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func reverse(s string)string{
	reversed_email:=[]byte(s); //string is immutable in Go
	//email reversing for loop
	for front,last:=0,len(reversed_email)-1;front<last;front,last=front+1,last-1{
		reversed_email[front],reversed_email[last]=reversed_email[last],reversed_email[front];
	}
	return string(reversed_email)
}

func main(){
	var email string
	fmt.Println(`Generate the md5( reverse(email)+"eZ$21#@54>4074W8Ndkf**WE32awe2376THWEKm")`);
	fmt.Print("Enter your email id : ");
	fmt.Scan(&email);

	fmt.Println("reverse("+email+") : ",reverse(email))
	fmt.Println("\n"+reverse(email)+"eZ$21#@54>4074W8Ndkf**WE32awe2376THWEKm")
	md5Version:=GetMD5Hash(reverse(email)+"eZ$21#@54>4074W8Ndkf**WE32awe2376THWEKm");
	fmt.Println("\n\nMD5("+reverse(email)+"eZ$21#@54>4074W8Ndkf**WE32awe2376THWEKm"+") : ",md5Version);
	fmt.Println()
}

 