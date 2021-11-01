package model

import "gorm.io/gorm"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

// 单表操作
func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Where("id =? ", a.ID).Updates(values).Error
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id= ? AND state =? ", a.ID, a.State)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, err

}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? ", a.ID, 0).Delete(&a).Error
}

type ArticleRow struct {
	ArticleID     uint32
	TagID         uint32
	TagName       string
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
}

// 查询Row
func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title AS article_title", "ar.desc AS article_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+"AS at").
		Joins("LEFT JOIN ? AS t ON at.tag_id=t.id", Tag{}.TableName()).
		Joins("LEFT JOIN ? AS ar ON at.article_id=ar.id", Article{}.TableName()).
		Where("at.tag_id=? AND ar.state=?", tagID, a.State).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articles []*ArticleRow
	err = db.ScanRows(rows, &articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int64
	err := db.Table(ArticleTag{}.TableName()+"AS at").
		Joins("LEFT JOIN ? AS t ON at.tag_id=t.id", Tag{}.TableName()).
		Joins("LEFT JOIN ? AS ar ON at.article_id=ar.id", Article{}.TableName()).
		Where("at.tag_id=? AND ar.state=?", tagID, a.State).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
