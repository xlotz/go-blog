package routers

import "go-blog/api"

//func SettingsRouter(router *gin.Engine) {
//	settingsApi := api.ApiGroupApp.SettingsApi
//	router.GET("/", settingsApi.SettingsInfoView)
//}

func (r RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	r.GET("/", settingsApi.SettingsInfoView)
}
