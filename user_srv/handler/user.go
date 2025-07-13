package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"mxshop_srvss/user_srv/gobal"
	"mxshop_srvss/user_srv/model"
	"mxshop_srvss/user_srv/proto"
)

type UserServer struct {
}

// 分页函数
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if page <= 0 {
			page = 1
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		NickName: user.Nickname,
		Gender:   user.Genders,
		Role:     int32(user.Role),
	}
	return userInfoRsp
}

func (s *UserServer) GetUserInfo(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	//获取用户列表
	var users []model.User
	result := gobal.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = uint32(int32(result.RowsAffected))
	gobal.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)
	for _, user := range users {
		userInfoRsp := ModelToResponse(user)

		rsp.Data = append(rsp.Data, data)
	}
	return rsp, nil
}

func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	//通过手机号查询用户
	var user model.User
	result := gobal.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	//通过用户ID查询用户
	var user model.User
	result := gobal.DB.Where(&user, req.Id).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

// CreateUser 新建用户
func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	//新建用户
	var user model.User
	result := gobal.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户存在")
	}
	user.Mobile = req.Mobile
	user.Nickname = req.NickName
	// 密码加密
	options := &password.Options{16, 10000, 32, sha512.New}

	salt, encodedPwd := password.Encode(req.Password, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	result = gobal.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}
