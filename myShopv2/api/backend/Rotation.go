package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationReq struct {
	g.Meta `path:"/backend/rotation/add" tags:"Rotation" method:"post" summary:"You first rotation api"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"       v:"required#跳转链接不能为空"      dc:"跳转链接"`
	Sort   int    `json:"sort"       dc:"排序"`
}
type RotationRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	RotationId int `json:"rotation_id"`
}
type RotationDeleteReq struct {
	g.Meta `path:"/backend/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图接口"`
	Id     uint `v:"min:1#请选择需要删除的轮播图" dc:"轮播图id"`
}
type RotationDeleteRes struct{}

//cmd.go 启动
//↓
//group.Bind(controller.Rotation)
//↓
//框架用反射扫描 controller 所有方法
//↓
//读取每个方法的第二个参数类型
//↓
//读取该参数结构体里的 g.Meta
//↓
//把 path + method 注册进路由表
