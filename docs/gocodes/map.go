package main

import "fmt"
import "encoding/json"

type User struct{
    Name string `json:"name"`
    Age int `json:"age"`
}

func main() {
   originMap:=map[string]User{};

   originMap["User1"]=User{"Rishikesh",24}
   originMap["User2"]=User{"Prajapati",24}
   originMap["User3"]=User{"Darshan",23};

   strBytes,_:= json.Marshal(originMap);//Marshalling
   fmt.Println("Unstructured form:-\n",string(strBytes),"\n\n");//Converting to JSON data

   strBytes,_= json.MarshalIndent(originMap,"","\t");//Marshalling
   fmt.Println("Structured form:-\n",string(strBytes),"\n");//Converting to JSON data  
}
