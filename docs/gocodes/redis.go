package main
import "gopkg.in/redis.v5"
import "fmt"

func ExampleNewClient() {
    client := redis.NewClient(&redis.Options{   //Type of *redis.Client
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    fmt.Printf("Type of client variable:%T\n",client)
    pong, err := client.Ping().Result()
    fmt.Printf("Returned value : %v , Error : %v\n",pong, err)
    // Output: PONG <nil>
// }

// func ExampleClient() {
    err = client.Set("key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err2 := client.Get("quote").Result()
    if err2!= nil {
        panic(err2)
    }
    fmt.Println("\nquote : ", val)

    val2, err3 := client.Get("counter : ").Result()
    if err3 == redis.Nil {
        fmt.Println("\nkey named counter does not exists")
    } else if err3 != nil {
        panic(err3)
    } else {
        fmt.Println("\ncounter : ", val2)
    }

    v,e:=client.Get("first_name").Result();
    if e!=nil{
        panic(e)
    }
    fmt.Println("My first name : ",v)

    v,e =client.Get("mysql_port").Result();
    if e!=nil{
        panic(e)
    }
    fmt.Println("Default port of MySQL : ", v)
    // Output: key value
    // key2 does not exists
}

func main(){
    ExampleNewClient();
    //ExampleClient()
}