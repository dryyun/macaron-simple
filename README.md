# macaron-simple

作为 go 小白，尝试使用 macaron 构建 curd 应用   
学习使用 go，go-web 知识

###工具
> govendor - https://github.com/kardianos/govendor - Go 包管理工具  
> realize  - https://github.com/tockins/realize    - Go 工具，实现文件监控，热编译等  
> goose    - https://github.com/pressly/goose      - 数据结构迁移工具

###使用第三方库   
> macaron  - https://github.com/go-macaron/macaron - Go Web 框架   
> go-ini   - https://github.com/go-ini/ini         - 配置文件，读写 INI 文件的功能，macaron 本身就依赖这个库
> bcrypt   - https://godoc.org/golang.org/x/crypto/bcrypt - bcrypt hash   
> validator - https://github.com/go-playground/validator - 数据验证  
> binding  - https://github.com/mholt/binding  - html 表单数据转成 go 数据，跟 `validator` 配合代替了 `macaron` 自带的 [数据绑定与验证](https://go-macaron.com/docs/middlewares/binding) 功能

###需要环境
> mysql   
> redis  


###Database Migration 工具 

PS: 这里存在时区问题，数据库时间的时区是按照 UTC 来的  

> 创建名为 macaron_simple 的数据库  
> $ cd databse  
> $ goose mysql "root:123456@/macaron_simple?charset=utf8mb4" up  
>  




