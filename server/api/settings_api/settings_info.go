package settings_api

import (
	"github.com/gin-gonic/gin"
	"go-blog/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//c.JSON(200, gin.H{"msg": "xxx"})
	// OK 测试
	//res.Ok("操作成功", c)
	// OkWithData 测试
	//data := map[string]string{
	//	"id": "123",
	//}
	//res.OkWithData(data, "请求成功", c)
	// OkWithMsg 测试
	//res.OkWithMsg("请求成功", c)

	// OkWithDetailed
	//res.OkWithDetailed(data, "数据详情", c)

	// Fail 测试
	//res.Fail("请求失败", c)

	// FailWithMsg 测试
	res.FailWithMsg("请求失败", c)
	// FailWithCode 系统错误、未知错误测试
	//res.FailWithCode(1001, c)
}
