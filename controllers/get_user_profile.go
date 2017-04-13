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

func GetUserProfile(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
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
	profileUid := strings.TrimSpace(string(ctx.FormValue("profile_uid")))

	/****************************************************************************************/
	fmt.Println("Extracted URL Parameters : ", uId, profileUid)
	if !middlewares.AreAllFieldsAreInValidForm(uId, profileUid) {
		views.ShowSuccessOrErrorAsJSON(ctx, "EmptyFieldOrKeyNameError", "You have to provide all fields (uid,profile_uid) except uid or check if  you have provided a wrong key uid", 103, 3)
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

	/****************** first part checking whether user is friend or not ***************/

	query := "select to_user_id,from_user_id from friends where to_user_id=" + uId + " and from_user_id=" + profileUid + ";"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error in query")
		return
	}
	found := false
	for rows.Next() {
		found = true
	}
	if found == false {
		fmt.Println("This user is not in the friend list")
		views.ShowSuccessOrErrorAsJSON(ctx, "FriendDoesNotExist", "This UserId Is Not In The Friend List", 136, 36)
		return
	}

	/*****************collecting user details ***********************/

	query = "select user_name,email,profile_pic,profile_video,dob,community_id, " +
		" login_type,category_id,token_updated,deleted_on " +
		" from auth_users where id=" + profileUid + ";"
	fmt.Println(query)
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println("Error in execution of query")
		return
	}
	type UserArr struct {
		Un     string `json:"un"`
		Umail  string `json:"umail"`
		Ppic   string `json:"p_pic"`
		Pvid   string `json:"p_vid"`
		Dob    string `json:"dob"`
		Cmtyid int    `json:"cmtyid"`

		Ltype int `json:"ltype"`
		//Status int `json:"sts"`
		Categoryid   int    `json:"category_id"`
		Tokenupdated string `json:"token_createdat"`
		Deletedon    string `json:"deleted_on"`
	}

	var Uname string
	var Umail string
	var Ppic string
	var Pvid string
	var Dob string
	var Cmtyid int
	var Ltype int
	//var Sts int
	var Categoryid int
	var Tokenupdated string
	var Deletedon string

	rows.Next()
	rows.Scan(&Uname, &Umail, &Ppic, &Pvid, &Dob, &Cmtyid, &Ltype, &Categoryid, &Tokenupdated,
		&Deletedon)

	UserDetails := UserArr{Uname, Umail, Ppic, Pvid, Dob, Cmtyid, Ltype, Categoryid, Tokenupdated, Deletedon}

	/********************collection user address*******************/
	query = "select * from users_details where user_id=" + profileUid + ";"
	rows, err = db.Query(query)
	if err != nil {
		views.ShowSuccessOrErrorAsJSON(ctx, "QueryExecutionError", "Error in execution query", 127, 27)
		fmt.Println("Error in execution of query")
		return
	}

	type AddrArr struct {
		Id                    int    `json:"id"`
		UserId                int    `json:"user_id"`
		CurrentAddress        string `json:"current_address"`
		CurrentAddressState   int    `json:"current_address_state"`
		CurrentAddressPincode int    `json:"current_adderss_state"`
		PresentAddress        string `json:"present_address"`
		PresentAddressState   int    `json:"present_address_state"`
		PresentAddressPincode int    `json:"present_address_pincode"`
		CreatedOn             string `json:"created_on"`
		UpdatedOn             string `json:"updated_on"`
	}

	var Id int
	var Userid int
	var CurrentAddress string
	var CurrentAddressState int
	var CurrentAddressPincode int
	var PresentAddress string
	var PresentAddressState int
	var PresentAddressPincode int
	var CreatedOn string
	var UpdatedOn string

	rows.Next()
	rows.Scan(&Id, &Userid, &CurrentAddress, &CurrentAddressState, &CurrentAddressPincode, &PresentAddress,
		&PresentAddressState, &PresentAddressPincode, CreatedOn, UpdatedOn)

	AddressDetails := AddrArr{Id, Userid, CurrentAddress, CurrentAddressState, CurrentAddressPincode,
		PresentAddress, PresentAddressState, PresentAddressPincode, CreatedOn, UpdatedOn}

	/**********************************************************/

	query = "select * from users_workexperience"
	rows, err = db.Query(query)
	if err != nil {
		views.ShowSuccessOrErrorAsJSON(ctx, "QueryExecutionError", "Error in execution query", 127, 27)
		fmt.Println("Error in execution of query")
		return
	}

	type ExpArr struct {
		Id        int    `json:"id"`
		UserId    int    `json:"user_id"`
		CompanyId int    `json:"company_id"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		CreatedOn string `json:"created_on"`
		UpdatedOn string `json:"updated_on"`
	}

	var id int
	var UserId int
	var CompanyId int
	var StartDate string
	var EndDate string
	var Createdon string
	var Updatedon string

	rows.Next()
	rows.Scan(&id, &UserId, &CompanyId, &StartDate, &EndDate, &Createdon, &Updatedon)

	ExperienceDetails := ExpArr{id, UserId, CompanyId, StartDate, EndDate, Createdon, Updatedon}

	fmt.Println("User profile displayed")
	fmt.Fprintf(ctx, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "u_arr", "addr_arr", "exp_arr"}, 1,
		"User details", UserDetails, AddressDetails, ExperienceDetails))
}
