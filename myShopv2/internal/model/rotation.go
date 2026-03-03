package model

// RotationCreateUpdateBase 创建/修改内容基类
type RotationCreateUpdateBase struct {
	PicUrl string // 内容模型
	Link   string // 标题
	Sort   int    // 内容

}

//RotationCreateInput 创建内容

type RotationCreateInput struct {
	RotationCreateUpdateBase
	//UserId uint
}

// RotationCreateOutput 创建内容返回结果

type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}
