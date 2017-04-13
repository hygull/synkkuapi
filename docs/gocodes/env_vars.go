package main
import ("fmt";"os")

func main(){
	var env_name string
	fmt.Print("Type the name of ENV VAR :  ");
	fmt.Scan(&env_name)

	gopath:=os.Getenv(env_name);
	fmt.Println("$"+env_name +" : ",gopath)
}