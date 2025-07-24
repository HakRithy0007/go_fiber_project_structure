package responses

type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

func NewResponse(message string, statusCode int, data interface{}) Response {
	return Response{
		Success:    true,
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	}
}

type ResponseWithPaging struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	Total      int         `json:"total"`
}

func NewResponseWithPaing(message string, statusCode int, data interface{}, page int, perpage int, total int) ResponseWithPaging {
	return ResponseWithPaging{
		Success:    true,
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
		Page:       page,
		PerPage:    perpage,
		Total:      total,
	}
}
