package endpoints

import "github.com/NaphatBuranathaworn/go-microservice/internal"

type GetRequest struct {
    Filters []internal.Filter `json:"filters,omitempty"`
}

type GetResponse struct {
    Documents []internal.Document `json:"documents"`
	Err       string              `json:"err,omitempty"`
}

type AddRequest struct {
	Document *internal.Document `json:"document"`
}

type AddResponse struct {
	TicketID string `json:"ticketID"`
	Err      string `json:"err"`
}