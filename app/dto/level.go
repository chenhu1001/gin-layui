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
 * 职级Dto
 * @author 半城风雨
 * @since 2021/9/6
 * @File : level
 */
package dto

// 分页查询
type LevelPageReq struct {
	Name  string `form:"name"`  // 职级名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加职级
type LevelAddReq struct {
	Name   string `form:"name"  binding:"required"`
	Status string `form:"status"    binding:"required"`
	Sort   string `form:"sort"`
}

// 编辑职级
type LevelUpdateReq struct {
	Id     string `form:"id" binding:"required"`
	Name   string `form:"name"  binding:"required"`
	Status string `form:"status"    binding:"required"`
	Sort   string `form:"sort"`
}

// 设置状态
type LevelStatusReq struct {
	Id     string `form:"id" binding:"required"`
	Status string `form:"status"    binding:"required"`
}