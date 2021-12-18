package v1

import (
	"Practice/go-programming-tour-book/blog-service/global"
	"Practice/go-programming-tour-book/blog-service/internal/service"
	"Practice/go-programming-tour-book/blog-service/pkg/app"
	"Practice/go-programming-tour-book/blog-service/pkg/convert"
	"Practice/go-programming-tour-book/blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	id := convert.StrTo(c.Param("id")).MustUInt32()
	response := app.NewResponse(c)
	svc := service.New(c)
	article, err := svc.GetArticleById(id)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleById err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}
	response.ToResponse(article)
}

//TODO

// @Summary 获取文章列表
// @Produce  json
// @Param title query string false "文章标题" maxlength(100)
// @Param desc query string  false "文章简述"  maxlength(255)
// @Param content query string false "文章内容" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{PageSize: app.GetPageSize(c), Page: app.GetPage(c)}
	count, err := svc.CountArticle(&service.CountArticleRequest{Title: param.Title, Desc: param.Desc, Content: param.Content, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountArticle err :%v", err)
		response.ToErrorResponse(errcode.ErrorCountArticleFail)
	}

	articles, err := svc.GetArticleList(&service.ArticleListRequest{Title: param.Title, Desc: param.Desc, Content: param.Content, State: param.State}, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err:%v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
	}
	response.ToResponseList(articles, count)
}

// @Summary 添加文章
// @Produce  json
// @Param title body string true "文章标签" minlength(3) maxlength(100)
// @Param desc body  string true "文章简述" minlength(3) maxlength(100)
// @Param content body string false "文章内容"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err :%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateArticle err:%v", err.Error())
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
}

// @Summary 更新文章
// @Produce  json
// @Param id path int true "文章ID"
// @Param title body string false "文章标签" minlength(3) maxlength(100)
// @Param desc body  string false "文章简述" minlength(3) maxlength(100)
// @Param content body string false "文章内容"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
}

// @Summary 删除文章
// @Produce  json
// @Param id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs :%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
}
