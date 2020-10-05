package exception

type ApiException struct {
	Code int
	Message string
	Data interface{}
}

func (apiException *ApiException) Error() string {
	return apiException.Message
}
