package controller

import (
	"Authority/app/api/model"
	"Authority/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllAuth(c *gin.Context)  {
	if res := haveAuth(c, "getAllAuth"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	authModel := model.Auth{}

	authJson := struct {
		Offset 	int 	`json:"offset" form:"offset"`
		Limit 	int 	`json:"limit" form:"limit"`
		Where 	struct{
			Title 	string 	`json:"title" form:"title"`
		}
	}{}

	if c.ShouldBind(&authJson) == nil {
		authJson.Offset = (authJson.Offset-1)*authJson.Limit
		res := authModel.GetAllAuth(authJson.Offset, authJson.Limit, authJson.Where.Title)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}
