package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	model_user "pchat/model/user"
	"pchat/utils"

	pb_common "pchat/pb/common"
	pb_user "pchat/pb/user"
	"pchat/repository/bson"
)

// @Description
// @Router		/users/{id} [put]
// @Tags		账户
// @Summary	更新个人信息
// @Accept		json
// @Produce	json
// @Param		id		path		string							true	"占位符"
// @Param		body	body		pb_user.UpdateProfileRequest	true	"body"
// @Success	200		{object}	nil
func updateProfile(ctx context.Context, req *pb_user.UpdateProfileRequest) (*pb_common.EmptyResponse, error) {
	setter := bson.M{}
	id := utils.GetUserId(ctx)
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("invalid user id")
	}
	if req.Avatar != nil {
		setter["avatar"] = req.Avatar.Value
	}
	if req.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password.Value), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		setter["password"] = string(hash)
	}
	if req.Name != nil {
		setter["name"] = req.Name.Value
	}
	if len(setter) > 0 {
		err := model_user.CUser.UpdateById(ctx, bson.NewObjectIdFromHex(id), bson.M{"$set": setter})
		if err != nil {
			return nil, err
		}
	}
	return &pb_common.EmptyResponse{}, nil
}
