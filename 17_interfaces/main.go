package main

import "fmt"

type payment struct{}

func (p payment) makePayment (amount float32){
    // razorPaymentGw := razorPay{}
    // razorPaymentGw.pay(amount)
    razorPaymentGw := stripe{}
    razorPaymentGw.pay(amount)
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
    newPayment := payment{}
    newPayment.makePayment(100)
}