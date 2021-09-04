package main

import(
	_ "context"
//	"flag"
	"fmt"
	_ "log"
//	"time"

	_ "golang.org/x/oauth2"
	_"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/credentials/oauth"
	_ "google.golang.org/grpc/examples/data"
	_ "google.golang.org/grpc/examples/features/proto/echo"
	"net/http"
	_"io/ioutil"
)

func main(){
	rpccred := grpc	
	/*conn, connerror := grpc.Dial("https://simple-books-api.glitch.me", grpc.WithInsecure())
	
	if(connerror!= nil){
		fmt.Println("this is an error")
	} else {		
		fmt.Println(conn)
	}
	*/
	response, connerror := http.Get("https://simple-books-api.glitch.me/books")
	if(connerror!= nil){
		fmt.Println("this is an error")
	} else {		
		theresp, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(response.Body))
	}	
	
}	
	