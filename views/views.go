/*      ############  JSON  VIEW  ###################

Aim 	   : To display the JSON data to the browser
Created on : 18/10/2016
*/

package views

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

/************* Construction of JSON messages for success/error ***********/

func ShowSuccessOrErrorAsJSON(ctx *fasthttp.RequestCtx, errType string, errDetails string, statusCode int, id int) {
	var jsonStrData []byte

	type ErrMsg struct {
		Message    string `json:"message"`
		Statuscode int    `json:"status code"`
		Id         int    `json:"id"`
	}
	type ErrMsgs struct {
		MessageDetails ErrMsg `json:"message details"`
	}
	msgsCollection := map[string]ErrMsgs{}
	msgsCollection[errType] = ErrMsgs{ErrMsg{errDetails, statusCode, id}}
	jsonStrData, _ = json.Marshal(msgsCollection)
	fmt.Fprintf(ctx, string(jsonStrData))
}

/************* To display JSON output {For testing the API}*************************************/

func ShowDataAsJSON(ctx *fasthttp.RequestCtx, jsonStrData string) {
	fmt.Fprintf(ctx, jsonStrData)
}

/***************************** Universal JSON Message creator ****************************/
/*                                  Very IMP for this API                                */
/*     s:=A([]string{"Name","Age","Guide"},"Rishikesh",25,"Rathnakara Sir & Arjun Sir")  */

func UniversalMessageCreatorAsJSON(keys []string, values ...interface{}) string {
	jsonMapData := map[string]interface{}{}
	i := 0
	for _, v := range values {
		jsonMapData[keys[i]] = v
		i += 1
	}
	var jsonStrData []byte
	jsonStrData, _ = json.Marshal(jsonMapData)
	return string(jsonStrData)
}
