 package main

 import (
 	"fmt"
 	"io"
 	"net/http"
 	"os"
 	"time"
 	"strings"
 )

 func uploadHandler(w http.ResponseWriter, r *http.Request) {

 	// the FormFile function takes in the POST input id file
 	file, header, err := r.FormFile("file")

 	if err != nil {
 		fmt.Fprintln(w, err)
 		return
 	}

 	defer file.Close();
 	gopath:=os.Getenv("GOPATH");
 	if gopath==""{
 		fmt.Println("GOPATH not found");
 		return
 	}
 	newFileName:=strings.Join(strings.Fields(time.Now().String()[0:19]),"")

 	out, err := os.Create(gopath+"/src/tmp/uploaded/"+"Media_"+newFileName)
 	if err != nil {
 		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
 		return
 	}

 	defer out.Close()

 	// write the content from POST to the file
 	_, err = io.Copy(out, file)
 	if err != nil {
 		fmt.Fprintln(w, err)
 	}

 	fmt.Fprintf(w, "File uploaded successfully : ")
 	fmt.Fprintf(w, header.Filename)
 }

 func main() {
 	http.HandleFunc("/", uploadHandler)
 	http.ListenAndServe(":8080", nil)
 }