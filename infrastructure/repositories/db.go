package repositories

import "gorm.io/gorm"

type Database interface {
	Shutdown() error
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Order(value interface{}) *gorm.DB
	Limit(limit int) *gorm.DB
	Offset(offset int) *gorm.DB
	Preload(query string, args ...interface{}) *gorm.DB
	Table(name string, args ...interface{}) *gorm.DB
}
