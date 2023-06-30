package model

type Response struct {
	Data    interface{} `json:"data"`
	Page    *Page       `json:"page"`
	Success bool        `json:"success"`
	Error   string      `json:"error"`
}

type Page struct {
	PageIndex int  `json:"pageIndex"`
	PageSize  int  `json:"pageSize"`
	Total     uint `json:"total"`
}
