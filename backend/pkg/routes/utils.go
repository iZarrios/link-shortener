package routes

import "math/rand"

type ApiResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func RandomURL(size int) string {
	var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	str := make([]rune, size)

	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}

	return string(str)
}
