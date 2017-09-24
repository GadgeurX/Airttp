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
}

func NewModule(name string, addr string, priority int) *module {
	module := new(module)
	module.connected = false
	module.name = name
	module.addr = addr
	module.priority = priority
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
		return params
	}
	var result ModuleParams
	err := mod.client.Call("Http.Module", params, &result)
	if err != nil {
		logger.GetInstance().Error("Error call module \"" + mod.name + "\" : " + err.Error())
		mod.connected = false
		mod.client = nil
		return params
	}
	return result
}
