package shop

type Category struct {
	Id       int64  `json:"id"`
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
	Level    int    `json:"level"`
	IsLeaf   string `json:"is_leaf"`
	Enable   string `json:"enable"`
}
