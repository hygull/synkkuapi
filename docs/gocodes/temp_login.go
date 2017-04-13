package controllers

import (
	// "crypto/md5"
	// "database/sql"
	// "encoding/hex"
	// // _ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
       "fmt"
	// // "net/http"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	// "regexp"
	// "strconv"
	// "strings"
	// "time"
)
// import (
// 	"synkku/conf"
// 	"synkku/controllers/validation"
// 	"synkku/views"
// )

// /********************** Middleware (Token validator) ***********************/
// func isAppSignInKeyCorrect(appSignInKeyMd5 string, authenticationToken string, email string) bool {
// 	reversed_email := []byte(email) //string is immutable in Go
// 	//email reversing for loop
// 	for front, last := 0, len(reversed_email)-1; front < last; front, last = front+1, last-1 {
// 		reversed_email[front], reversed_email[last] = reversed_email[last], reversed_email[front]
// 	}
// 	reversed_email_and_auth_token := string(reversed_email) + authenticationToken

// 	fmt.Printf("Reversed email and auth token      :  %v (%T)", reversed_email_and_auth_token, reversed_email_and_auth_token)
// 	fmt.Println()
// 	fmt.Println("\nYour app generated app_sign_key : ", appSignInKeyMd5)

// 	md5_of_reversed_email_and_auth_token := getMD5Hash(reversed_email_and_auth_token)
// 	fmt.Println("Server generated app_sign_key   : ", md5_of_reversed_email_and_auth_token)

// 	return md5_of_reversed_email_and_auth_token == appSignInKeyMd5
// }

// /*************************START (User login callee's)******************************************/

// func getStoken(userId int, authenticationToken string) string {
// 	unixTime := strconv.FormatInt(time.Now().UTC().Unix(), 10)
// 	tempStoken := strconv.Itoa(userId) + authenticationToken + unixTime //<userId><authenticationToken><unixTime>
// 	stoken := getMD5Hash(tempStoken)                                    //md5("<userId><authenticationToken><unixTime>")
// 	return stoken
// }

// func getMD5Hash(text string) string {
// 	hasher := md5.New()
// 	hasher.Write([]byte(text))
// 	return hex.EncodeToString(hasher.Sum(nil))
// }

// func IsEmailValid(email string) bool {
// 	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
// 	return Re.MatchString(email)
// }

/*************************END of callee of UserLogin2 (User login)***********************/

