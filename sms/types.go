package sms

import (
	"github.com/sleeyax/smspool-go/internal"
)

type OrderSMSRequest struct {
	// The country in ISO format or number retrieved from the Country List endpoint
	Country string `form:"country"`
	// The service in ISO format or number retrieved from the Service List endpoint
	Service string `form:"service"`
	// The pool or number retrieved from the Pool List endpoint (optional)
	Pool string `form:"pool"`
	// The max price you are willing to pay per number
	MaxPrice string `form:"max_price"`
	// The pricing option to use
	PricingOption string `form:"pricing_option"`
	// Quantity of numbers ordered
	Quantity string `form:"quantity"`
	// The area code to use
	AreaCode string `form:"areacode"`
	// The exclude option to use
	Exclude string `form:"exclude"`
	// The create token option to use
	CreateToken string `form:"create_token"`
}

type OrderSMSResponse struct {
	internal.Response `json:",inline"`
	Pools             map[string]Pool `json:"pools"`
	Errors            []Error         `json:"errors"`
	// Error type.
	Type        string  `json:"type"`
	CC          string  `json:"cc"`
	PhoneNumber string  `json:"phonenumber"`
	Number      int     `json:"number"`
	OrderID     string  `json:"order_id"`
	Country     string  `json:"country"`
	Service     string  `json:"service"`
	Pool        int     `json:"pool"`
	ExpiresIn   int     `json:"expires_in"`
	Expiration  int     `json:"expiration"`
	Cost        string  `json:"cost"`
	CostInCents float64 `json:"cost_in_cents"`
	URL         string  `json:"url"`
	Token       string  `json:"token"`
}

type Error struct {
	Message string `json:"message"`
	Param   string `json:"param"`
}

type Pool struct {
	internal.Response `json:",inline"`
	Errors            []Error `json:"errors"`
	Type              string  `json:"type"`
}

type OrderIDRequest struct {
	OrderID string `form:"orderid"`
}

type OrderStatusCode int

const (
	OrderStatusCodePending  OrderStatusCode = 1
	OrderStatusCodeComplete                 = 3
	OrderStatusCodeRefunded                 = 6
)

type OrderStatus string

const (
	OrderStatusPending  OrderStatus = "pending"
	OrderStatusComplete             = "complete"
	OrderStatusRefunded             = "refunded"
)

type CheckSMSResponse struct {
	Status     OrderStatusCode `json:"status"`
	Message    string          `json:"message"`
	SMS        string          `json:"sms"`
	FullSMS    string          `json:"full_sms"`
	Resend     int             `json:"resend"`
	Expiration int             `json:"expiration"`
	TimeLeft   int             `json:"time_left"`
}

type CancelSMSResponse struct {
	internal.Response `json:",inline"`
}

type CancelAllSMSResponse struct {
	internal.Response `json:",inline"`
	RefundedOrders    []string `json:"refunded_orders"`
}

type ActivateSMSResponse struct {
	internal.Response `json:",inline"`
}

type ReactivateSMSResponse struct {
	internal.Response `json:",inline"`
}

type ActiveOrdersResponse []Order

type Order struct {
	// The cost of the SMS request
	Cost string `json:"cost"`
	// The order code associated with the SMS request.
	OrderCode string `json:"order_code"`
	// The phone number to which the SMS was sent.
	PhoneNumber string `json:"phonenumber"`
	// The code associated with the SMS request.
	Code string `json:"code"`
	// The full code associated with the SMS request.
	FullCode string `json:"full_code"`
	// The short name associated with the SMS request.
	ShortName string `json:"short_name"`
	// The service associated with the SMS request.
	Service string `json:"service"`
	//  The status of the SMS request.
	Status OrderStatus `json:"status"`
	// The pool associated with the SMS request.
	Pool int `json:"pool"`
	// The timestamp of the SMS request.
	Timestamp string `json:"timestamp"`
	// The completion date of the SMS request.
	CompletedOn string `json:"completed_on"`
	// The expiry time of the SMS request.
	Expiry int `json:"expiry"`
	// The time left for the SMS request to expire.
	TimeLeft int `json:"time_left"`
}

type ArchiveAllOrdersResponse struct {
	internal.Response `json:",inline"`
}

type CheckResendResponse struct {
	internal.Response `json:",inline"`
	Resends           int     `json:"resends"`
	ResendCost        float64 `json:"resendCost"`
	ExpiresInHour     int     `json:"expires_in_hour"`
}

type ResendResponse struct {
	internal.Response `json:",inline"`
	Errors            []Error `json:"errors"`
	OrderID           string  `json:"order_id"`
	Charge            int     `json:"charge"`
}

type OneTimeSMSStockRequest struct {
	// Country retrieved from "Country list" endpoint.
	Country string `form:"country"`
	// Service retrieved from "Service list" endpoint.
	Service string `form:"service"`
	// Pool retrieved from "Pool list" endpoint.
	Pool string `form:"pool"`
}

type OneTimeSMSStockResponse struct {
	Success int `json:"success"`
	Amount  int `json:"amount"`
}

type OrderHistoryRequest struct {
	// Which row you'd like to start from.
	Start int `form:"start"`
	// Max length of all rows.
	Length int `form:"length"`
	// Search query for phone number, order_code, service or country.
	Search string `form:"search"`
}

type OrderHistoryResponse []Order
