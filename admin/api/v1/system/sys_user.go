package system

import (
	"strconv"

	"catering/global"
	"catering/model/common/request"
	"catering/model/common/response"
	"catering/model/system"
	systemReq "catering/model/system/request"
	systemRes "catering/model/system/response"
	"catering/pkg"
	"catering/pkg/valid"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	msg, err := valid.BindAndValid(c, &l)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		if err, user := userService.Login(u); err != nil {
			global.Log.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			b.tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

// 登录以后签发jwt
func (b *BaseApi) tokenNext(c *gin.Context, user system.SysUser) {
	j := &pkg.JWT{SigningKey: []byte(global.Config.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          uint(user.ID),
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.Log.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithDetailed(systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
	return
}

// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body systemReq.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse,msg=string} "用户注册账号,返回包括用户信息"
// @Router /user/register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	msg, err := valid.BindAndValid(c, &r)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities}
	err, userReturn := userService.Register(*user)
	if err != nil {
		global.Log.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body systemReq.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {object} response.Response{msg=string} "用户修改密码"
// @Router /user/changePassword [post]
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var user systemReq.ChangePasswordStruct
	msg, err := valid.BindAndValid(c, &user)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	u := &system.SysUser{Username: user.Username, Password: user.Password}
	if err, _ := userService.ChangePassword(u, user.NewPassword); err != nil {
		global.Log.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router /user/getUserList [post]
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	msg, err := valid.BindAndValid(c, &pageInfo)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	if err, list, total := userService.GetUserInfoList(pageInfo); err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags SysUser
// @Summary 更改用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuth true "用户UUID, 角色ID"
// @Success 200 {object} response.Response{msg=string} "设置用户权限"
// @Router /user/setUserAuthority [post]
func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var sua systemReq.SetUserAuth
	msg, err := valid.BindAndValid(c, &sua)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	userID := pkg.GetUserID(c)
	uuid := pkg.GetUserUuid(c)
	if err := userService.SetUserAuthority(userID, uuid, sua.AuthorityId); err != nil {
		global.Log.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		claims := pkg.GetUserInfo(c)
		j := &pkg.JWT{SigningKey: []byte(global.Config.JWT.SigningKey)} // 唯一签名
		claims.AuthorityId = sua.AuthorityId
		if token, err := j.CreateToken(*claims); err != nil {
			global.Log.Error("修改失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
		} else {
			c.Header("new-token", token)
			c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
			response.OkWithMessage("修改成功", c)
		}

	}
}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuthorities true "用户UUID, 角色ID"
// @Success 200 {object} response.Response{msg=string} "设置用户权限"
// @Router /user/setUserAuthorities [post]
func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	_ = c.ShouldBindJSON(&sua)
	if err := userService.SetUserAuthorities(sua.ID, sua.AuthorityIds); err != nil {
		global.Log.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {object} response.Response{msg=string} "删除用户"
// @Router /user/deleteUser [delete]
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	msg, err := valid.BindAndValid(c, &reqId)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	jwtId := pkg.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	if err := userService.DeleteUser(reqId.ID); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysUser true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "设置用户信息"
// @Router /user/setUserInfo [put]
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user system.SysUser
	msg, err := valid.BindAndValid(c, &user)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	user.Username = ""
	if err, ReqUser := userService.SetUserInfo(user); err != nil {
		global.Log.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
	}
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysUser true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "设置用户信息"
// @Router /user/SetSelfInfo [put]
func (b *BaseApi) SetSelfInfo(c *gin.Context) {
	var user system.SysUser
	_ = c.ShouldBindJSON(&user)
	user.Username = ""
	user.ID = uint64(pkg.GetUserID(c))
	if err, ReqUser := userService.SetUserInfo(user); err != nil {
		global.Log.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
	}
}

// @Tags SysUser
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取用户信息"
// @Router /user/getUserInfo [get]
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := pkg.GetUserUuid(c)
	if err, ReqUser := userService.GetUserInfo(uuid); err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
	}
}

// @Tags SysUser
// @Summary 重置用户密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body system.SysUser true "ID"
// @Success 200 {object} response.Response{msg=string} "重置用户密码"
// @Router /user/resetPassword [post]
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var user system.SysUser
	_ = c.ShouldBindJSON(&user)
	if err := userService.ResetPassword(uint(user.ID)); err != nil {
		global.Log.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
	} else {
		response.OkWithMessage("重置成功", c)
	}
}
