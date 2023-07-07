package batainer

type ServiceError struct {
	Code    int    `json:"code"`
	Message string `json:"Message"`
}
