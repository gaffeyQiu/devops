package devops

type Interface interface {
	// PipelineOperator 流水线相关
	PipelineOperator
}

//func GetDevOpsStatusCode(devopsErr error) int {
//	errStr := strings.TrimPrefix(devopsErr.Error(), "unexpected status code: ")
//	if code, err := strconv.Atoi(errStr); err == nil {
//		message := http.StatusText(code)
//		if !govalidator.IsNull(message) {
//			return code
//		}
//	}
//	if jErr, ok := devopsErr.(*ErrorResponse); ok {
//		return jErr.Response.StatusCode
//	}
//	return http.StatusInternalServerError
//}
//
//type ErrorResponse struct {
//	Body     []byte
//	Response *http.Response
//	Message  string
//}
//
//func (e *ErrorResponse) Error() string {
//	var u string
//	var method string
//	if e.Response != nil && e.Response.Request != nil {
//		req := e.Response.Request
//		u = fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.RequestURI())
//		method = req.Method
//	}
//	return fmt.Sprintf("%s %s: %d %s", method, u, e.Response.StatusCode, e.Message)
//}
