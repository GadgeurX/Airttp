package modules

import (
	"Airttp/http"
	"Airttp/config"
	"Airttp/logger"
	"strconv"
	"net"
	"time"
)

var m_ModuleManager *moduleManager = nil

type moduleManager struct {
	modules []*module
	maxPriority int
}

func NewModuleMangert() *moduleManager {
	moduleManager := new(moduleManager)
	return moduleManager
}

func (manager *moduleManager) LoadModules() {
	modulesList := config.GetConfigInstance().GetModulesList()
	for moduleName,moduleInfo := range modulesList {
		priority, err := strconv.Atoi(moduleInfo["priority"])
		require, err2 := strconv.ParseBool(moduleInfo["require"])
		if (err != nil || err2 != nil) {
			logger.GetInstance().Error("Module moduleName invalid propriety")
		} else {
			logger.GetInstance().Info("Load module \"" + moduleName + "\" at \"" + moduleInfo["addr"] + "\" and priority " + moduleInfo["priority"])
			manager.modules = append(manager.modules, NewModule(moduleName, moduleInfo["addr"], priority, require))
		}
	}
	manager.maxPriority = manager.SetMaxPriority()
	go manager.CheckModulesConnection()
}

func (manager *moduleManager) CheckModulesConnection() {
	for {
		time.Sleep(10 * time.Second)
		for _, moduleElem := range manager.modules {
			if (!moduleElem.connected) {
				moduleElem.Connect()
			}
		}
	}
}

func (manager *moduleManager) SetMaxPriority() int {
	maxPrio := 0
	for _, moduleElem := range manager.modules {
		if (moduleElem.priority > maxPrio) {
			maxPrio = moduleElem.priority
		}
	}
	manager.maxPriority = maxPrio
	return maxPrio
}

func (manager *moduleManager) ExecRequest(conn net.Conn, sess *http.Session, req *http.Request, res *http.Responce) {
	start := time.Now()
	var params ModuleParams
	params.Sess = *sess
	params.Req = *req
	params.Res = *res
	prio := 1
	for prio <= manager.maxPriority {
		for _, moduleElem := range manager.modules {
			if (moduleElem.priority == prio) {
				params = moduleElem.execRequest(params)
			}
		}
		prio++
	}
	elapsed := time.Since(start)
	logger.GetInstance().Info(params.Req.Method + " " + params.Req.Uri + " " + elapsed.String())
	conn.Write(params.Res.Raw)
}

func GetManagerInstance() *moduleManager {
	if (m_ModuleManager != nil) {
		return m_ModuleManager
	} else {
		m_ModuleManager = NewModuleMangert()
		return m_ModuleManager
	}
}