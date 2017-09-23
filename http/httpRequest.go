package http

import "fmt"

type Request struct {
	Raw []byte
	Header []byte
	Method string
	Uri string
	Headers map[string]string
	Params map[string]string
	Body []byte
}

func NewRequest(data []byte) *Request {
	request := new(Request)
	request.Headers = make(map[string]string)
	request.Params = make(map[string]string)
	request.Raw = data
	return request
}

func (req *Request) Draw() {
	fmt.Println("Request : " + req.Method)
	fmt.Println("Uri : " + req.Uri)
	fmt.Println("Raw : " + string(req.Raw))
	fmt.Println("Header : ")
	for key, value := range req.Headers {
		fmt.Println("     " + key + " : " + value)
	}
	fmt.Println("Params : ")
	for key, value := range req.Params {
		fmt.Println("     " + key + " : " + value)
	}
	fmt.Println("Body : ")
	fmt.Println(string(req.Body))
}

func (req *Request) Copy(req2 Request) {
	req.Raw = req2.Raw
	req.Body = req2.Body
	req.Header = req2.Header
	req.Method = req2.Method
	req.Uri = req2.Uri
	req.Headers = req2.Headers
	req.Params = req2.Params
}
