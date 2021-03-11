package schema

import (
	"go/ast"
	"reflect"

	"geeorm/dialect"
)

/*
type User struct {
    Name string `geeorm:"PRIMARY KEY"`
    Age  int
}

------>

CREATE TABLE `User` (`Name` text PRIMARY KEY, `Age` integer);

*/

// Field represents a column of database
type Field struct {
	Name string // 字段名
	Type string // 字段类型
	Tag  string // 约束条件;例:primary key
}

// Schema represents a table of database
type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// Parse: 解析
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// reflect.Indirect:获取指针的实例
	// 本处取得参数的结构体信息
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(), // 结构体的名词
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)

		// Anonyomous: 匿名的
		// IsExported: 报告名称是否为导出的 Go 符号（即，它是否以大写字母开头）。
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,

				// 进行对应数据库的类型转换
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))), // 进行对应数据库的类型转换
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}

// 该函数根据数据库中列的顺序,从对象中找到对应的值,依次平铺
// u1:= &User{Name:"Tom",Age:"18"} --> ("Tom",18)
func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}
