// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 栏目-服务类
 * @author 半城风雨
 * @since 2021/11/13
 * @File : item_cate
 */
package service

import (
	"errors"
	"github.com/chenhu1001/gin-layui/app/dto"
	"github.com/chenhu1001/gin-layui/app/model"
	"github.com/chenhu1001/gin-layui/app/vo"
	"github.com/chenhu1001/gin-layui/utils"
	"github.com/chenhu1001/gin-layui/utils/gconv"
	"reflect"
	"strings"
	"time"
)

var ItemCate = new(itemCateService)

type itemCateService struct{}

func (s *itemCateService) GetList(req *dto.ItemCateQueryReq) []vo.ItemCateInfoVo {
	// 创建查询对象
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 栏目名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 排序
	query = query.OrderBy("sort asc")
	// 对象转换
	var list []model.ItemCate
	query.Find(&list)

	// 数据处理
	var result []vo.ItemCateInfoVo
	for _, v := range list {
		item := vo.ItemCateInfoVo{}
		item.ItemCate = v
		// 站点封面
		if v.IsCover == 1 && v.Cover != "" {
			item.Cover = utils.GetImageUrl(v.Cover)
		}
		// 获取栏目
		if v.ItemId > 0 {
			var itemInfo model.Item
			utils.XormDb.Id(item.Id).Get(&itemInfo)
			item.ItemName = itemInfo.Name
		}
		// 加入数组
		result = append(result, item)
	}
	return result
}

func (s *itemCateService) Add(req *dto.ItemCateAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.ItemCate
	entity.Name = req.Name
	entity.Pid = gconv.Int(req.Pid)
	entity.ItemId = gconv.Int(req.ItemId)
	entity.Pinyin = req.Pinyin
	entity.Code = req.Code
	entity.Status = gconv.Int(req.Status)
	entity.Note = req.Note
	entity.Sort = gconv.Int(req.Sort)

	// 封面
	entity.IsCover = gconv.Int(req.IsCover)
	if entity.IsCover == 1 {
		// 有封面
		cover, err := utils.SaveImage(req.Cover, "item_cate")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		// 没封面
		entity.Cover = ""
	}
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.Mark = 1

	// 插入数据
	return entity.Insert()
}

func (s *itemCateService) Update(req *dto.ItemCateUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &model.ItemCate{Id: gconv.Int(req.Id)}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 设置对象
	entity.Name = req.Name
	entity.Pid = gconv.Int(req.Pid)
	entity.ItemId = gconv.Int(req.ItemId)
	entity.Pinyin = req.Pinyin
	entity.Code = req.Code
	entity.Status = gconv.Int(req.Status)
	entity.Note = req.Note
	entity.Sort = gconv.Int(req.Sort)

	// 封面
	entity.IsCover = gconv.Int(req.IsCover)
	if entity.IsCover == 1 {
		// 有封面
		cover, err := utils.SaveImage(req.Cover, "item_cate")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		// 没封面
		entity.Cover = ""
	}
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()

	// 更新记录
	return entity.Update()
}

func (s *itemCateService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.ItemCate{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, errors.New("删除失败")
		}
		return rows, nil
	} else {
		// 批量删除
		return 0, nil
	}
}

func (s *itemCateService) GetCateName(cateId int, delimiter string) string {
	// 声明数组
	list := make([]string, 0)
	for {
		if cateId <= 0 {
			// 退出
			break
		}
		// 业务处理
		var info model.ItemCate
		has, err := utils.XormDb.Id(cateId).Get(&info)
		if err != nil || !has {
			break
		}
		// 上级栏目ID
		cateId = info.Pid
		// 加入数组
		list = append(list, info.Name)
	}
	// 结果数据处理
	if len(list) > 0 {
		// 数组反转
		utils.Reverse(&list)
		// 拼接字符串
		return strings.Join(list, delimiter)
	}
	return ""
}

// 获取子级菜单
func (s *itemCateService) GetCateTreeList(itemId int, pid int) ([]*vo.CateTreeNode, error) {
	var cateNote vo.CateTreeNode
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 站点ID
	if itemId > 0 {
		query = query.Where("item_id=?", itemId)
	}
	// 返回字段
	query.Cols("id,name,pid")
	// 排序
	query = query.OrderBy("sort asc")
	// 查询所有
	data := make([]model.ItemCate, 0)
	err := query.Find(&data)
	if err != nil {
		return nil, errors.New("系统错误")
	}
	makeCateTree(data, &cateNote)
	return cateNote.Children, nil
}

//递归生成分类列表
func makeCateTree(cate []model.ItemCate, tn *vo.CateTreeNode) {
	for _, c := range cate {
		if c.Pid == tn.Id {
			child := &vo.CateTreeNode{}
			child.ItemCate = c
			tn.Children = append(tn.Children, child)
			makeCateTree(cate, child)
		}
	}
}

// 数据源转换
func (s *itemCateService) MakeList(data []*vo.CateTreeNode) []map[string]string {
	cateList := make([]map[string]string, 0)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		// 一级栏目
		for _, val := range data {
			item := map[string]string{}
			item["id"] = gconv.String(val.Id)
			item["name"] = val.Name
			cateList = append(cateList, item)

			// 二级栏目
			for _, v := range val.Children {
				item2 := map[string]string{}
				item2["id"] = gconv.String(v.Id)
				item2["name"] = "|--" + v.Name
				cateList = append(cateList, item2)

				// 三级栏目
				for _, vt := range v.Children {
					item3 := map[string]string{}
					item3["id"] = gconv.String(vt.Id)
					item3["name"] = "|--|--" + vt.Name
					cateList = append(cateList, item3)
				}
			}
		}
	}
	return cateList
}