/********************** User Login ******************************************************/
func UserLogin(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	const AuthenticationKey = "eZ$21#@54>4074W8Ndkf**WE32awe2376THWEKm"
	ctx.SetContentType("text/plain; charset=utf8")

	if ctx.IsPost(){ //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "PostMethodNotFoundError", "The data is not being sent using POST method", 107, 7)
		fmt.Println("The data is not being sent using POST method")
		return
	}
	// req.ParseForm()

	// if len(req.Form) == 0 { //If there's no post data then inform the user
	// 	views.ShowSuccessOrErrorAsJSON(rw, "NoPostDataError", "You haven't sent the POST data", 104, 4)
	// 	fmt.Println("You haven't send the POST data")
	// 	return
	// }
	/***************************Getting data from URLs *************************************/
	un := ""
	// p_pic := ""
	// cmtyid := "0"
	// atype := "0"
	// ltype := "0"

	if len(ctx.FormValue("name"))!= "" {
		// un = strings.TrimSpace(req.Form["un"][0])
		un=strings.TrimSpace(strings(ctx.FormValue("name")))
	}
	fmt.Println(un)
	// if req.Form["p_pic"][0] != "" {
	// 	p_pic = strings.TrimSpace(req.Form["p_pic"][0]) //From google  (optional)
	// }

	// if req.Form["cmtyid"][0] != "" {
	// 	cmtyid = strings.TrimSpace(req.Form["cmtyid"][0]) //community id (required)
	// }
	// if req.Form["ltype"][0] != "" {
	// 	ltype = strings.TrimSpace(req.Form["ltype"][0]) //1 for google (required)
	// }
	// if req.Form["atype"][0] != "" {
	// 	atype = strings.TrimSpace(req.Form["atype"][0]) //1 for google (required)
	// }
	// umail := strings.TrimSpace(req.Form["umail"][0])

	// app_sign_in_key := strings.TrimSpace(req.Form["appsigninkey"][0]) //app sign in key while first login (required)
	// /****************************************************************************************/
	// fmt.Println("Extracted URL Parameters : ", un, umail, p_pic, cmtyid, ltype, atype, app_sign_in_key)
	// if !validation.AreAllFieldsAreInValidForm(umail, app_sign_in_key, cmtyid, ltype, atype) {
	// 	views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds,profile pic & username should be empty", 103, 3)
	// 	return
	// }

	// if !IsEmailValid(umail) {
	// 	views.ShowSuccessOrErrorAsJSON(rw, "InvalidEmailFormatError", "The enetered email is not valid", 133, 33)
	// 	fmt.Println("Email is not valid")
	// 	return
	// }
	// if !isAppSignInKeyCorrect(app_sign_in_key, AuthenticationKey, umail) { //Call to a Middleware
	// 	views.ShowSuccessOrErrorAsJSON(rw, "InvalidAppSignKey", "App signin key is invalid", 132, 32)
	// 	fmt.Println("The token is invalid")
	// 	return
	// }

	// fmt.Println("Connecting to DB...")
	// //db, err := sql.Open("mysql", conf.DBUserName+":"+conf.DBPassword+"@tcp"+"("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName)
	// db, err := sql.Open("mysql", conf.DBUserName+":"+conf.DBPassword+"@tcp"+"("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName)

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close()

	// fmt.Println("Connected")
	// rows, err := db.Query("select  user_name, email_id from users where email_id= '" + umail + "' AND active=1;")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// found := false
	// for rows.Next() {
	// 	found = true
	// }
	// if found {
	// 	rows, err = db.Query("select user_id, profilepic_url, profilevideo_url, community_id, account_type from users" +
	// 		" where email_id='" + umail + "' ;") //I need to check about stoken
	// 	fmt.Println("Details extracted from DB")
	// 	var ppic, pvid, stoken string
	// 	var cmtid, act, uid int
	// 	stoken = getStoken(uid, AuthenticationKey)
	// 	rows.Next()
	// 	rows.Scan(&uid, &ppic, &pvid, &cmtid, &act)

	// 	stmt, err := db.Prepare("update users set stoken='" + stoken + "' , stoken_updated_on='" + time.Now().String()[0:19] + "' where email_id='" + umail + "';")
	// 	if err != nil {
	// 		fmt.Println("Error in execution of update query")
	// 		views.ShowSuccessOrErrorAsJSON(rw, "UpdateQueryExecutionError", "Error in execution of update query", 134, 34)
	// 		return
	// 	}
	// 	stmt.Exec()
	// 	fmt.Println("Details updated as the user is already having an account")
	// 	fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "uid", "p_pic", "p_vid", "cmtyid", "atype", "stoken"}, 1, "You successfully logged in", uid, ppic, pvid, cmtid, act, stoken))
	// 	//http.Redirect(rw, req, "/", 301)
	// } else {
	// 	timeNow := time.Now().String()[0:19]

	// 	query := "insert into users(user_name, email_id, date_of_birth, city_id, gender, " +
	// 		" marital_status, studied_at, employed_at, community_id, profilepic_url, profilevideo_url, login_type, " +
	// 		"account_type,stoken , stoken_updated_on, details_updated_on, updated_on, active, " +
	// 		" deleted_on ) values(" + "'" + un + "' , '" + umail + "' , '1900-01-01',0, 0, 0, 'Default School', 'None', " +
	// 		cmtyid + ", '" + p_pic + "', '', 0, 1, '' , '1900-01-01 12:00:00', '" + timeNow + "' , '" + timeNow +
	// 		"' , 1, '1900-01-01 12:00:00');"
	// 	fmt.Println(query)
	// 	fmt.Println("Preparing to insert new details on the Database")
	// 	stmt, err := db.Prepare(query)
	// 	if err != nil {
	// 		fmt.Println("Error in query execution")
	// 		views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "Error in executing the insert query", 127, 27)
	// 		return
	// 	}
	// 	stmt.Exec()
	// 	fmt.Println("New details inserted into Database(A new user login)")

	// 	rows, err = db.Query("select user_id, profilepic_url, profilevideo_url, community_id, account_type from users" +
	// 		" where email_id='" + umail + "' ;") //I need to check about stoken
	// 	fmt.Println("Details extracted from DB")
	// 	var ppic, pvid, stoken string
	// 	var cmtid, act, uid int
	// 	stoken = getStoken(uid, AuthenticationKey)
	// 	rows.Next()
	// 	rows.Scan(&uid, &ppic, &pvid, &cmtid, &act)

	// 	stmt, err = db.Prepare("update users set stoken='" + stoken + "' , stoken_updated_on='" + time.Now().String()[0:19] + "' where email_id='" + umail + "';")
	// 	if err != nil {
	// 		fmt.Println("Error in execution of update query")
	// 		views.ShowSuccessOrErrorAsJSON(rw, "UpdateQueryExecutionError", "Error in execution of update query", 134, 34)
	// 		return
	// 	}
	// 	stmt.Exec()
		// fmt.Println("New user details updated")
		// fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "uid", "p_pic", "p_vid", "cmtyid", "atype", "stoken"}, 1, "You successfully logged in", uid, ppic, pvid, cmtid, act, stoken))
	//}
}
