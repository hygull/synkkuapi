package main

import "fmt"
import "encoding/base64"
import _ "reflect"
import "regexp"
import "strings"

func main() {
    var email string 
    
    fmt.Print("Enter your email : ")
    fmt.Scan(&email)
    //Validating email(Checking whether it is empty or not)
    if email==""{
        fmt.Println("You have just pressed the RETURN/ENTER key")
        return
    }
    //If not empty then check for the format
    if !IsEmailValid(email){
        fmt.Println("Enter a proper email")
        return
    }
    //Without reversing the email
    email_synkku_base64:=base64.StdEncoding.EncodeToString([]byte(email+".synkku"))
    fmt.Printf("%-72v : %-20v\n","email(in original form)  ",email)
    fmt.Printf("%-72v : %-20v\n","synkku (in encoded form)  ",base64.StdEncoding.EncodeToString([]byte("synkku")))
    fmt.Printf("%-72v : %-20v\n","email.synkku (in base64 encoded form)  ",email_synkku_base64);

    //After reversing the email
    reversed_email_slice:=[]rune(email) //string is immutable in Go
    for front,last:=0,len(reversed_email_slice)-1;front<last;front,last=front+1,last-1{
        reversed_email_slice[front],reversed_email_slice[last]=reversed_email_slice[last],reversed_email_slice[front]
    }
    reversed_email_synkku_base64:=base64.StdEncoding.EncodeToString([]byte(string(reversed_email_slice)+".synkku"))
    fmt.Printf("%-72v : %-20v\n","reversed_email_slice (in original form)  ",string(reversed_email_slice))
    fmt.Printf("%-72v : %-20v\n","synkku (in encoded form)  ",base64.StdEncoding.EncodeToString([]byte("synkku")))
    fmt.Printf("%-72v : %-20v\n",string(reversed_email_slice)+".synkku (in base64 encoded form)  ",reversed_email_synkku_base64);

    fmt.Println("\n")
    var new_entered_base64 string
    fmt.Print("Enter the last generated base64 encoded form to get original email       : ")
    fmt.Scan(&new_entered_base64);
    fmt.Println();
    decoded,_:=base64.StdEncoding.DecodeString(new_entered_base64); //decode is a slice of bytes
    //fmt.Println(reflect.TypeOf(decoded))
    if new_entered_base64==reversed_email_synkku_base64{
        fmt.Printf("%-72v : %-20v",reversed_email_synkku_base64+"(decoded)  ",string(decoded));
        fmt.Println();
    }else{
        fmt.Println("You have not pasted properly");
        return;
    }

    str:=""
    regained_str_synkku_reversed,_:=base64.StdEncoding.DecodeString(string(reversed_email_synkku_base64)); 
    two:=strings.Split(string(regained_str_synkku_reversed),".");
    for _,part:=range two[:len(two)-1]{
         str+=part  
    }
    fmt.Printf("%-72v : %-20v","Concatenated format  ",str);
    fmt.Println();
    fmt.Printf("%-72v : %-20v","Reversed email format  ",str);

    new_reversed_slice:=[]rune(str)
    for front,last:=0,len(new_reversed_slice)-1;front<last;front,last=front+1,last-1{
        new_reversed_slice[front],new_reversed_slice[last]=new_reversed_slice[last],new_reversed_slice[front];
    }
    fmt.Println();
    fmt.Printf("%-72v : %-20v","Original email ",string(new_reversed_slice));
    fmt.Println("\n")//to bring the cursor to bottom
}

func IsEmailValid(email string) bool {
    Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return Re.MatchString(email);
}
