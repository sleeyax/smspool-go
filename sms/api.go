package sms

import (
	"encoding/json"
	"github.com/sleeyax/smspool-go/internal"
)

type Api struct {
	internal.Api
}

func Create(apiKey string) Api {
	return Api{
		Api: internal.Create(apiKey),
	}
}

// OrderSMS orders an SMS.
func (a Api) OrderSMS(request OrderSMSRequest) (OrderSMSResponse, error) {
	var t OrderSMSResponse

	body, err := a.Send(request, "/purchase/sms")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}

// CheckSMS checks the status of an SMS order.
// You may run into rate limits while using this endpoint for high rate requests, we recommend using ActiveOrders instead as it will list all your orders at the same time.
func (a Api) CheckSMS(orderID string) (CheckSMSResponse, error) {
	var t CheckSMSResponse

	body, err := a.Send(OrderIDRequest{OrderID: orderID}, "/sms/check")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}

// CancelSMS cancels an SMS order.
func (a Api) CancelSMS(orderID string) (CancelSMSResponse, error) {
	var t CancelSMSResponse

	body, err := a.Send(OrderIDRequest{OrderID: orderID}, "/sms/cancel")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}

// CancelAllSMS cancels all SMS orders.
func (a Api) CancelAllSMS() (CancelAllSMSResponse, error) {
	var t CancelAllSMSResponse

	body, err := a.Send(struct{}{}, "/sms/cancel_all")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}

// ActivateSMS activates an SMS order.
// This endpoint will forcefully activate an SMS to set it available for future resends, please keep in mind that these orders are NOT eligible for a refund after activating them.
func (a Api) ActivateSMS(orderID string) (ActivateSMSResponse, error) {
	var t ActivateSMSResponse

	body, err := a.Send(OrderIDRequest{OrderID: orderID}, "/sms/activate")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}

// ReActivateSMS reactivates an SMS order.
// This endpoint will reactivate an order, please keep in mind that these orders are NOT eligible for a refund after activating them.
func (a Api) ReActivateSMS(orderID string) (ReactivateSMSResponse, error) {
	var t ReactivateSMSResponse

	body, err := a.Send(OrderIDRequest{OrderID: orderID}, "/sms/reactivate")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}

// ActiveOrders lists all active orders.
func (a Api) ActiveOrders() (ActiveOrdersResponse, error) {
	var t ActiveOrdersResponse

	body, err := a.Send(struct {
	}{}, "/request/active")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}

// ArchiveAllOrders archives all orders.
// This API endpoint is used to archive a request in the SMS Pool system.
func (a Api) ArchiveAllOrders() (ArchiveAllOrdersResponse, error) {
	var t ArchiveAllOrdersResponse

	body, err := a.Send(struct {
	}{}, "/request/archive")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}

// CheckResend checks if an SMS order can be resent.
func (a Api) CheckResend(orderID string) (CheckResendResponse, error) {
	var t CheckResendResponse

	body, err := a.Send(OrderIDRequest{OrderID: orderID}, "/sms/check_resend")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}

// Resend resends an SMS order.
// Keep in mind that some pools will have a charge per resend.
func (a Api) Resend(orderID string) (ResendResponse, error) {
	var t ResendResponse

	body, err := a.Send(OrderIDRequest{OrderID: orderID}, "/sms/resend")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}

// OneTimeSMSStock retrieves the stock of a one-time SMS.
func (a Api) OneTimeSMSStock(request OneTimeSMSStockRequest) (OneTimeSMSStockResponse, error) {
	var t OneTimeSMSStockResponse

	body, err := a.Send(request, "/sms/stock")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}

// OrderHistory retrieves the order history of SMS requests.
func (a Api) OrderHistory(request OrderHistoryRequest) (OrderHistoryResponse, error) {
	var t OrderHistoryResponse

	body, err := a.Send(request, "/request/history")
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)

	return t, err
}
