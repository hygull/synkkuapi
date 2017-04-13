/*
	Created on : 18/10/2016.
*/
package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"strings"
	"time"
)

import (
	"synkkuapi/conf"
	"synkkuapi/controllers/middlewares"
	"synkkuapi/views"
)

/********************** User Login ******************************************************/
func AccountLogin(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	AuthenticationKey := conf.AuthenticationKey
	ctx.SetContentType("application/json; charset=utf8")

	if !ctx.IsPost() { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(ctx, "PostMethodNotFoundError", "The data is not being sent using POST method", 107, 7)
		fmt.Println("The data is not being sent using POST method")
		return
	}
	l := len(ctx.PostBody())
	if l == 0 { //If there's no post data then inform the user...len([]byte )==0
		views.ShowSuccessOrErrorAsJSON(ctx, "NoPostDataError", "You haven't sent the POST data", 104, 4)
		fmt.Println("You haven't send the POST data(as it is empty)")
		return
	}
	fmt.Println("Length of PostBody() : ", l)
	/***************************Getting data from URLs *************************************/
	email := strings.TrimSpace(string(ctx.FormValue("email")))

	app_sign_in_key := strings.TrimSpace(string(ctx.FormValue("app_sign_in_key"))) //app sign in key while first login (required)
	/****************************************************************************************/
	fmt.Println("Extracted URL Parameters : ", email, app_sign_in_key)
	if !middlewares.AreAllFieldsAreInValidForm(email, app_sign_in_key) {
		views.ShowSuccessOrErrorAsJSON(ctx, "EmptyFieldOrKeyNameError", "This error is because of one of these : (1) You have to provide atleast 2 details email & app_sign_in_key (2) You have provided a wrong key name", 103, 3)
		return
	}

	if !middlewares.IsEmailValid(email) {
		views.ShowSuccessOrErrorAsJSON(ctx, "InvalidEmailFormatError", "The enetered email is not valid", 102, 2)
		fmt.Println("Email is not valid")
		return
	}
	if !middlewares.IsAppSignInKeyCorrect(app_sign_in_key, AuthenticationKey, email) { //Call to a Middleware
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

	err = db.Ping()
	if err != nil {
		fmt.Println("Error in connection test...")
		fmt.Println(err.Error())
		views.ShowSuccessOrErrorAsJSON(ctx, "DBConnectionTestError", "Error in connection with database", 112, 12)
		return
	}

	fmt.Println("Connected...")
	rows, err := db.Query("select email from auth_users where email= '" + email + "' AND is_active=TRUE;")
	if err != nil {
		fmt.Println("Error in select query...")
		fmt.Println(err.Error())
		views.ShowSuccessOrErrorAsJSON(ctx, "SelectQueryExecutionError", "Error in syntax of select query", 109, 9)
		return
	}

	found := false
	for rows.Next() {
		found = true
	}

	if found {
		rows, err = db.Query("select id from auth_users" +
			" where email='" + email + "' ;") //I need to check about stoken
		if err != nil {
			fmt.Println("Error in query...")
			fmt.Println(err.Error())
			views.ShowSuccessOrErrorAsJSON(ctx, "SelectQueryExecutionError", "Error in syntax of query", 109, 9)
			return
		}
		fmt.Println("Details extracted from DB")
		var stoken string
		var uid int
		rows.Next()
		rows.Scan(&uid)
		stoken = middlewares.GetStoken(uid, AuthenticationKey)

		q := "update auth_users set token='" + stoken + "' , token_updated='" +
			time.Now().String()[0:19] + "' where email='" + email + "';"
		fmt.Println(q)
		stmt, err := db.Prepare(q)
		if err != nil {
			fmt.Println("Error in execution of update query")
			views.ShowSuccessOrErrorAsJSON(ctx, "UpdateQueryExecutionError", "Error in execution of update query", 107, 7)
			return
		}
		stmt.Exec()
		fmt.Println("Details updated as the user is already having an account")
		fmt.Fprintf(ctx, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "uid", "stoken"}, 1, "You successfully logged in", uid, stoken))
		//http.Redirect(rw, req, "/", 301)
	} else {
		views.ShowSuccessOrErrorAsJSON(ctx, "UserDoesNotExist", "You are not a registered user", 105, 5)
		return
	}
}
