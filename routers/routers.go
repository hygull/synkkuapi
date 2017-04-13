/*
	project started on : 18/10/2016.

*/
package routers

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)
import (
	"synkkuapi/conf"
	"synkkuapi/controllers"
)

/*@***************** Routes to Handlers ****************************************************************@*/

func RequestRouters() error {

	//To define a fasthttp router...
	fastRouter := fasthttprouter.New()

	//To show the accessible endpoints (Temporary as it will be replaced soon with other contents)...
	fastRouter.POST("/", controllers.Home)

	//To show the available & accessible endpoints...
	fastRouter.POST("/v1/endpoints", controllers.Endpoints)

	//To login into account...
	fastRouter.POST("/v1/account/login", controllers.AccountLogin)

	//To create an  account...
	fastRouter.POST("/v1/account/new", controllers.AccountNew)

	//To getposts
	fastRouter.POST("/v1/getpost/list", controllers.GetPosts) //DONE

	//To getcomments
	fastRouter.POST("/v1/getcomments/list", controllers.GetComments) //DONE

	//TO getlikes
	fastRouter.POST("/v1/getlikes", controllers.GetLikes) //DONE

	//To getfriendlist

	//fastRouter.POST("/v1/getfriend/list", controllers.GetListOfFriendRequest)

	//To getmedia files'

	fastRouter.POST("/v1/getmedia/files", controllers.GetMediaFiles) //DONE

	//To getuser profile
	fastRouter.POST("/v1/getuser/profile", controllers.GetUserProfile) //DONE

	//To getcommunities
	fastRouter.POST("/v1/get/communities", controllers.GetCommunities) //DONE

	//To get a send request
	fastRouter.POST("/v1/send/request", controllers.SendRequest) //DONE

	//Accept friend request
	fastRouter.POST("/v1/accept/friend/request", controllers.AcceptFriendRequest)

	/*@*************************** Ensure and Serve the requests **************************************@*/

	//To Serve requests...
	err := fasthttp.ListenAndServe(conf.HostServerIP+":"+conf.Port, fastRouter.Handler)

	if err != nil {
		fmt.Println("[synkku:error] Invalid URL Requested, check for proper port and host\n")
		log.Fatal(err)
		return err //If there is any error in connection...
	} else {
		return nil //If everything is fine...
	}
}
