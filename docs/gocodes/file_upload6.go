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
 	"strconv"
 	"reflect"
    "regexp"
    "encoding/json"
 )
func uploadHandler(w http.ResponseWriter, r *http.Request) {
// the FormFile function takes in the POST input id file
    supportedFileFormats:=[]string{"jpg", "png", "gif", "bmp","jpeg","PNG","JPG","JPEG","GIF","BMP","mp4","avi"}
 	file, header, err := r.FormFile("file")

	defer file.Close();

 	gopath:=os.Getenv("GOPATH");
 	if gopath==""{
 		fmt.Println("GOPATH not found");
 		return
 	}
 	fmt.Println("Great")
 	fmt.Println(reflect.TypeOf(header.Filename),header.Filename)
    fmt.Println("Work")

    re := regexp.MustCompile("^[a-zA-Z0-9_.]*$")
    if !re.MatchString(header.Filename){
        fmt.Println("The file name should be alphanumeric, only _ and . are permitted as special characters")
        return
    }

 	two:=strings.Split(header.Filename,".")
 	if len(two)<2{
 		fmt.Println("File name is not in proper format")
 		return
 	}
    IsFileFormatOk:=false;
    for _,extension:=range supportedFileFormats{
        if extension==two[len(two)-1]{
            IsFileFormatOk=true
            break
        }
    }
    if !IsFileFormatOk{
        MsgMap:=make(map[string]string )
        fmt.Println(`Only "jpg", "png", "gif", "bmp","jpeg","PNG","JPG","JPEG","GIF","BMP","mp4","avi" files are allowed`);
        MsgMap["Message"]="Only jpg,jpeg,png,gif,bmp,mp4,avi files are allowed"
        MsgStr,_:=json.Marshal(MsgMap);
        fmt.Fprintf(w,string(MsgStr));
        return
    }

 	str:=""
    for _,part:=range two[:len(two)-1]{
         str+=part  
    }
    if str==""{
        fmt.Println("There should be a valid file name")
        return
    }

   s:=strings.Join(strings.Fields(time.Now().String()[0:19]),":")
   chars:=[]string{" ","-",":"}
   for _,char:=range chars{
        s=strings.Replace(s,char,"_",-1)
   }

 	s=str+"_"+s+"__synkku."
 	newFileName:=s+two[len(two)-1]//strings.Join(strings.Fields(time.Now().String()[0:19])
 	i:=1
 	fmt.Println("File name decided...")
 	_,err=os.Stat(gopath+"/src/tmp/uploaded/"+newFileName);
 	for !os.IsNotExist(err){//Does not exist
 		s=newFileName
 		newFileName+=strconv.Itoa(i)
		i+=1
		_,err=os.Stat(gopath+"/src/tmp/uploaded/"+newFileName)
	}

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