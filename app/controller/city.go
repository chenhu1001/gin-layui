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
 * 城市管理-控制器
 * @author 半城风雨
 * @since 2021/11/13
 * @File : city
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

var City = new(cityCtl)

type cityCtl struct{}

func (c *cityCtl) Index(ctx *gin.Context) {
	// 渲染模板
	response.BuildTpl(ctx, "city_index.html").WriteTpl()
}

func (c *cityCtl) List(ctx *gin.Context) {
	//// 参数
	//var req *dto.CityQueryReq
	//if err := ctx.ShouldBind(&req); err != nil {
	//	ctx.JSON(http.StatusOK, common.JsonResult{
	//		Code: -1,
	//		Msg:  err.Error(),
	//	})
	//	return
	//}

	// 调用获取列表方法
	list := service.City.GetList(nil)

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}

func (c *cityCtl) Edit(ctx *gin.Context) {
	id := gconv.Int(ctx.Query("id"))
	if id > 0 {
		// 编辑
		info := &model.City{Id: id}
		has, err := info.Get()
		if !has || err != nil {
			ctx.JSON(http.StatusOK, common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		// 渲染模板
		response.BuildTpl(ctx, "city_edit.html").WriteTpl(gin.H{
			"info":      info,
			"levelList": common.CITY_LEVEL,
		})
	} else {
		// 添加
		// 渲染模板
		response.BuildTpl(ctx, "city_edit.html").WriteTpl()
	}
}

func (c *cityCtl) Add(ctx *gin.Context) {
	// 参数
	var req *dto.CityAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用添加方法
	rows, err := service.City.Add(req, utils.Uid(ctx))
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

func (c *cityCtl) Update(ctx *gin.Context) {
	// 参数
	var req *dto.CityUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用更新方法
	rows, err := service.City.Update(req, utils.Uid(ctx))
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

func (c *cityCtl) Delete(ctx *gin.Context) {
	// 参数
	ids := ctx.Param("ids")
	if ids == "" {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
		return
	}

	// 调用删除方法
	rows, err := service.City.Delete(ids)
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

func (c *cityCtl) GetChilds(ctx *gin.Context) {
	// 参数验证
	var req *dto.CityChildReq
	if err := ctx.ShouldBind(&req); err != nil {
		// 返回结果
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用获取子级城市
	list, err := service.City.GetChilds(req.CityCode)
	if err != nil {
		// 返回结果
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}