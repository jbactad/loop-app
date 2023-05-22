package repositories

//go:generate mockery --name=Database --output=./mocks --outpkg=mocks --with-expecter
type Database interface {
	Error() error
	Shutdown() error
	Find(dest interface{}, conds ...interface{}) Database
	First(dest interface{}, conds ...interface{}) Database
	Create(value interface{}) Database
	Save(value interface{}) Database
	Delete(value interface{}) Database
	Where(query interface{}, args ...interface{}) Database
	Order(value interface{}) Database
	Limit(limit int) Database
	Offset(offset int) Database
	Preload(query string, args ...interface{}) Database
	Table(name string, args ...interface{}) Database
}
