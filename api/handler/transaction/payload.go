package transaction

type TransactionPayload struct {
	CustomerName string `json:"customer_name"`
	Tickets []TicketPayload `json:"tickets"`
}

type TicketPayload struct {
	TicketType string `json:"ticket"`
	Qty int `json:"qty"`
}
