package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

// RegisterDialect 和 GetDialect 两个方法用于注册和获取 dialect 实例。
// 如果新增加对某个数据库的支持，那么调用 RegisterDialect 即可注册到全局。
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
