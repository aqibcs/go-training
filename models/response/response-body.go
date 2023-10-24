package response

type ResponseBody struct {
	Code       int   `json:"code"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}
