package response

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	statusOk    = "ok"
	statusError = "error"
)

func OK() Response {
	return Response{
		Status: statusOk,
	}
}

func Error(msg string) Response {
	return Response{
		Status: statusError,
		Error:  msg,
	}
}
