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
 * 会员等级-控制器
 * @author 半城风雨
 * @since 2021/11/13
 * @File : member_level
 */
package controller

import (
	"github.com/chenhu1001/gin-layui/app/dto"
	"github.com/chenhu1001/gin-layui/app/model"
	"github.com/chenhu1001/gin-layui/app/service"
	"github.com/chenhu1001/gin-layui/utils"
	"github.com/chenhu1001/gin-layui/utils/common"
	"github.com/chenhu1001/gin-layui/utils/gconv"
	"github.com/chenhu1001/gin-layui/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

var MemberLevel = new(memberLevelCtl)

type memberLevelCtl struct{}

func (c *memberLevelCtl) Index(ctx *gin.Context) {
	// 渲染模板
	response.BuildTpl(ctx, "memberlevel_index.html").WriteTpl()
}

func (c *memberLevelCtl) List(ctx *gin.Context) {
	// 参数
	var req *dto.MemberLevelPageReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用分页查询方法
	list, count, err := service.MemberLevel.GetList(req)
	if err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code:  0,
		Msg:   "查询成功",
		Data:  list,
		Count: count,
	})
}

func (c *memberLevelCtl) Edit(ctx *gin.Context) {
	// 记录ID
	id := gconv.Int(ctx.Query("id"))
	if id > 0 {
		// 编辑
		info := &model.MemberLevel{Id: id}
		has, err := info.Get()
		if !has || err != nil {
			ctx.JSON(http.StatusOK, common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		// 渲染模板
		response.BuildTpl(ctx, "memberlevel_edit.html").WriteTpl(gin.H{
			"info": info,
		})
	} else {
		// 添加
		response.BuildTpl(ctx, "memberlevel_edit.html").WriteTpl()
	}
}

func (c *memberLevelCtl) Add(ctx *gin.Context) {
	// 参数
	var req *dto.MemberLevelAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用添加方法
	rows, err := service.MemberLevel.Add(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "添加成功",
	})
}

func (c *memberLevelCtl) Update(ctx *gin.Context) {
	// 参数
	var req *dto.MemberLevelUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用更新方法
	rows, err := service.MemberLevel.Update(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (c *memberLevelCtl) Delete(ctx *gin.Context) {
	// 记录ID
	ids := ctx.Param("ids")
	if ids == "" {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
		return
	}

	// 调用删除方法
	rows, err := service.MemberLevel.Delete(ids)
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}

func (c *memberLevelCtl) GetMemberLevelList(ctx *gin.Context) {
	list := make([]model.MemberLevel, 0)
	utils.XormDb.Where("mark=1").OrderBy("sort asc").Find(&list)
	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}