# DO

用于数据表与Dao数据映射

实现如下接口,可以自定义结构体对应的数据表表名
```go

type tabler interface {
	TableName() string
}

```