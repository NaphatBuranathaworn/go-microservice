package counter

type AddRequest struct {
	V int `json:"value"`
}

type AddResponse struct {
	V int `json:"value"`
}