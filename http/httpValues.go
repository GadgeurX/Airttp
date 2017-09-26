package http

type Value struct {
	Code int
	Message string
}

var Values = map[string]Value{
	"OK": Value{Code: 200, Message: "OK"},
	"NOT_FOUND": Value{Code: 404, Message: "Not Found"},
	"SERVER_ERROR": Value{Code: 500, Message: "Server Error"},
}
