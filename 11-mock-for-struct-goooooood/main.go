package main

import (
	"fmt"
)

//https://tutorialedge.net/golang/improving-your-tests-with-testify-go/

// MessageService handles notifying clients they have
// been charged
type MessageService interface {
	SendChargeNotification(int) error
}

// SMSService is our implementation of MessageService
type SMSService struct{}

// MyService uses the MessageService to notify clients
type MyService struct {
	messageService MessageService
}

// SendChargeNotification notifies clients they have been
// charged via SMS
// This is the method we are going to mock
func (sms SMSService) SendChargeNotification(value int) error {
	fmt.Println("Sending Production Charge Notification")
	return nil
}

// ChargeCustomer performs the charge to the customer
// In a real system we would maybe mock this as well
// but here, I want to make some money every time I run my tests
func (a MyService) ChargeCustomer(value int) error {
	a.messageService.SendChargeNotification(value)
	fmt.Printf("Charging Customer For the value of %d\n", value)
	return nil
}

// A "Production" Example
func main() {

	smsService := SMSService{}
	myService := MyService{smsService}
	myService.ChargeCustomer(100)

}
