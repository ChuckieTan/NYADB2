## TBM中对表的可见性管理
目前在TBM中, 并没有针对表进行可见性管理.
不管事务有没有结束, 一旦它创建了一张表, 该表将会立即被其他事务可见.
目前这样做的原因是为了简洁代码.

TODO: 后期将可见性的逻辑再次抽象一下, 并将TBM中的语义分析逻辑分离, 实现对表的可见性管理.

PS: 目前对TBM的代码抽象还不够满意, TBM应该实现三部分逻辑: 1)对语句进行语义解析, 2)管理表,字段,记录等结构, 3)管理表的可见性. 目前1)和2)被夹杂实现, 3)还未实现.


## 无Drop语句
添加Drop语句后, 就会涉及对表可见性的管理了, 为了简洁代码, 目前暂时未实现该条语句.

TODO: 增加Drop语句.

## Read无读取字段的筛选
如果执行Read name from student, 其结果实际上是执行Read * from student.
即没有对Read的读取字段做筛选.
目前没有做他的原因是为了简洁代码.

TODO: 对Read进行读取字段的筛选.

## Where只支持单字段, 且不能嵌套
目前Where文法较为简单, 见NYADB语法.

TODO: 扩展Where语法.

## 日志自动归档和压缩
TODO: 为日志文件增加自动归档和压缩功能.

## 增加Vacuum
TODO: 增加Vacuum功能.

## field.CalExp不能处理0和INF
TODO: 特殊处理这两种情况.
