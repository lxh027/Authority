package controller

import (
	"Authority/app/api/model"
	"Authority/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)
// TODO 注册权限
func GetAllRole(c *gin.Context)  {
	roleModel := model.Role{}

	roleJson := struct {
		Offset 	int 	`json:"offset" form:"offset"`
		Limit 	int 	`json:"limit" form:"limit"`
		Where 	struct{
			Name 	string 	`json:"name" form:"name"`
			Desc 	string 	`json:"desc" form:"desc"`
		}
	}{}

	if c.ShouldBind(&roleJson) == nil {
		roleJson.Offset = (roleJson.Offset-1)*roleJson.Limit
		res := roleModel.GetAllRole(roleJson.Offset, roleJson.Limit, roleJson.Where.Name, roleJson.Where.Desc)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}