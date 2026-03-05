package service

import (
	"context"
	"myShopv2/internal/model"
)

type (
	IRotation interface {
		// Create 创建内容
		Create(ctx context.Context, in *model.RotationCreateInput) (out model.RotationCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) error
	}
)

var (
	localRotation IRotation
)

func Rotation() IRotation {
	if localRotation == nil {
		panic("implement not found for interface IRotation, forgot register?")
	}
	return localRotation
}

func RegisterRotation(i IRotation) {
	localRotation = i
}
