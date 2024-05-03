# smspool-go

## Installation
```bash
$ go get github.com/sleeyax/smspool-go
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/sleeyax/smspool-go/sms"
	"log"
)

func main() {
	pool := sms.Create("YOUR_API_KEY")
	order, err := pool.OrderSMS(sms.OrderSMSRequest{
		Country: "US",
		Service: "example",
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
}
```
