package registry

import (
	"github.com/muhammadnurbasari/test-be-majoo/routes/reportRoutes"
	"github.com/muhammadnurbasari/test-be-majoo/routes/userRoutes"
)

//initializeDomainModules calls the domain module routes
//in folder routes/*
func (reg *AppRegistry) initializeDomainModules() {
	userRoutes.UserRoutes(reg.serverHttp.GetRouteEngine(), reg.ConnPublic.DB())
	reportRoutes.ReportRoutes(reg.serverHttp.GetRouteEngine(), reg.ConnPublic.DB())
}
