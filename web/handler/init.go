package handler

func InitHandler() {
	initMiddlewareGlobal()
	initControllerServer()
	initControllerUser()
}
