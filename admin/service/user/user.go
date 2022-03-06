package user

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var UserService userService = NewUserService()

func NewUserService() userService {
	return userServiceImpl{}
}

type userServiceImpl struct {
}

func (impl userServiceImpl) GetById(id uint64) *model.User {
	return model.GetUserById(id)
}
func (impl userServiceImpl) List(params *model.User) []*model.User {
	return model.ListUser(params)
}

func (impl userServiceImpl) Count() int {
	return model.CountUser()
}

func (impl userServiceImpl) ListPage(pageNum, pageSize int, params *model.User) *response.ApiResponse {
	users, err := model.ListUserPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	total := impl.Count()
	res := &response.ApiResponse{List: users, Total: total}
	return res
}
