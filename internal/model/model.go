package model

import (
	"fmt"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type Model struct {
	ID         uint32         `gorm:"primary_key" json:"id"`
	CreatedBy  string         `json:"created_by"`
	ModifiedBy string         `json:"modified_by"`
	CreatedOn  time.Time      `json:"created_on"`
	ModifiedOn time.Time      `json:"modified_on"`
	DeletedOn  gorm.DeletedAt `json:"deleted_on,omitempty"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "blog_",
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.Logger = logger.Default.LogMode(logger.Info)
	}
	s, err := db.DB()
	if err != nil {
		return nil, err
	}
	s.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	s.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil

}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	m.CreatedOn = now
	m.ModifiedOn = now
	return
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	m.ModifiedOn = now

	return
}

func (m *Model) BeforeDelete(tx *gorm.DB) (err error) {

	return
}
