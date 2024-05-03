package main

import (
	"flag"
	"fmt"
	"github.com/sleeyax/smspool-go/sms"
	"log"
)

func main() {
	apiKey := flag.String("apiKey", "", "The API key to use.")
	flag.Parse()

	if *apiKey == "" {
		log.Fatalln("The API key flag '-apiKey' is required.")
	}

	pool := sms.Create(*apiKey)
	order, err := pool.OrderSMS(sms.OrderSMSRequest{
		Country: "US",
		Service: "1",
	})
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to order SMS: %v", err))
	}

	if !order.IsSuccess() {
		log.Fatalln(fmt.Sprintf("Order SMS not successful: %s", order.Errors))
	}

	log.Println("Successfully ordered SMS.")
	log.Println("Order ID: ", order.OrderID)
	log.Println("Phone Number: ", order.PhoneNumber)

	log.Println("Cancelling all orders...")

	cancelled, err := pool.CancelAllSMS()
	if err != nil {
		log.Fatalln(err)
	}
	if !cancelled.IsSuccess() {
		log.Fatalln(cancelled.Message)
	}
	log.Println("Successfully cancelled all SMS.", cancelled.Message)
	log.Println("Refunded orders: ", cancelled.RefundedOrders)
}
