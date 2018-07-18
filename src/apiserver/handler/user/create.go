package user

import (
	"github.com/gin-gonic/gin"
	"apiserver/pkg/errno"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"apiserver/util"
	"apiserver/handler"
	"apiserver/model"
)

// create creates a new user account
func Create(c *gin.Context) {

	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := r.checkParam(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	//对密码加密
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	//向数据库添加记录
	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	handler.SendResponse(c, nil, rsp)
}

func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	return nil
}