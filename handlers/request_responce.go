package handlers

type RequestBody struct {
	Name string `json:"name"`
}
type ResponseBody struct {
	Code       int   `json:"code"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}
