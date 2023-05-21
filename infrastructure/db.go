package infrastructure

import "gorm.io/gorm"

type GormDatabase struct {
	db *gorm.DB
}

func NewGormDB(db *gorm.DB) *GormDatabase {
	return &GormDatabase{
		db: db,
	}
}

func (db *GormDatabase) Shutdown() error {
	conn, err := db.db.DB()
	if err != nil {
		return err
	}

	return conn.Close()
}

func (g *GormDatabase) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return g.db.Find(dest, conds...)
}

func (g *GormDatabase) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return g.db.First(dest, conds...)
}

func (g *GormDatabase) Create(value interface{}) *gorm.DB {
	return g.db.Create(value)
}

func (g *GormDatabase) Save(value interface{}) *gorm.DB {
	return g.db.Save(value)
}

func (g *GormDatabase) Delete(value interface{}) *gorm.DB {
	return g.db.Delete(value)
}

func (g *GormDatabase) Where(query interface{}, args ...interface{}) *gorm.DB {
	return g.db.Where(query, args...)
}

func (g *GormDatabase) Order(value interface{}) *gorm.DB {
	return g.db.Order(value)
}

func (g *GormDatabase) Limit(limit int) *gorm.DB {
	return g.db.Limit(limit)
}

func (g *GormDatabase) Offset(offset int) *gorm.DB {
	return g.db.Offset(offset)
}

func (g *GormDatabase) Preload(query string, args ...interface{}) *gorm.DB {
	return g.db.Preload(query, args...)
}

func (g *GormDatabase) Table(name string, args ...interface{}) *gorm.DB {
	return g.db.Table(name, args...)
}
