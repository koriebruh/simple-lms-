package helper

import "jamal/tgs/models/web"

func WebResponsesFunc(code int, status string, data interface{}) web.WebResponse {
	return web.WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}

func WebResponsesFindAll(code int, status string, data interface{}) web.WebResponse {
	return web.WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
