package main
import ("fmt";"os")

func main(){
	var fname string
	fmt.Print("Enter the name of file that you wanna search inside /tmp/ : ");
	fmt.Scan(&fname)
	gopath:=os.Getenv("GOPATH");
 	if gopath==""{
 		fmt.Println("GOPATH not found");
 		return
 	}
	if _,err:=os.Stat(gopath+"/src/tmp/uploaded/"+fname);os.IsNotExist(err){
		fmt.Println("Does not exist")
	}else{
		fmt.Println("Does exist");
	}
}