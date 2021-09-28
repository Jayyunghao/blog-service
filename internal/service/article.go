package service

type CountArticleRequest struct {
	Title string `form:"title"  binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleRequest struct {
	Title    string `form:"title" binding:"max=100"`
	Desc     string `form:"desc" binding:"max=255"`
	ImageUrl string `form:"cover_image_url" binding:"max=255"`
	Content  string `form:"content"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title    string `form:"title" binding:"max=100"`
	Desc     string `form:"desc" binding:"max=255"`
	ImageUrl string `form:"cover_image_url" binding:"max=255"`
	Content  string `form:"content"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreateBy string `form:"create_by" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID       uint32 `form:"id" binding:"required,gte=1"`
	Title    string `form:"title" binding:"max=100"`
	Desc     string `form:"desc" binding:"max=255"`
	ImageUrl string `form:"cover_image_url" binding:"max=255"`
	Content  string `form:"content"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifyBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
