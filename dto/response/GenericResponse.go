package response

type GenericResponse struct {
	Body      any
	Status    int
	Message   string
	IsSuccess bool
}

func (r *GenericResponse) SetBody(body any) {

	r.Body = body
}

func (r *GenericResponse) getBody() any {

	return r.Body
}

func (r *GenericResponse) SetStatus(status int) {
	r.Status = status
}

func (r *GenericResponse) GetStatus() int {
	return r.Status
}

func (r *GenericResponse) SetMessage(message string) {

	r.Message = message
}

func (r *GenericResponse) GetMessage() string {

	return r.Message
}

func (r *GenericResponse) GetIsSuccess() bool {

	return r.IsSuccess
}

func (r *GenericResponse) SetIsSuccess(isSuccess bool) {

	r.IsSuccess = isSuccess
}
