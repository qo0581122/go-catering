import service from '@/utils/request'
// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {
  return service({
    url: '/base/login',
    method: 'post',
    data: data
  })
}

// @Summary 获取验证码
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/captcha [post]
export const captcha = (data) => {
  return service({
    url: '/base/captcha',
    method: 'post',
    data: data
  })
}

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/resige [post]
export const register = (data) => {
  return service({
    url: '/systemUser/register',
    method: 'post',
    data: data
  })
}

// @Summary 修改密码
// @Produce  application/json
// @Param data body {username:"string",password:"string",newPassword:"string"}
// @Router /systemUser/changePassword [post]
export const changePassword = (data) => {
  return service({
    url: '/systemUser/changePassword',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemUser/getUserList [post]
export const getUserList = (data) => {
  return service({
    url: '/systemUser/getUserList',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.SetUserAuth true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /systemUser/setUserAuthority [post]
export const setUserAuthority = (data) => {
  return service({
    url: '/systemUser/setUserAuthority',
    method: 'post',
    data: data
  })
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /systemUser/deleteUser [delete]
export const deleteUser = (data) => {
  return service({
    url: '/systemUser/deleteUser',
    method: 'delete',
    data: data
  })
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /systemUser/setUserInfo [put]
export const setUserInfo = (data) => {
  return service({
    url: '/systemUser/setUserInfo',
    method: 'put',
    data: data
  })
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /systemUser/setSelfInfo [put]
export const setSelfInfo = (data) => {
  return service({
    url: '/systemUser/setSelfInfo',
    method: 'put',
    data: data
  })
}

// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.setUserAuthorities true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /systemUser/setUserAuthorities [post]
export const setUserAuthorities = (data) => {
  return service({
    url: '/systemUser/setUserAuthorities',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemUser/getUserInfo [get]
export const getUserInfo = () => {
  return service({
    url: '/systemUser/getUserInfo',
    method: 'get'
  })
}

export const resetPassword = (data) => {
  return service({
    url: '/systemUser/resetPassword',
    method: 'post',
    data: data
  })
}
