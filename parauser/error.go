package parauser

type ServiceError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
