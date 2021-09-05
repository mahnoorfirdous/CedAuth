package main

import(
	_ "context"
//	"flag"
	"fmt"
	_ "log"
	"bytes"
	"encoding/json"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	_"google.golang.org/grpc/credentials"
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
			AccessToken: access_token,
		}
	} else{
		return nil	
	}		
}

func main(){

	//fmt.Println(post_to_api()["accessToken"].(string))
	rpccred := oauth.NewOauthAccess(fetchToken())
		
	conn, connerror := grpc.Dial("https://api.coredge.ai/defender_test_tenant_2/api/login", grpc.WithTransportCredentials(rpccred), grpc.WithPerRPCCredentials(credentials)) //there needs to be an auth server url
	//conn, connerror := grpc.Dial("https://api.coredge.ai/defender_test_tenant_2/api/login", grpc.WithInsecure())
	if(connerror!= nil){
		fmt.Println("this is an error ")
		fmt.Print("%d", connerror)
	} else {		
		fmt.Println(conn)
	}
	defer conn.Close()	
}	
	
	//does WithInsecure() contradict WithPerRPCCredentials()? and does oauth dialing fail because the server is not using it?
	//need to figure out TLS credentials that match with the server