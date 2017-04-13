/*
	Created on 18 Oct 2016
*/

package conf

/*@*********** Database Server settings **********************************************************@*/

//Database host eg 127.0.0.1 or localhost
var DBHost = "127.0.0.1"

//Port number for Database server, eg. for Mysql 3306, for PostgreSQl 5432/5433
var DBPort = "5433"

//The name of database
var DBName = "synkkunew" /*	"college"	*/

//The username for Database
var DBUserName = "postgres" /*	"rishikesh"	*/

//The password of your database
var DBPassword = "admin@321" /*	"rishikesh"	*/

//The AuthenticationKey
var AuthenticationKey = "eZ$21#@54>4074W8Ndkf**WE32awe2376THWEKm"

/*@************* 2nd combination for vagrant PostgreSQL Database *********************************@*/

//Database host eg 127.0.0.1 or localhost
// var DBHost = "192.168.33.10" //"127.0.0.1" //"198.168.50.4" //"10.0.2.2"

// //var DBHost = "192.168.1.10"

// //Port number for Database server, eg. for Mysql 3306, for PostgreSQl 5432/5433
// var DBPort = "2222" //"80" //"5939"

// //The name of database
// var DBName = "synkku" /*	"college" */

// //The username for Database
// var DBUserName = "rishikesh" /* "postgres" */

// //The password of your database
// var DBPassword = "rishikesh" /* "admin@321" /*

/*@*********** Server Port and IP settings ********************************************************@*/

//Default IP of localhost, you can change it to your own, once initialize it with "" and try
var HostServerIP = "127.0.0.1" /* "localhost" */

//Port Number on which server will listen for the client requests, SynnkuPort:="8000"   is illegal
var Port = "8000"
