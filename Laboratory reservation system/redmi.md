在Go语言的结构体中定义一个字段，用于存储类似于'YYYY-MM-DD HH:MM:SS'这样的时间格式，你可以使用Go语言的时间包（time package）和GORM的标签来实现。

以下是一个示例代码：

```go
import (
    "time"
)

type MyStruct struct {
    ID        uint      `gorm:"primary_key"`
    CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
    // 其他字段...
}
```

在这个示例代码中，结构体`MyStruct`中定义了一个`CreatedAt`字段，类型为`time.Time`。在GORM的标签中，我们使用`column`来指定数据库中对应的列名，并使用`type`来指定数据库中的数据类型为`datetime`。

这样，当你使用GORM进行数据库操作时，它会自动将`CreatedAt`字段的值转换为'YYYY-MM-DD HH:MM:SS'的格式进行存储。



