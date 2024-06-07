package models

type Response struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

type BadResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type ResponseMany struct {
	Code  int         `json:"code"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
	Prev  string      `json:"prev"`
	Next  string      `json:"next"`
}

type LoginResponse struct {
	User
	AccessToken string `json:"accessToken"`
}

type FilterFetch struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
	Count  bool
}
