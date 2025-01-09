package main

import "fmt"

type paymenter interface{
    pay(amount float32)
}

type payment struct{
    gateway paymenter
}

func (p payment) makePayment (amount float32){
    // razorPaymentGw := razorPay{}
    // razorPaymentGw.pay(amount)
    p.gateway.pay(amount)
}

type razorPay struct{}

func (r razorPay) pay (amount float32){
    // Logic to make payment
    fmt.Println("Making payment using RazorPay", amount)
}

type stripe struct{}

func (s stripe) pay  (amount float32){
    // Logic to make payment
    fmt.Println("Making payment using Stripe", amount)
}

func main()  {
    // stripePaymentGw := stripe{}
    razorPaymentGw := razorPay{}
    newPayment := payment{
        gateway: razorPaymentGw,
    }
    newPayment.makePayment(100)
}