package dao

import (
	"Practice/go-programming-tour-book/blog-service/internal/model"
	"Practice/go-programming-tour-book/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetTagById(id uint32) (model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Get(d.engine)
}

func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createBy},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}

	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}

func (d *Dao) CountArticle(title, desc, content string, state uint8) (int, error) {
	article := model.Article{Title: title, Desc: desc, Content: content, State: state}
	return article.Count(d.engine)
}

func (d *Dao) GetArticleList(title, desc, content string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: title, Desc: desc, Content: content, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetArticleById(id uint32) (model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Get(d.engine)
}

func (d *Dao) CreateArticle(title, desc, content string, state uint8, createBy string) error {
	article := model.Article{
		Title:   title,
		Desc:    desc,
		Content: content,
		State:   state,
		Model:   &model.Model{CreatedBy: createBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title, desc, content string, state uint8, modifiedBy string) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"title":       title,
		"desc":        desc,
		"content":     content,
		"modified_by": modifiedBy,
	}
	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
