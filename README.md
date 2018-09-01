# goTools
做一些go常用的工具
> InArray  判断某元素是否包含在另一数组元素中  其中类型无论

> OutputCsv  导出csv文件

> SubStr   取子串

> IsEqual  浮点数 指定精度的相等判断

### 预实现
1、对应test方法的实现

2、excel操作 读取,导出

3、mysql数据库操作的封装 ,基于 sqlx 包

   getOne(), 获取指定条件的一条数据

   getList(), 获取指定条件的多条数据, 可分页获取

   update(),  按指定条件更新数据

   del(),    按指定条件删除数据

   addOne(), 增加一条数据

   addBatch() 批量增加数据

