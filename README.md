## 项目分包
分包如下：
```
|---cmd  //启动函数
|    |---main.go  
|    |---wire.go  
|    |---wire_gen.go  
|---etc  //配置相关
|    |---config.go  
|---global //全局变量引用，如数据库，kafka等
|    |---application.go
|    |---db.go
|    |---engine.go
|    |---httpserver.go
|    |---server.go
|---pkg  //公共包go常用封装
|    |---result.go  
|---internal  //具体业务
|    |---api  
|    |    |---handlers
|    |    |    |---pinghanler.go  
|    |    |---routers.go  
|    |    |---pingrouter.go  
|    |---services
|    |    |---
|    |---repos
|    |    |---
|    |---models

```
## 启动流程

## 开发步骤

## 依赖安装

### 

### ezr包

## 一些问题  

1. 下载ezr包的时候失败
```
go: ********/arch/agollo@v1.0.12: verifying go.mod: ********/arch/agollo@v1.0.12/go.mod: reading https://sum.golang.google.cn/lookup/********/arch/agollo@v1.0.12: 404 Not Found
        server response: not found: ********/arch/agollo@v1.0.12: unrecognized import path "********/arch/agollo": https fetch: Get "https://********/arch/agollo?go-get=1": dial tcp: lookup ******** on 8.8.8.8:53: no such host
```
解决方案：
```
go env -w GOPRIVATE="*.in"
```
2. 启动报错
```
..\internal\api\handlers\pinghandler.go:5:2: no required module provides package ********/godemo/internal/pkg; to add it:
        go get ********/godemo/internal/pkg
```
解决方案：
```
go run .\main.go .\wire_gen.go
```
3. wire依赖报错

解决方案：
```
go install github.com/google/wire/cmd/wire@latest
```
4. wire vscode告警
```
[{
	"resource": "/d:/workspace/godemo/cmd/wire.go",
	"owner": "_generated_diagnostic_collection_name_#0",
	"severity": 4,
	"message": "No packages found for open file D:\\workspace\\godemo\\cmd\\wire.go.\nThis file may be excluded due to its build tags; try adding \"-tags=<build tag>\" to your gopls \"buildFlags\" configuration\nSee the documentation for more information on working with build tags:\nhttps://github.com/golang/tools/blob/master/gopls/doc/settings.md#buildflags.",
	"source": "go list",
	"startLineNumber": 5,
	"startColumn": 9,
	"endLineNumber": 5,
	"endColumn": 13
}]
```

