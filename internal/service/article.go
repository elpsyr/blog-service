package service

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"` //2^32
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	State         uint8  `form:"state" binding:"oneof=0 1"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
}
type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"` //2^32
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	State         uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
