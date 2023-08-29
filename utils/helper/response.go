package helper

func FailedResponse(msg string, code int, error_msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":        "Failed.",
		"message":       msg,
		"error_message": error_msg,
		"code":          code,
	}
}

func SuccessResponse(msg string, code int) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success.",
		"message": msg,
		"code":    code,
	}
}

func FailedWithDataResponse(msg string, code int, error_msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":        "Failed With Data.",
		"message":       msg,
		"error_message": error_msg,
		"code":          code,
		"data":          data,
	}
}

func SuccessWithDataResponse(msg string, code int, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success With Data.",
		"message": msg,
		"code":    code,
		"data":    data,
	}
}
