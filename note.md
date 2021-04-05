# GeeOrm

## schema(对象与表之间的转换)

```go

表的格式:
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

例子:

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

----------------------->

schema.Name = "User"
schema.Fields[0].Name = "Name"
schema.Fields[0].Type = "string"
schema.Fields[0].Tag = "PRIMARY KEY"
```



## session(实现与数据库的交互)

数据库的连接与中断

增删查改



## 总体框架

```mermaid
graph LR
A[原生sql] --> B []
原生sql
```

