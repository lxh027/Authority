package controller

import (
	"Authority/app/api/model"
	"Authority/app/api/validate"
	"Authority/app/common"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
// TODO
func AddUserRoles(c *gin.Context)  {
	if res := haveAuth(c, "roleAssign"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userRoleValidate := validate.UserRoleValidate
	userRoleModel := model.UserRole{}

	if res, err:= userRoleValidate.Validate(c, "addGroup"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	userRolesJson := struct {
		Uid 	int 	`json:"uid" form:"uid"`
		Rids	string 	`json:"rids" form:"rids"`
	}{}

	if c.ShouldBind(&userRolesJson) == nil {
		var rids []int
		_ = json.Unmarshal([]byte((userRolesJson.Rids)), &rids)
		fmt.Println(rids)
		for _, rid := range rids {
			res := userRoleModel.AddUserRole(model.UserRole{Uid: userRolesJson.Uid, Rid: rid})
			if res.Status != common.CodeSuccess {
				c.JSON(http.StatusOK, common.ApiReturn(res.Status, "编号为"+string(rune(rid))+"的权限添加失败", res.Data))
			}
		}
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeSuccess, "添加成功", true))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func DeleteUserRoles(c *gin.Context)  {
	if res := haveAuth(c, "roleAssign"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userRoleValidate := validate.UserRoleValidate
	userRoleModel := model.UserRole{}

	if res, err:= userRoleValidate.Validate(c, "deleteGroup"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	userRolesJson := struct {
		Uid 	int 	`json:"uid" form:"uid"`
		Rids	string 	`json:"rids" form:"rids"`
	}{}

	if c.ShouldBind(&userRolesJson) == nil {
		var rids []int
		_ = json.Unmarshal([]byte((userRolesJson.Rids)), &rids)
		fmt.Println(rids)
		for _, rid := range rids {
			res := userRoleModel.DeleteUserRole(model.UserRole{Uid: userRolesJson.Uid, Rid: rid})
			if res.Status != common.CodeSuccess {
				c.JSON(http.StatusOK, common.ApiReturn(res.Status, "编号为"+string(rune(rid))+"的权限删除失败", res.Data))
			}
		}
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeSuccess, "删除成功", true))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func GetUserRolesList(c *gin.Context) {
	if res := haveAuth(c, "roleAssign"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userRoleValidate := validate.UserRoleValidate
	roleModel := model.Role{}

	if res, err:= userRoleValidate.Validate(c, "getUserRole"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	allRoles := roleModel.GetRoleNoRules()

	roleJson := struct {
		Uid 	int 	`json:"uid" form:"uid"`
	}{}

	if c.ShouldBind(&roleJson) == nil {
		res := roleModel.GetUserRole(roleJson.Uid)
		roles := res.Data.([]model.Role)
		var val []int
		for _, role := range roles {
			val = append(val, role.Rid)
		}
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, map[string]interface{} {
			"allRoles" 	: allRoles.Data,
			"values"	: val,

		}))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}
