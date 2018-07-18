package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"apiserver/model"
	"apiserver/handler"
	"apiserver/pkg/errno"
)

func Delete(c *gin.Context) {

	//解析出id的值
	userId, _ := strconv.Atoi(c.Param("id"))

	//删除用户
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
