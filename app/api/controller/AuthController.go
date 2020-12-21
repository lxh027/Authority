package controller

import (
	"Authority/app/api/model"
	"Authority/app/api/validate"
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

func GetParentAuth(c *gin.Context)  {
	if res := haveAuth(c, "getAllAuth"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	authValidate := validate.AuthValidate
	authModel := model.Auth{}

	if res, err:= authValidate.Validate(c, "findParent"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	authJson := struct {
		Parent 	int 	`json:"parent" form:"parent"`
	}{}

	if c.ShouldBind(&authJson) == nil {
		res := authModel.GetParentAuth(authJson.Parent)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func AddAuth(c *gin.Context)  {
	if res := haveAuth(c, "addAuth"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	authValidate := validate.AuthValidate
	authModel := model.Auth{}

	if res, err:= authValidate.Validate(c, "add"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	var authJson model.Auth

	if c.ShouldBind(&authJson) == nil {
		authJson.Target = "_self"
		if authJson.Type != 1 {
			authJson.Href = ""
		}
 		res := authModel.AddAuth(authJson)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func DeleteAuth(c *gin.Context)  {
	if res := haveAuth(c, "deleteAuth"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	authValidate := validate.AuthValidate
	authModel := model.Auth{}

	if res, err:= authValidate.Validate(c, "delete"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	authIDJson := struct {
		Aid	int `json:"aid" form:"aid"`
	}{}

	if c.ShouldBind(&authIDJson) == nil {
		res := authModel.DeleteAuth(authIDJson.Aid)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}