package infrastructure

import (
	"github.com/jbactad/loop/infrastructure/repositories"
	"gorm.io/gorm"
)

type GormDatabase struct {
	db *gorm.DB
}

func (g *GormDatabase) Error() error {
	return g.db.Error
}

func NewGormDB(db *gorm.DB) *GormDatabase {
	return &GormDatabase{
		db: db,
	}
}

func (gdb *GormDatabase) Shutdown() error {
	conn, err := gdb.db.DB()
	if err != nil {
		return err
	}

	return conn.Close()
}

func (g *GormDatabase) Find(dest interface{}, conds ...interface{}) repositories.Database {
	return NewGormDB(g.db.Find(dest, conds...))
}

func (g *GormDatabase) First(dest interface{}, conds ...interface{}) repositories.Database {
	return NewGormDB(g.db.First(dest, conds...))
}

func (g *GormDatabase) Create(value interface{}) repositories.Database {
	return NewGormDB(g.db.Create(value))
}

func (g *GormDatabase) Save(value interface{}) repositories.Database {
	return NewGormDB(g.db.Save(value))
}

func (g *GormDatabase) Delete(value interface{}) repositories.Database {
	return NewGormDB(g.db.Delete(value))
}

func (g *GormDatabase) Where(query interface{}, args ...interface{}) repositories.Database {
	return NewGormDB(g.db.Where(query, args...))
}

func (g *GormDatabase) Order(value interface{}) repositories.Database {
	return NewGormDB(g.db.Order(value))
}

func (g *GormDatabase) Limit(limit int) repositories.Database {
	return NewGormDB(g.db.Limit(limit))
}

func (g *GormDatabase) Offset(offset int) repositories.Database {
	return NewGormDB(g.db.Offset(offset))
}

func (g *GormDatabase) Preload(query string, args ...interface{}) repositories.Database {
	return NewGormDB(g.db.Preload(query, args...))
}

func (g *GormDatabase) Table(name string, args ...interface{}) repositories.Database {
	return NewGormDB(g.db.Table(name, args...))
}
