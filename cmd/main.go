package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
)

const IP_URL = "https://api.ipify.org?format=json"

type IPAddr struct {
	IP string `json:"ip"`
}

var IPCache IPAddr

func main() {

	AWS_KEY := os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET := os.Getenv("AWS_SECRET_ACCESS_KEY")

	fmt.Printf("%s -> %s\n", AWS_KEY, AWS_SECRET)

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		log.Fatal("Error loading aws config")
	}

	fmt.Println(cfg)

	for {
		time.Sleep(5 * time.Second)

		resp, err := http.Get(IP_URL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		addr := IPAddr{}
		json.NewDecoder(resp.Body).Decode(&addr)

		if addr.IP != IPCache.IP {
			log.Printf("New IP Detected -> %q", addr.IP)
			// TODO: Logic to update route53
			IPCache.IP = addr.IP
		}
	}
}
