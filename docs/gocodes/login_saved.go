/*
	Created on : 18/10/2016.
*/
package controllers

import (
	"database/sql"
	_ "github.com/lib/pq"

	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"regexp"
	"strconv"
	"strings"
	"time"
)

import (
	"synkkuapi/conf"
	"synkkuapi/controllers/validation"
	"synkkuapi/views"
)

/********************** Middleware (Token validator) ***********************/
func isAppSignInKeyCorrect(appSignInKeyMd5 string, authenticationToken string, email string) bool {
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

/*************************START (User login callee's)******************************************/

func getStoken(userId int, authenticationToken string) string {
	unixTime := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	tempStoken := strconv.Itoa(userId) + authenticationToken + unixTime //<userId><authenticationToken><unixTime>
	stoken := getMD5Hash(tempStoken)                                    //md5("<userId><authenticationToken><unixTime>")
	return stoken
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func IsEmailValid(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

/*************************END of callee of UserLogin2 (User login)***********************/

/********************** User Login ******************************************************/
func Login(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	AuthenticationKey := conf.AuthenticationKey
	ctx.SetContentType("application/json; charset=utf8")

	if !ctx.IsPost() { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(ctx, "PostMethodNotFoundError", "The data is not being sent using POST method", 107, 7)
		fmt.Println("The data is not being sent using POST method")
		return
	}

	if len(ctx.PostBody()) == 0 { //If there's no post data then inform the user...len([]byte )==0
		views.ShowSuccessOrErrorAsJSON(ctx, "NoPostDataError", "You haven't sent the POST data", 104, 4)
		fmt.Println("You haven't send the POST data")
		return
	}
	/***************************Getting data from URLs *************************************/

	umail := strings.TrimSpace(string(ctx.FormValue("umail")))

	app_sign_in_key := strings.TrimSpace(string(ctx.FormValue("appsigninkey"))) //app sign in key while first login (required)
	/****************************************************************************************/
	fmt.Println("Extracted email & app_sign_in_key : ",umail,app_sign_in_key)
	if !validation.AreAllFieldsAreInValidForm(umail, app_sign_in_key) {
		views.ShowSuccessOrErrorAsJSON(ctx, "EmptyFieldOrKeyNameError", "email & app_sign_in_key keys, both are mandatory for existing users, check if you have provided  wrong key names", 103, 3)
		return
	}

	if !IsEmailValid(umail) {
		views.ShowSuccessOrErrorAsJSON(ctx, "InvalidEmailFormatError", "The enetered email is not valid", 102, 2)
		fmt.Println("Email is not valid")
		return
	}
	if !isAppSignInKeyCorrect(app_sign_in_key, AuthenticationKey, umail) { //Call to a Middleware
		views.ShowSuccessOrErrorAsJSON(ctx, "InvalidAppSignKey", "App signin key is invalid", 101, 1)
		fmt.Println("The token is invalid")
		return
	}

	fmt.Println("Connecting to DB...")
	db, err := sql.Open("postgres", "postgres://"+conf.DBUserName+":"+conf.DBPassword+"@"+conf.DBHost+":"+conf.DBPort+"/"+conf.DBName+"?sslmode=disable")

	if err != nil {
		fmt.Println("Error in connection...")
		fmt.Println(err.Error())
		views.ShowSuccessOrErrorAsJSON(ctx, "DBConnectionError", "Error in connection with database", 104, 4)
		return
	}
	defer db.Close()

	fmt.Println("Connected")
	rows, err := db.Query("select email from auth_user where email= '" + umail + "' AND is_active=TRUE;")
	if err != nil {
		fmt.Println("Error in query...")
		fmt.Println(err.Error())
		views.ShowSuccessOrErrorAsJSON(ctx, "SelectQueryExecutionError", "Error in execution of select query,check the syntax", 108,8)
		return
	}

	found := false
	for rows.Next() {
		found = true
	}
	if found {
		rows, err = db.Query("select id from auth_user" +
			" where email='" + umail + "' ;") //I need to check about stoken
		if err != nil {
			fmt.Println("Error in query...")
			fmt.Println(err.Error())
			views.ShowSuccessOrErrorAsJSON(ctx, "QueryError", "Error in syntax of query", 106, 6)
			return
		}
		fmt.Println("Details extracted from DB")
		var stoken string
		var uid int
		rows.Next()
		rows.Scan(&uid)
		stoken = getStoken(uid, AuthenticationKey)

		q := "update auth_user set stoken='" + stoken + "' , stoken_created_on='" +
			time.Now().String()[0:19] + "' where email='" + umail + "';"
		fmt.Println(q)
		stmt, err := db.Prepare(q)
		if err != nil {
			fmt.Println("Error in execution of update query")
			views.ShowSuccessOrErrorAsJSON(ctx, "UpdateQueryExecutionError", "Error in execution of update query", 107, 7)
			return
		}
		stmt.Exec()
		fmt.Println("Details updated as the user is already having an account")
		fmt.Fprintf(ctx, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "uid","stoken"}, 1, "You successfully logged in", uid,stoken))
		//http.Redirect(rw, req, "/", 301)
	} else {
		views.ShowSuccessOrErrorAsJSON(ctx, "UserDoesNotExist", "You are not a registered user, go to http://"+conf.HostServerIP+":"+conf.Port+, 105, 5)
		return
	}
}
