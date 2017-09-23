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
}

func NewModule(name string, addr string, priority int) *module {
	module := new(module)
	module.name = name
	module.addr = addr
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		logger.GetInstance().Error("Error connection to module \"" + module.name + "\"")
		return module
	}
	logger.GetInstance().Notice("Connected to module \"" + module.name + "\" priority " + strconv.Itoa(priority))
	module.priority = priority
	module.client = rpc.NewClient(conn)
	return module
}

func (mod *module) execRequest(params ModuleParams) ModuleParams {
	if (mod.client == nil) {
		return params
	}
	var result ModuleParams
	err := mod.client.Call("Http.Module", params, &result)
	if err != nil {
		logger.GetInstance().Error("Error call module \"" + mod.name + "\" : " + err.Error())
		return params
	}
	return result
}
