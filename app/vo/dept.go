package vo

import "github.com/chenhu1001/gin-layui/app/model"

// 部门树结构
type DeptTreeNode struct {
	model.Dept
	Children []*DeptTreeNode `json:"children"` // 子栏目
}
