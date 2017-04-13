    /*
    two:=strings.Split(s,".")
    for _,part:=range two[:len(two)-1]{
         str+=part  
    }
    fmt.Println(two)
    fmt.Println(str)
    */

package main

import ("fmt"
 	"io"
 	"net/http"
 	"os"
 	"time"
 	"strings"
 	_"strconv"
 	"reflect"
 )
func uploadHandler(w http.ResponseWriter, r *http.Request) {
// the FormFile function takes in the POST input id file
	fmt.Println("Ok")
 	file, header, err := r.FormFile("file")
 	fmt.Println("Ok2")
	defer file.Close();
	fmt.Println("Ok3")
 	gopath:=os.Getenv("GOPATH");
 	if gopath==""{
 		fmt.Println("GOPATH not found");
 		return
 	}
 	fmt.Println("Great")
 	fmt.Println(reflect.TypeOf(header.Filename),header.Filename)

 	fmt.Println("Work")
 	two:=strings.Split(header.Filename,".")
 	if len(two)<2{
 		fmt.Println("File name is not in proper format")
 		return
 	}
 	str:=""
    for _,part:=range two[:len(two)-1]{
         str+=part  
    }

 	newFileName:=str+"__"+strings.Join(strings.Fields(time.Now().String()[0:19]),"")
 	newFileName=newFileName+"__synkku."+two[len(two)-1]//strings.Join(strings.Fields(time.Now().String()[0:19])
 	//i:=1
 	fmt.Println("File name decided...")
 	/*_,err=os.Stat(gopath+"/src/tmp/uploaded/"+newFileName);
 	for !os.IsNotExist(err){//Does not exist
 		s:=newFileName
 		newFileName=+strconv.Itoa(i)
		i+=1
		_,err=os.Stat(gopath+"/src/tmp/uploaded/"+newFileName)
	}*/

	out, err2 := os.Create(gopath+"/src/tmp/uploaded/"+newFileName)
	if err2 != nil {
		 		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		 		return
	}

 	defer out.Close()

 	// write the content from POST to the file
 	_, err2 = io.Copy(out, file)
 	if err2 != nil {
 		fmt.Fprintln(w, err)
 	}

 	fmt.Fprintf(w, "File uploaded successfully : ")
 	fmt.Fprintf(w, header.Filename)
 }

 func main() {
 	http.HandleFunc("/", uploadHandler)
 	http.ListenAndServe(":8080", nil)
 }