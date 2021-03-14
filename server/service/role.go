package service

import (
	"context"
	"errors"

	"github.com/hellojqk/config/server/entity"
	"github.com/hellojqk/config/server/repository"
)

// RoleInsertOne .
func RoleInsertOne(ctx context.Context, role entity.Role) (err error) {
	role.SetCreator("")
	insertResult, err := repository.DB.Collection("role").InsertOne(ctx, role)
	if err != nil {
		return err
	}
	if insertResult.InsertedID == nil {
		return errors.New("创建失败")
	}
	return nil
}
