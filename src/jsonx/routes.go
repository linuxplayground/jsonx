package main

import (
	
	"latham.nz/featly.common"
)

var routes = common.Routes{
	common.Route{
		"Create",
		"POST",
		"/",
		Create,
	},
}
