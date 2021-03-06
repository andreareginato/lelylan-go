package main

import (
	"os"
	"github.com/bradrydzewski/go.auth/oauth2"
	"github.com/joho/godotenv"
	"github.com/apellizzn/lelylan-go/lelylan"
)

type Device struct {
	Id string
	Name string
	Type string
}

func main(){

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	oauthClient := oauth2.Client {
		ClientId: os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		AccessTokenURL: "http://people.lelylan.com/oauth/token"}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	token, err := oauthClient.GrantTokenPassword(username, password, "resources")
	
	if err == nil{
		client := lelylan.Client{ Token: token }

		//get all devices
		client.Devices() 

		// create a new device
		device := client.CreateDevice("New Device", "518be84900045e1521000007")

		// provide device id
		client.Device(device.Id)
		
		// provide device new name
		client.UpdateDevice("Old Device", device.Id)

		// delete a device
		client.DeleteDevice(device.Id)
	}
}

