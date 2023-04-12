package controllers

import "Final_Project/services"

type HttpServer struct {
	app services.ServiceInterface
}

func NewHttpServer(app services.ServiceInterface) HttpServer {
	return HttpServer{app: app}
}
