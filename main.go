package main

import(
	"fmt"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/joho/godotenv"
	"log"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := twilio.NewRestClient()
	
	params := &api.CreateMessageParams{}
	params.SetBody("My first Twilio message Using API")
	params.SetFrom("+1")
	params.SetTo("+17166179181")
	
	resp, err := client.Api.CreateMessage(params)
	if err != nil{
		fmt.Println(err)

	}else {
		if resp.Body != nil{
			fmt.Println(*resp.Body)
		}else {
			fmt.Println(resp.Body)
		}

	}
	
}