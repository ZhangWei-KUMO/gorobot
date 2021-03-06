# gorobot

![](https://img.shields.io/badge/go-v1.16.5-blue)

本文将会基于Golang官方教程教初学者如何编写一个自己的Go模块。话不多说，首先创建一个空文件夹`gorobot`

```bash
mkdir gorobot
cd gorobot
```

> 注意：在这里我使用的是自己的github仓库，开发者应按照实际情况对应自己仓库便于该包后续下载。

我们接着在根目录中创建模块的入口文件`main.go`,并写入简单的代码：

```bash
touch main.go
```

```go
// main.go
// 首先给这个包取个名称gorobot
package gorobot

import "fmt"
//  name string中的string为输入值类型，后一个string为输出值类型
func Hello(name string) string {
    // %v 为fmt包打印输出字符串占位符
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    // :=符号位go特有的初始化声明变量赋值方式，go语言对于变量不需进行类型声明。
    return message
}
```

我们给代码打上标签并提交至github仓库中。

```bash
git add -A
git commit -m "add initial code"
# 直接提交到master分支
git push origin master

# 如果是多人协作，则开启一个branch
git checkout -b v1
git push -u origin v1

# 给当前代码打上标签
git tag v1.0.1
git push --tag

```

## 测试模块是否已经成功提交并正常使用

创建模块的角度来看，我们已经完成了我们的所我们还是在当前根目录文件夹中创建一个`example\`文件夹,进入其中并创建一个`app.go`,写入如下代码：

```go
package main

import (
	"fmt"
	"github.com/ZhangWei-KUMO/gorobot"
)

func main() {
	message := gorobot.Hello("zhangwei")
	fmt.Println(message)
}
```

执行命令行`go mod init test`,给当前测试包创建一个`go.mod`文件，然后通过`go mod tidy`自动加载代码中所需的依赖。

> 注意：由于缓存的原因，go mod tidy可能会有代码版本的滞后，我们可以手动编辑`go.mod`文件指定版本或者执行go clean -modcache清除缓存。

然后我们运行代码:

```go
go run .
```

好了！您的第一个golang包成功了。

## 如何编写go语言的测试文件

对于go语言自身而言，它自带了一个轻量级的测试框架，其中包含`go test`命令行和`testing`包。在项目中我们的测试文件以`*_test.go`的方式命名。其中的测试函数以`Test`开头，参数类型为` *testing.T`，对于错误我们可以通过`t.Error()`或`t.Fail()`两种方式报错。