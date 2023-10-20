package main

import "fmt"

// Before SRP
// type Order struct {
// 	items []string
// }

// func (o *Order) addItem(item string) {
// 	o.items = append(o.items, item)
// }

// func (o *Order) calculateTotal() float64 {
// 	total := 0.0
// 	for _, item := range o.items {
// 		total += getItemPrice(item)
// 	}
// 	return total
// }

// func getItemPrice(item string) float64 {
// 	return 10.0
// }

// func main() {
// 	order := Order{}
// 	order.addItem("shoes")
// 	order.addItem("bag")
// 	total := order.calculateTotal()
// 	fmt.Printf("Total is $%.2f\n", total)
// }

// After SRP
// type Order struct {
// 	items []string
// }

// func (o *Order) addItem(item string) {
// 	o.items = append(o.items, item)
// }

// func CalculateTotal(items []string) float64 {
// 	total := 0.0
// 	for _, item := range items {
// 		total += getItemPrice(item)
// 	}
// 	return total
// }

// func getItemPrice(item string) float64 {
// 	return 10.0
// }

// func main() {
// 	order := Order{}
// 	order.addItem("shoes")
// 	order.addItem("bag")
// 	total := CalculateTotal(order.items)
// 	fmt.Printf("Total is $%.2f\n", total)
// }

//Before OCP

// type Circle struct {
// 	Radius float64
// }

// type Square struct {
// 	SideLength float64
// }

// func (c *Circle) Area() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

// func (s *Square) Area() float64 {
// 	return s.SideLength * s.SideLength
// }

// func main() {
// 	circle := &Circle{Radius: 5}
// 	square := &Square{SideLength: 5}
// 	fmt.Printf("Circle Area: %f\n", circle.Area())
// 	fmt.Printf("Square Area: %f\n", square.Area())
// }

// After OCP

// type Circle struct {
// 	Radius float64
// }

// type Square struct {
// 	SideLength float64
// }

// func (c *Circle) Area() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

// func (s *Square) Area() float64 {
// 	return s.SideLength * s.SideLength
// }

// type Shape interface {
// 	Area() float64
// }

// func CalculateArea(shapes Shape) float64 {
// 	return shapes.Area()
// }

// func main() {
// 	circle := Circle{Radius: 5}
// 	square := &Square{SideLength: 5}
// 	fmt.Printf("Circle Area: %f\n", CalculateArea(circle))
// 	fmt.Printf("Square Area: %f\n", CalculateArea(square))
// }

// Before LSP
// type Bird struct{}

// func (b *Bird) Eat() {
// 	fmt.Println("Bird is eating")
// }

// type Ostrich struct{}

// func (o *Ostrich) Eat() {
// 	fmt.Println("Ostrich can't eating")
// }

// func feedBird(b *Bird) {
// 	b.Eat()
// }

// func main() {
// 	bird := &Bird{}
// 	ostrich := &Ostrich{}
// 	feedBird(bird)
// 	feedBird(ostrich)
// }

//After LSP
// type Bird interface {
// 	Eat()
// }

// type Dara struct{}

// func (d *Dara) Eat() {
// 	fmt.Println("Dara is eating")
// }

// type Ostrich struct{}

// func (o *Ostrich) Eat() {
// 	fmt.Println("Ostrich can't eating")
// }

// func feedBird(b Bird) {
// 	b.Eat()
// }

// func main() {
// 	dara := &Dara{}
// 	ostrich := &Ostrich{}
// 	feedBird(dara)
// 	feedBird(ostrich)
// }

// Before ISP

// type Worker interface {
// 	Work()
// 	Eat()
// }

// type Engineer struct{}
// type Robot struct{}

// func (e *Engineer) Work() {
// 	fmt.Println("Engineer is working")
// }
// func (e *Engineer) Eat() {
// 	fmt.Println("Engineer is eating")
// }
// func (r *Robot) Work() {
// 	fmt.Println("Robot is working")
// }

// func doJob(w Worker) {
// 	w.Work()
// 	w.Eat()
// }

// func main() {
// 	Engineer := &Engineer{}
// 	robot := &Robot{}
// 	doJob(Engineer)
// 	doJob(robot)
// }

// After ISP

// type Worker interface {
// 	Work()
// }

// type Eater interface {
// 	Eat()
// }

// type Engineer struct{}
// type Robot struct{}

// func (e *Engineer) Work() {
// 	fmt.Println("Engineer is working")
// }
// func (e *Engineer) Eat() {
// 	fmt.Println("Engineer is eating")
// }
// func (r *Robot) Work() {
// 	fmt.Println("Robot is working")
// }

// func doJob(w Worker) {
// 	w.Work()
// }
// func doEat(e Eater) {
// 	e.Eat()
// }

// func main() {
// 	Engineer := &Engineer{}
// 	robot := &Robot{}
// 	doJob(Engineer)
// 	doEat(Engineer)
// 	doJob(robot)
// }

// Before DIP
// type EmailService struct{}
// type SmsService struct{}

// func (es *EmailService) SendEmail(to, subject, message string) error {
// 	fmt.Println("Sending email to", to)
// 	return nil
// }

// func (ss *SmsService) SendSms(to, message string) error {
// 	fmt.Println("Sending SMS to", to)
// 	return nil
// }

// type MessageSender struct {
// 	EmailService *EmailService
// 	SmsService   *SmsService
// }

// func (ms *MessageSender) SendMail(to, subject, message string) error {
// 	fmt.Println("Sending email message...")
// 	return ms.EmailService.SendEmail(to, subject, message)
// }

// func (ms *MessageSender) SendSms(to, message string) error {
// 	fmt.Println("Sending sms message...")
// 	return ms.SmsService.SendSms(to, message)
// }

// func main() {
// 	emailService := &EmailService{}
// 	smsService := &SmsService{}
// 	messageSender := &MessageSender{
// 		EmailService: emailService,
// 		SmsService:   smsService,
// 	}

// 	emailErr := messageSender.SendMail("recipient@mail.com", "Hai", "Ini Message")
// 	if emailErr != nil {
// 		fmt.Println("Error sending email:", emailErr)
// 	}

// 	smsErr := messageSender.SendSms("081277281092", "Ini Message")
// 	if smsErr != nil {
// 		fmt.Println("Error sending sms:", smsErr)
// 	}

// }

// After DIP

type MessageSender interface {
	Send(to, subject, message string) error
}

type EmailService struct{}
type SmsService struct{}

func (es *EmailService) Send(to, subject, message string) error {
	fmt.Println("Sending email to", to)
	return nil
}

func (ss *SmsService) Send(to, subject, message string) error {
	fmt.Println("Sending SMS to", to)
	return nil
}

type MessageService struct {
	sender MessageSender
}

func (ms *MessageService) SendMail(to, subject, message string) error {
	fmt.Println("Sending email message...")
	return ms.sender.Send(to, subject, message)
}

func (ms *MessageService) SendSms(to, subject, message string) error {
	fmt.Println("Sending sms message...")
	return ms.sender.Send(to, subject, message)
}

func main() {
	emailService := &EmailService{}
	smsService := &SmsService{}

	emailMessageService := &MessageService{sender: emailService}
	smsMessageService := &MessageService{sender: smsService}

	emailErr := emailMessageService.SendMail("recipient@mail.com", "Hai", "Ini Message")
	if emailErr != nil {
		fmt.Println("Error sending email:", emailErr)
	}

	smsErr := smsMessageService.SendSms("081277281092", "", "Ini Message")
	if smsErr != nil {
		fmt.Println("Error sending sms:", smsErr)
	}

}
