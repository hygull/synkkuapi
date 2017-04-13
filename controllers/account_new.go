/*
	Created on : 20/10/2016.
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

/*******************************End of callees********************************************************/

func AccountNew(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	AuthenticationKey := conf.AuthenticationKey
	ctx.SetContentType("application/json; charset=utf8")
	fmt.Println("Visiting : ", ctx.Path())

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
	/***************************Getting data from URLs*************************************/
	uname := ""
	fname := ""
	lname := ""
	cmtyid := "1"
	ppic_url := ""
	pvideo_url := ""

	email := strings.TrimSpace(string(ctx.FormValue("email")))
	uname = strings.TrimSpace(string(ctx.FormValue("uname")))
	fname = strings.TrimSpace(string(ctx.FormValue("fname")))
	lname = strings.TrimSpace(string(ctx.FormValue("lname")))
	dob := strings.TrimSpace(string(ctx.FormValue("dob")))
	cmtyid = strings.TrimSpace(string(ctx.FormValue("cmtyid")))                    //Default changes need to be removed
	app_sign_in_key := strings.TrimSpace(string(ctx.FormValue("app_sign_in_key"))) //app sign in key while first login (required)
	ppic_url = strings.TrimSpace(string(ctx.FormValue("p_pic")))
	pvideo_url = strings.TrimSpace(string(ctx.FormValue("p_vid"))) //It is not in endpoints spec...
	ltype := strings.TrimSpace(string(ctx.FormValue("ltype")))
	catid := strings.TrimSpace(string(ctx.FormValue("catid"))) //We have taken this...

	if catid == "" {
		catid = "1"
	}
	if ltype == "" {
		ltype = "1" //Google login...It can be changed later
	}
	if dob == "" {
		dob = "1900-01-01"
	}

	/****************************************************************************************/
	fmt.Println("Extracted URL Parameters : ", email, uname, fname, lname, dob, cmtyid, app_sign_in_key, ppic_url, pvideo_url, cmtyid, catid, app_sign_in_key)
	if !middlewares.AreAllFieldsAreInValidForm(email, app_sign_in_key) {
		views.ShowSuccessOrErrorAsJSON(ctx, "EmptyFieldOrKeyNameError", "You have to provide all fields (email,uname, fname,lname,dob,doj,app_sign_in_key) except lname or check if  you have provided a wrong key name", 103, 3)
		return
	}

	if !middlewares.IsEmailValid(email) {
		views.ShowSuccessOrErrorAsJSON(ctx, "InvalidEmailFormatError", "The enetered email is not valid", 102, 2)
		fmt.Println("Email is not valid")
		return
	}
	fmt.Println("Is App sign in key for first login correct : ", app_sign_in_key == AuthenticationKey)
	fmt.Println(AuthenticationKey, len(AuthenticationKey))
	fmt.Println(app_sign_in_key, len(app_sign_in_key))
	if app_sign_in_key != AuthenticationKey {
		views.ShowSuccessOrErrorAsJSON(ctx, "InvalidFirstLoginAppSignKeyError", "You have not specified the proper App SignIn Key required for first login. Demand it from your App sevice provider", 114, 14)
		fmt.Println("First time login requires a fixed app sign key...that will be provided to you by the App service provider")
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
		fmt.Println("Error in query...")
		fmt.Println(err.Error())
		views.ShowSuccessOrErrorAsJSON(ctx, "SelectQueryExecutionError", "Error in syntax of select query", 109, 9)
		return
	}

	found := false
	for rows.Next() {
		found = true
	}
	if found {
		fmt.Println("You don't need to create an account...as this email is already registered")
		views.ShowSuccessOrErrorAsJSON(ctx, "AccountRecreationAttemptError", "This email ID is already registered...so please login", 113, 13)
		return
	}

	timeNow := time.Now().String()[0:19]
	doj := timeNow //This variable is not required...this is for understaning
	stoken := "^synkku$"
	query := "insert into auth_users(id, user_name, email, first_name, last_name, dob, date_joined, " +
		"community_id,profile_pic, profile_video,token , token_updated, is_active,is_superuser,deleted_on,login_type,category_id)" +
		" values( default, " + "'" + uname + "' , '" + email + "' , '" + fname + "' , '" + lname + "' , '" +
		dob + "' , '" + doj + "' , " + cmtyid + ", '" + ppic_url + "' , '" + pvideo_url + "' ,'" + stoken +
		"' , '" + timeNow + "', TRUE , FALSE, '1900-01-01 12:00:00', " + ltype + ", " + catid + " );"
	fmt.Println("Executing : ", query)
	fmt.Println("Preparing to insert new details into the Database")
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Error in execution of insert query\n", err.Error())
		views.ShowSuccessOrErrorAsJSON(ctx, "InsertQueryExecutionError", "Error in executing the insert query", 108, 8)
		return
	}
	stmt.Exec()
	fmt.Println("New details inserted into Database(A new user login)")

	rows, err = db.Query("select id from auth_users" +
		" where email='" + email + "' ;") //I need to check about stoken
	if err != nil {
		fmt.Println("Error in execution of select query\n", err.Error())
		views.ShowSuccessOrErrorAsJSON(ctx, "SelectQueryExecutionError", "Error in execution of select query", 109, 9)
		return
	}
	fmt.Println("Details extracted from DB")
	var uid int
	stoken = middlewares.GetStoken(uid, AuthenticationKey)
	rows.Next()
	rows.Scan(&uid)

	query = "update auth_users set token='" + stoken + "' , token_updated='" +
		timeNow + "' where email='" + email + "';"
	fmt.Println("Executing : ", query)
	stmt, err = db.Prepare(query)
	if err != nil {
		fmt.Println("Error in execution of update query")
		views.ShowSuccessOrErrorAsJSON(ctx, "UpdateQueryExecutionError", "Error in execution of update query", 110, 10)
		return
	}
	stmt.Exec()
	fmt.Println("New user details inserted")
	fmt.Fprintf(ctx, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "uid", "stoken"}, 1, "You successfully logged in", uid, stoken))
}
