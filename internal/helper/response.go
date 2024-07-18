package helper

func ResponseFormat(status string, code int, message string, data any, meta any) map[string]any {
	var result = make(map[string]any)
	result["code"] = code
	result["status"] = status
	result["message"] = message
	if data != nil {
		result["data"] = data
	}
	if meta != nil {
		result["meta"] = meta
	}
	return result
}

func ResponseFormatNonData(code int, message string, status string) map[string]any {
	var result = make(map[string]any)
	result["code"] = code
	result["status"] = status
	result["message"] = message
	return result
}
