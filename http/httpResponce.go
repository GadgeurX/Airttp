package http

import "fmt"

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

func (res *Responce) Draw() {
	fmt.Println("Responce : ")
	fmt.Println("Raw : " + string(res.Raw))
	fmt.Println("Header : ")
	for key, value := range res.Headers {
		fmt.Println("     " + key + " : " + value)
	}
	fmt.Println("Body : ")
	fmt.Println(string(res.Body))
}