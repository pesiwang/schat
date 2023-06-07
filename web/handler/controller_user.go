package handler

import (
	"schat/app"
	"schat/infrastructure/logger"
	"schat/types"
	"schat/web/base"
	"schat/web/define"
	"time"

	"github.com/gin-gonic/gin"
)

func initControllerUser() {
	userGroup := base.GinRouter().Group("/user")

	userGroup.GET("/check_whether_register", checkWhetherRegister)
	userGroup.POST("/register", register)
	userGroup.POST("/login", login)
	userGroup.POST("/get_auth_code", getAuthCode)

	userGroup.GET("/get_info", getInfo)
}

func checkWhetherRegister(c *gin.Context) {
	var checkRegister define.CheckWhetherRegisterRequest
	err := c.ShouldBindJSON(&checkRegister)
	if err != nil {
		logger.Warnf("check params failed: %v", err)
		base.GinSimpleResponse(c, define.RESULT_CODE_CHECK_PARAMS_FAILED, err.Error())
		return
	}

	registered, err := app.NewUser().CheckWhetherRegister(&checkRegister.Phone)
	if err != nil {
		logger.Warnf("check params failed: %v", err)
		base.GinSimpleResponse(c, define.RESULT_CODE_SERVER_INNER_ERROR, err.Error())
		return
	}

	checkResp := define.CheckWhetherRegisterResponse{Registered: registered}

	base.GinDataResponse(c, define.RESULT_CODE_SUCC, "succ", checkResp)
}

func getAuthCode(c *gin.Context) {

}

func getInfo(c *gin.Context) {
	uid, _ := c.Get("uid")
	logger.Infof("uid from context: %v", uid)
	base.GinDataResponse(c, define.RESULT_CODE_SUCC, "succ", gin.H{"uid": uid, "name": "wolf"})
}

func login(c *gin.Context) {
	var login define.LoginRequest
	err := c.ShouldBindJSON(&login)
	if err != nil {
		logger.Warnf("check params failed: %v", err)

		base.GinSimpleResponse(c, define.RESULT_CODE_CHECK_PARAMS_FAILED, err.Error())
		return
	}

	logger.Infof("input:%v", login)

	// application service comes in

	var uid types.UserId = 100000
	tokenString, err := base.GenerateToken(uid, 30*time.Second)
	if err != nil {
		base.GinSimpleResponse(c, define.RESULT_CODE_SERVER_INNER_ERROR, err.Error())
		return
	}
	authResponse := define.AuthResponse{
		TokenString: tokenString,
	}
	base.GinDataResponse(c, define.RESULT_CODE_SUCC, "succ", authResponse)
}

func register(c *gin.Context) {
	var register define.RegisterRequest
	// use query because time_format validate tag not work in json format, need append "T00:00:00Z" to the end of date string
	err := c.ShouldBindQuery(&register)
	if err != nil {
		logger.Warnf("check params failed: %v", err)

		base.GinSimpleResponse(c, define.RESULT_CODE_CHECK_PARAMS_FAILED, err.Error())
		return
	}

	logger.Infof("input:%v", register)

	birthdayStr := register.Birthday.Format("2006-01-02")
	logger.Infof("birthday: %v", birthdayStr)

	// application service comes in

	var uid types.UserId = 100000
	tokenString, err := base.GenerateToken(uid, 10*time.Second)
	if err != nil {
		base.GinSimpleResponse(c, define.RESULT_CODE_SERVER_INNER_ERROR, err.Error())
		return
	}

	authResponse := define.AuthResponse{
		TokenString: tokenString,
	}
	base.GinDataResponse(c, define.RESULT_CODE_SUCC, "succ", authResponse)
}
