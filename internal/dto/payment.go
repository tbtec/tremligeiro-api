package dto

type PaymentWebHook struct {
	Resource  string `json:"resource"`
	Topic     string `json:"topic"`
	Status    string `json:"status"`
	PaymentId string `json:"paymentId"`
}

type Payment struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
