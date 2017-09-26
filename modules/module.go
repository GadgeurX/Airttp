package modules

import (
	"net/rpc"
	"Airttp/http"
	"Airttp/logger"
	"net"
	"strconv"
)

type ModuleParams struct {
	Sess http.Session
	Req http.Request
	Res http.Responce
}

func (params *ModuleParams) Copy(param2 ModuleParams) {
	params.Req.Copy(param2.Req)
	params.Res.Copy(param2.Res)
}

type module struct {
	name string
	addr string
	client *rpc.Client
	priority int
	connected bool
	require bool
}

func NewModule(name string, addr string, priority int, require bool) *module {
	module := new(module)
	module.connected = false
	module.name = name
	module.addr = addr
	module.priority = priority
	module.require = require
	module.Connect()
	return module
}

func (mod *module) Connect() {
	conn, err := net.Dial("tcp", mod.addr)
	if err != nil {
		logger.GetInstance().Error("Error connection to module \"" + mod.name + "\"")
	} else {
		logger.GetInstance().Notice("Connected to module \"" + mod.name + "\" priority " + strconv.Itoa(mod.priority))
		mod.connected = true
		mod.client = rpc.NewClient(conn)
	}
}

func (mod *module) execRequest(params ModuleParams) ModuleParams {
	if (mod.client == nil) {
		if (mod.require) {
			handleMissingRequireModule(&params)
		}
		return params
	}
	var result ModuleParams
	err := mod.client.Call("Http.Module", params, &result)
	if err != nil {
		logger.GetInstance().Error("Error call module \"" + mod.name + "\" : " + err.Error())
		mod.connected = false
		mod.client = nil
		handleMissingRequireModule(&params)
		return params
	}
	return result
}

func handleMissingRequireModule(params *ModuleParams) {
	params.Res.Body = []byte("Error missing module")
	params.Res.Headers["Content-Length"] = strconv.Itoa(len(params.Res.Body))
	params.Res.Code = http.Values["SERVER_ERROR"].Code
	params.Res.Message = http.Values["SERVER_ERROR"].Message

	params.Res.Raw = []byte("HTTP/1.1 " + strconv.Itoa(params.Res.Code) + " " + params.Res.Message + "\r\n")
	for key, value := range params.Res.Headers {
		params.Res.Raw = append(params.Res.Raw[:], []byte(key + ": " + value + "\r\n")[:]...)
	}
	params.Res.Raw = append(params.Res.Raw[:], []byte("\r\n")[:]...)
	params.Res.Raw = append(params.Res.Raw[:], params.Res.Body[:]...)
}
