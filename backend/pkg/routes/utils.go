package routes

type ApiResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}
