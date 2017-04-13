package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"strings"
	//"time"
	//"strconv"
)
import (
	"synkkuapi/conf"
	"synkkuapi/controllers/middlewares"
	"synkkuapi/views"
)

func GetCommunities(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
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

	/****************************************************************************************/
	fmt.Println("Extracted URL Parameters : ", uId)
	if !middlewares.AreAllFieldsAreInValidForm(uId) {
		views.ShowSuccessOrErrorAsJSON(ctx, "EmptyFieldOrKeyNameError", "You have to provide all fields (uid) except uid or check if  you have provided a wrong key uid", 103, 3)
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

	rows, e := db.Query("select id from auth_users where id=" + uId + ";")
	if e != nil {
		panic(e)
	} else {
		fmt.Println("userId and postId details successfully retreived from DB....")
	}

	found := false
	for rows.Next() {
		found = true
	}

	if found == false {
		views.ShowSuccessOrErrorAsJSON(ctx, "UserDoesNotExist", "This id does not exist in the database", 101, 1)
		return
	}
	rows, e = db.Query("select id, name from community;")
	if e != nil {
		views.ShowSuccessOrErrorAsJSON(ctx, "QueryExecutionError", "Error in executing select query for community", 127, 27)
		return
	} else {
		fmt.Println("userId successfully retreived from DB....")
	}
	/*----------------------------------*/
	type CommunityDetail struct {
		ComId   int    `json:"cmtyid"`
		ComName string `json:"cmtyn"`
	}
	CommArr := []CommunityDetail{}
	var commId int
	var commName string
	/*----------------------------------*/
	found = false
	for rows.Next() {
		rows.Scan(&commId, &commName)
		CommArr = append(CommArr, CommunityDetail{commId, commName})
		found = true
	}
	if found == false {
		fmt.Println("No community found")
		views.ShowSuccessOrErrorAsJSON(ctx, "NoCommunityFound", "There is no any community", 130, 30)
		return
	}
	fmt.Println("Successfully got the community list")
	fmt.Fprintf(ctx, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "cmty_arr"}, 1, "List of communities", CommArr))
}
