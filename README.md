# go-catering
## admin模块结构设计

> 管理后台的服务
>
> 技术栈采用 gorm gin redis来进行主要开发
>
> 主要是分成MVC三层架构进行代码设计

- 组合式API
- RESTful API

## web模块设计

> 管理后台
>
> 技术栈采用 vue vouter vuex 和 element-ui进行页面开发
>
> 代码风格具有promise和callback两种风格

### 待完成
- 前端图片的显示问题和查询删除按钮的逻辑完善
- Session携带token数据校验失败
- Dockerfile打包
- 自定义参数校验规则待完善
- Kafka事件服务的代码编写部署和协议的编写


## 开发日记

- 2022-03-06 完成区域、代金券、优惠卷、店铺的RESTfulApi，Service层和Model层的重构。
- 2022-03-09 完成用户积分，用户VIP，用户地址模块的RESTfulApi，Service层和Model层的重构。
- 2022-03-12 完成商品模块的RESTfulApi，Service层和Model层的重构。
- 2022-03-15 对无用代码进行删减
- 2022-03-16 修改数据库的配置问题
- 2022-03-17 修改参数校验规则，重新命名工具包，修改配置文件和结构体标签
- 2022-03-18 微服务的first commit
- 2022-03-20 管理后台的first commit
- 2022-03-21 增加session
- 2022-03-22 修改admin模块启动的逻辑和ip访问限制问题，增加注释和降低商铺和产品的粒度性
- 2022-03-24 修复admin模块路由重复注册导致启动失败和linux环境下启动失败的问题
- 2022-03-24 完成admin模块的Dockerfile打包和docker的部署
- 2022-03-25 完善商铺和产品的关联页面，修复更改gorm原始模型后部分功能失效问题
