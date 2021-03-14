package service

import (
	"context"
	"errors"

	"github.com/hellojqk/config/server/entity"
	"github.com/hellojqk/config/server/repository"
)

// GroupInsertOne .
func GroupInsertOne(ctx context.Context, group entity.Group) (err error) {
	group.SetCreator("")
	insertResult, err := repository.DB.Collection("group").InsertOne(ctx, group)
	if err != nil {
		return err
	}
	if insertResult.InsertedID == nil {
		return errors.New("创建失败")
	}
	return nil
}
