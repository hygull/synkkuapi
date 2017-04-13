package middlewares

import (
	"crypto/md5" 
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv" 
	"time"
)

/******************* Middleware (Post data validator) ******************************/
func AreAllFieldsAreInValidForm(postData ...string) bool {
	isAnyBlank := false
	for _, data := range postData {
		if data == "" {
			isAnyBlank = true
			break
		}
	}
	fmt.Println("Is Any Field blank : ", isAnyBlank)
	return !isAnyBlank //An efficient way to check for empty fields
}

/********************** Middleware (Token validator) ***********************/
func IsAppSignInKeyCorrect(appSignInKeyMd5 string, authenticationToken string, email string) bool {
	reversed_email := []byte(email) //string is immutable in Go
	//email reversing for loop
	for front, last := 0, len(reversed_email)-1; front < last; front, last = front+1, last-1 {
		reversed_email[front], reversed_email[last] = reversed_email[last], reversed_email[front]
	}
	reversed_email_and_auth_token := string(reversed_email) + authenticationToken

	fmt.Printf("Reversed email and auth token      :  %v (%T)", reversed_email_and_auth_token, reversed_email_and_auth_token)
	fmt.Println()
	fmt.Println("\nYour app generated app_sign_key : ", appSignInKeyMd5)

	md5_of_reversed_email_and_auth_token := getMD5Hash(reversed_email_and_auth_token)
	fmt.Println("Server generated app_sign_key   : ", md5_of_reversed_email_and_auth_token)

	return md5_of_reversed_email_and_auth_token == appSignInKeyMd5
}

/****************************** Midaleware (Email validator) ************************************************/
func IsEmailValid(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

/***************************** Callee's of middleware functions *********************************************/
/*********** If everything is fine then generate token (middleware will call this function) *****************/
func GetStoken(userId int, authenticationToken string) string {
	unixTime := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	tempStoken := strconv.Itoa(userId) + authenticationToken + unixTime //<userId><authenticationToken><unixTime>
	stoken := getMD5Hash(tempStoken)                                    //md5("<userId><authenticationToken><unixTime>")
	return stoken
}

/************************** Generating md5 hash of the string parameter *************************************/
func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
