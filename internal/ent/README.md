# ORM ent. 用法

## 1. 生成`schema`

`schema`对应MySQL里面的表。

导航到`internal`，示例生成两个schema`User`和`Pet`：

```shell
go run -mod=mod entgo.io/ent/cmd/ent new User Pet
```

将会生成`ent`目录。

在`ent/schema`目录对应文件加上字段，索引，关系等。

可以使用`atlas`工具查看SQL语句。

> Atlas: https://github.com/ariga/atlas

## 2. 生成资源文件

```shell
go generate ./ent
```

会生成数据库访问对象`Client`等。

## 3. 迁移（Migration）

把`schema`修改同步到数据库。

### 自动迁移

自动迁移建议仅用于开发和测试。

`main`里面添加如下函数：

```go
if err := client.Schema.Create(ctx); err != nil {
    log.Fatalf("failed creating schema resources: %v", err)
}
```

用上一步生成的`Client`对象执行数据迁移。

> 自动迁移更多选项：https://entgo.io/zh/docs/migrate

### 版本迁移

需要安装`atlas`。

- 生成迁移文件：`atlas migrate diff ...`
- 应用迁移：`atlas migrate apply ...`

> 详细文档：https://entgo.io/zh/docs/versioned-migrations

## 4. 查询

先创建第2步生成的`Client`对象，在用`Client`执行对数据库的CRUD：

```go
package main

import (
  "github.com/ilaziness/gintpl/internal/app/web"
  "github.com/ilaziness/gintpl/internal/ent"
  "github.com/ilaziness/gokit/storage/mysql"
  "github.com/ilaziness/gokit/log"
)

func main() {
  client := ent.NewClient(ent.Driver(mysql.EntDriver(web.Config.Db)))
  
  // 创建数据
  _, err := client.User.Create().SetName("user1").SetAge(20).SetUsername("username1").Save(context.Background())
  if err != nil {
    log.Logger.Errorf("failed creating user: %v", err)
  }
}
```
