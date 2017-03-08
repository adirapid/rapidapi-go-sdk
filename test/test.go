package main

import (
	"fmt"
	"os"
	"../RapidAPISDK"
)

func handleResponse(response map[string]interface{}){
	if response["success"] != nil {
		fmt.Println(response["success"])
	} else {
		fmt.Println(response["error"])
	}
}

func TestPublicPack(rapidApi RapidAPISDK.RapidAPI) {
	params := map[string]RapidAPISDK.Param{
		"apiKey": {"data","AIzaSyCDogEcpeA84USVXMS471PDt3zsG-caYDM"},
		"string": {"data", "test"},
		"targetLanguage": {"data", "he"},
		"sourceLanguage": {"data",""},
	}
	response := rapidApi.Call("GoogleTranslate", "translate", params)
	handleResponse(response)
}

func TestPackWithImage(rapidApi RapidAPISDK.RapidAPI) {
	params :=  map[string]RapidAPISDK.Param{
		"subscriptionKey": {"data", "57e9164516844d99ae455a9953aca0c2"},
		"image" : {"file","test/cute_dog.jpg" },
		"details": {"data", ""},
		"visualFeatures": {"data",""},
	}
	response := rapidApi.Call("MicrosoftComputerVision", "analyzeImage", params)
	handleResponse(response)
}

func TestPackWithWriter(rapidApi RapidAPISDK.RapidAPI) {

	file, err := os.Open("test/cute_dog.jpg")
	if err != nil {
		panic(err)
	}
	params := map[string]RapidAPISDK.Param{
		"subscriptionKey": {"data", "57e9164516844d99ae455a9953aca0c2"},
		"image":           {"writer", file},
		"details":         {"data", ""},
		"visualFeatures":  {"data", ""},
	}
	defer file.Close()

	response := rapidApi.Call("MicrosoftComputerVision", "analyzeImage", params)
	handleResponse(response)
}

func TestPackWithURL(rapidApi RapidAPISDK.RapidAPI) {

	params := map[string]RapidAPISDK.Param{
		"subscriptionKey": {"data", "57e9164516844d99ae455a9953aca0c2"},
		"image":           {"data", "https://i.ytimg.com/vi/opKg3fyqWt4/hqdefault.jpg"},
		"details":         {"data", ""},
		"visualFeatures":  {"data", ""},
	}

	response := rapidApi.Call("MicrosoftComputerVision", "analyzeImage", params)
	handleResponse(response)
}

func TestWebhookEvents(rapidApi RapidAPISDK.RapidAPI) {
	params := map[string]string{
		"token":  "dxJoEmQ93TypcotygRB0eok2",
		"command":  "/kaki3",
	}
	callbacks := make(map[string]func(msg interface{}))
	callbacks["onMessage"] = func(msg interface{}) {
		fmt.Println(msg)
	}
	callbacks["onJoin"] = func(msg interface{}) {
		fmt.Println("Joined!")
	}
	callbacks["onError"] = func(msg interface{}) {
		fmt.Println("Error")
	}
	rapidApi.Listen("Slack", "slashCommand", params, callbacks)
}

func main() {
	rapidApi := RapidAPISDK.RapidAPI{"withoutImage", "72352b8b-9384-4a9a-abb1-195d5e234418"}

	TestPublicPack(rapidApi)
	TestPackWithImage(rapidApi)
	TestPackWithURL(rapidApi)
	TestPackWithWriter(rapidApi)
/*
	webhooksRapid := RapidAPISDK.RapidAPI{"testa", "a3787239-bb1e-4fa4-85c3-a423fa6af51f"}
	TestWebhookEvents(webhooksRapid)
*/
}
