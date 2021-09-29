package service

import (
	"Practice/go-programming-tour-book/blog-service/internal/model"
	"Practice/go-programming-tour-book/blog-service/pkg/app"
)

type CountArticleRequest struct {
	Title string `form:"title"  binding:"max=100"`
	Desc     string `form:"desc" binding:"max=255"`
	Content  string `form:"content"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleRequest struct {
	Title    string `form:"title" binding:"max=100"`
	Desc     string `form:"desc" binding:"max=255"`
//	ImageUrl string `form:"cover_image_url" binding:"max=255"`
	Content  string `form:"content"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title    string `form:"title" binding:"max=100"`
	Desc     string `form:"desc" binding:"max=255"`
//	ImageUrl string `form:"cover_image_url" binding:"max=255"`
	Content  string `form:"content"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreateBy string `form:"create_by" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID       uint32 `form:"id" binding:"required,gte=1"`
	Title    string `form:"title" binding:"max=100"`
	Desc     string `form:"desc" binding:"max=255"`
//	ImageUrl string `form:"cover_image_url" binding:"max=255"`
	Content  string `form:"content"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifyBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountArticle(param *CountArticleRequest) (int,error) {
	return svc.dao.CountArticle(param.Title, param.Desc, param.Content, param.State)
}

func (svc *Service) GetArticleList(param *ArticleRequest, pager *app.Pager) ([]*model.Article,error) {
	return svc.dao.GetArticleList(param.Title, param.Desc, param.Content, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.State, param.ModifyBy)
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}