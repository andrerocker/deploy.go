package engine

import (
	"./command"
	"./config"
	"./http"
	"./http/gin"
	"fmt"
	"strings"
)

type Engine struct {
	http   http.Engine
	config config.Configuration
}

func New() Engine {
	return Engine{gin.New(), config.New()}
}

func (self Engine) Draw() {
	for endpoint, commands := range self.config.Commands {
		formattedEndpoint := fmt.Sprintf("/%s/:%s", endpoint, endpoint)

		for _, verbs := range commands {
			for verb, command := range verbs {
				wrappedHandler := wrapValuesHandler(endpoint, command.(string))
				self.http.Register(verb, formattedEndpoint, wrappedHandler)
			}
		}
	}
}

func (self Engine) Start() {
	daemon := self.config.Daemon
	listen := daemon.BindUrl()
	self.http.Start(listen)
}

func wrapValuesHandler(paramName, commandTemplate string) func(http.Request) {
	return func(request http.Request) {
		target := fmt.Sprintf("{%s}", paramName)
		content := request.Parameter(paramName)
		compiled := strings.Replace(commandTemplate, target, content, -1)
		command.ExecuteCommand(request.Writer(), compiled)
	}
}
