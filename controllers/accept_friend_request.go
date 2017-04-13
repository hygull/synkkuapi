package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"strings"
	"time"
	//"strconv"
)
import (
	"synkkuapi/conf"
	"synkkuapi/controllers/middlewares"
	"synkkuapi/views"
)

func AcceptFriendRequest(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	//AuthenticationKey := conf.AuthenticationKey
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

	uId := strings.TrimSpace(string(ctx.FormValue("uid")))
	fromUid := strings.TrimSpace(string(ctx.FormValue("fromuid")))

	/****************************************************************************************/
	fmt.Println("Extracted URL Parameters : ", uId, fromUid)
	if !middlewares.AreAllFieldsAreInValidForm(uId, fromUid) {
		views.ShowSuccessOrErrorAsJSON(ctx, "EmptyFieldOrKeyNameError", "You have to provide all fields (uid,fromuid) except uid or check if  you have provided a wrong key uid", 103, 3)
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

	var t string
	t = time.Now().String()[0:19]

	query := "update friends set accepted_on='" + t + "', status=2 where from_user_id=" + fromUid + " and to_user_id=" + uId + ";"
	row2, err2 := db.Prepare(query)
	if err2 != nil {
		views.ShowSuccessOrErrorAsJSON(ctx, "QueryExecutionError", "Error in query", 127, 27)
		fmt.Println("error in query")
		return
	}
	row2.Exec()
	fmt.Println("Successfully accepted (Request accepted)")
	views.ShowSuccessOrErrorAsJSON(ctx, "FriendRequestAccepted", "Friend request successfully accepted", 138, 38)

}
