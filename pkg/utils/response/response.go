package response

import "library-sevice/internal/dto"

type RespOk struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RespOkWithPagination struct {
	Success    bool               `json:"success"`
	Message    string             `json:"message"`
	Data       interface{}        `json:"data"`
	Pagiration dto.PaginationInfo `json:"pagination"`
}

type RespLogin struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	AccessToken string      `json:"accessToken"`
	TokenType   string      `json:"tokenType"`
	Data        interface{} `json:"data"`
}

type RespError struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}
