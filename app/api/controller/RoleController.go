package controller

import (
	"Authority/app/api/model"
	"Authority/app/api/validate"
	"Authority/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)
// TODO 注册权限
func GetAllRole(c *gin.Context)  {//??
	if res := haveAuth(c, "getAllRole"); res != common.Authed {//getAllUser怎么改？
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
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

func GetRoleByID(c *gin.Context) {//jun
	if res := haveAuth(c, "getAllRole"); res != common.Authed {//getAllUser怎么改？
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	roleValidate := validate.RoleValidate
	roleModel := model.Role{}

	if res, err:= roleValidate.Validate(c, "find"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	var roleJson model.Role

	if c.ShouldBind(&roleJson) == nil {
		res := roleModel.GetRoleByID(roleJson.Rid)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func AddRole(c *gin.Context) { //jun
	if res := haveAuth(c, "addRole"); res != common.Authed {//getAllUser怎么改？
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	roleValidate := validate.RoleValidate
	roleModel := model.Role{}

	if res, err:= roleValidate.Validate(c, "add"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	var roleJson model.Role
	if c.ShouldBind(&roleJson) == nil {
		//userJson.Password = common.GetMd5(userJson.Password)
		res := roleModel.AddRole(roleJson)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}
