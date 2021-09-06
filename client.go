package main

import(
	_ "context"
	"crypto/tls"
	"fmt"
	_ "log"
	"bytes"
	"encoding/json"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	_ "google.golang.org/grpc/examples/data"
	_ "google.golang.org/grpc/examples/features/proto/echo"
	"net/http"
	_"io/ioutil"
)

func post_to_api() map[string]interface{} {
	token_request, _ := json.Marshal(map[string]string{"username":"defender_admin2", "password":"123",})
	requestBody		 := bytes.NewBuffer(token_request)
	response, connerror := http.Post("https://api.coredge.ai/defender_test_tenant_2/api/login", "application/json", requestBody)
	var respstring map[string]interface{}
	if(connerror!= nil){
		fmt.Println("this is an error")
	} /*else {		
		theresp, _ := ioutil.ReadAll(response.Body)
		var respstring string = string(theresp)
		fmt.Println(respstring)
		
	}*/
	json.NewDecoder(response.Body).Decode(&respstring)
	defer response.Body.Close()
	
	return respstring
}

func fetchToken() *oauth2.Token {
	access_token := post_to_api()["accessToken"].(string)
	if(access_token != "nil"){
		fmt.Println("there is a token")
		return &oauth2.Token{
			AccessToken: "0",
		}
	} else{
		return nil	
	}		
}

func main(){

	rpccred := oauth.NewOauthAccess(fetchToken())
	creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})
	transportSecurity := grpc.WithTransportCredentials(creds)
		
	conn, connerror := grpc.Dial("data.coredge.ai:443", transportSecurity, grpc.WithPerRPCCredentials(rpccred))
	if(connerror!= nil){
		fmt.Println("this is an error ")
		fmt.Print("%d", connerror)
	} else {		
		fmt.Printf("%+v", conn)
	}
	defer conn.Close()	
}	