package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"strings"
	//"time"
)
import (
	"synkkuapi/conf"
	"synkkuapi/controllers/middlewares"
	"synkkuapi/views"
)

func GetMediaFiles(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
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
	pId := strings.TrimSpace(string(ctx.FormValue("pid")))

	/****************************************************************************************/
	fmt.Println("Extracted URL Parameters : ", uId, pId)
	if !middlewares.AreAllFieldsAreInValidForm(uId, pId) {
		views.ShowSuccessOrErrorAsJSON(ctx, "EmptyFieldOrKeyNameError", "You have to provide all fields (uid,pid except uid or check if  you have provided a wrong key uid", 103, 3)
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

	rows, e := db.Query("select post_id,user_id from media where post_id = " + pId + " AND user_id = " + uId + ";")
	//fmt.Println("Executing : " + rows)
	if e != nil {
		panic(e)
	} else {
		fmt.Println("All the details of users successfully retreived from db")
	}
	found := false
	for rows.Next() {
		found = true
		break
	}

	if found == true {
		q := "select id,media_type,media_url,status,created_on,updated_on from media where user_id=" + uId + " AND post_id=" + pId + " AND status=1;"

		rows2, e2 := db.Query(q)
		fmt.Println("Executing : " + q)
		if e2 != nil {

			fmt.Println("error in query")
		} else {
			fmt.Println("All the details of users successfully retreived from db")

		}
		fmt.Println("ROWS DATA :\n", rows2)

		type MediaData struct {
			Mid   int    `json:"mid"`
			Mtype int    `json:"mtype"`
			Murl  string `json:"murl"`
			Sts   int    `json:"sts"`
			Cat   string `json:"c_at"`
			Uat   string `json:"u_at"`
		}
		MediaArr := []MediaData{}

		var id int
		var media_type int
		var media_url string
		var status int
		var created_on string
		var updated_on string

		fmt.Println("Getting media details from DB...")
		for rows2.Next() {
			fmt.Println("data : ", rows2)
			//cretaed_on updated_on is not coming from database ,we have to check it
			rows2.Scan(&id, &media_type, &media_url, &status, &created_on, &updated_on)
			MediaArr = append(MediaArr, MediaData{id, media_type, media_url, status, created_on, updated_on})
		}
		fmt.Fprintf(ctx, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "me_arr"}, 1, "Media files are here", MediaArr))
	} else {
		fmt.Println(ctx, "MediaNotFound", "This media id does not exist")
		//This check does not require...it's for to make sure for proper data input while testing the API
		views.ShowSuccessOrErrorAsJSON(ctx, "MediaNotFound", "The media id related to the specified user does not exist in the database", 132, 32)
		return
	}
}
