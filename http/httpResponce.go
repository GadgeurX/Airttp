package http

type Responce struct {
	Raw []byte
	Header []byte
	Headers map[string]string
	Body []byte
	Code int
	Message string
}

func NewResponce() *Responce {
	responce := new(Responce)
	responce.Headers = make(map[string]string)
	return responce
}

func (res *Responce) Copy(res2 Responce) {
	res.Raw = res2.Raw
	res.Body = res2.Body
	res.Header = res2.Header
	res.Code = res2.Code
	res.Headers = res2.Headers
	res.Message = res2.Message
}