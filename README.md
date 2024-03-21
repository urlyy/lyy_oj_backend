# 生成proto代码
1. 下载protoc，并配置环境变量[https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)
2. 
```shell
cd rpc 
protoc --go_out=./pkg/ --go-grpc_out=./pkg/  proto/judge.proto
```

# go开发参考资料
- [解决go数据表查询结构体对应字段null问题（sqlx converting NULL to string is unsupported）](https://blog.csdn.net/Ming13416908424/article/details/123748041?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522171032218416800182165601%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=171032218416800182165601&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduend~default-2-123748041-null-null.142^v99^pc_search_result_base9&utm_term=converting%20NULL%20to%20string%20is%20unsupported&spm=1018.2226.3001.4187)
- [vscode go关闭超链接跳转](https://blog.csdn.net/Apale_8/article/details/113922392)
- [Gin框架获取请求参数的各种方式详解](https://juejin.cn/post/7213176141462126653)
- [golang如何发送邮件（qq邮箱）](https://cloud.tencent.com/developer/article/2217677)
- [【Golang第11章：单元测试】GO语言单元测试](https://blog.csdn.net/weixin_45652150/article/details/128534305?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522171034076416800182168106%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=171034076416800182168106&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~rank_v31_ecpm-2-128534305-null-null.142^v99^pc_search_result_base9&utm_term=go%E8%BF%9B%E8%A1%8C%E6%A8%A1%E5%9D%97%E6%B5%8B%E8%AF%95&spm=1018.2226.3001.4187)
- [sqlx文档](https://jmoiron.github.io/sqlx/)
- [gin文档](https://gin-gonic.com/zh-cn/docs/examples/multipart-urlencoded-form/)
- [解决go gin框架 binding:"required"`无法接收零值的问题](https://www.cnblogs.com/rainbow-tan/p/15457818.html)
- [Go: How to get last insert id on Postgresql with NamedExec()](https://stackoverflow.com/questions/33382981/go-how-to-get-last-insert-id-on-postgresql-with-namedexec)

# Hint
- sqlx的Select和Get支持在如`SELECT id FROM ...`这样的语句时，只用传入&[]int{}或者&int，对于`Select`不会`err`，需要判断`len>0`，对于`Get`需要手动判断`err!=nil`  