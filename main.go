package main

import "fmt"

import (
	"os"
	"strconv"
	"synkkuapi/conf"
	"synkkuapi/routers"
)

/******************* Start the Server ***************************************/

func main() {
	//To check whether the port specified in config.go is an integer and it is valid or not
	isThereAnyFaultInPortSetting := false
	_, err := strconv.Atoi(conf.Port)

	//If user specifies Port as   "9000a" , "800a" , "8abc" , "9a0b" , "ab89" , "" which is not convertible into integer
	if err != nil {
		if conf.Port == "" {
			fmt.Println("[synkku:error] You have specified blank port")
		} else {
			fmt.Println("[synkku:error] You have specified port as : ", conf.Port)
		}
		fmt.Println("[synkku:error] It is not valid, please try with 8000, 9000 etc")
		isThereAnyFaultInPortSetting = true
	}

	//Console message to show,the status of working API
	fmt.Println("\n[synkku] **************** Welcome to ****************")
	fmt.Println("[synkku] ****************** SYNKKU ******************\n")
	fmt.Print("\nSYNKKU Server is running on " + conf.HostServerIP + ":" + conf.Port + "\n\n")
	if isThereAnyFaultInPortSetting {
		fmt.Println("[synkku] Server stopped")
		return
	} else {
		err = routers.RequestRouters()
		if err != nil {
			os.Exit(1)
		}
	}
}
