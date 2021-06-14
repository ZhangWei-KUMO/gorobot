# gorobot

本文将会基于Golang官方教程教初学者如何编写一个自己的Go模块。话不多说，首先创建一个空文件夹

```bash
mkdir go-greetings
cd go-greetings
```

然后我们需要使用`go mod init`命令行对根目录进行初始化，并创建go.mod文件。

```bash
go mod init github.com/ZhangWei-KUMO/gorobot
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
git tag v1.0.1
git push --tag
git checkout -b v1
git push -u origin v1
```

## 在其他项目中引入当前模块

我们还是在当前根目录文件夹中创建一个`test\`文件夹,进入其中并创建一个`app.go`,写入如下代码：

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

执行命令行`go mod init github.com/ZhangWei-KUMO/test`,在当前目录中会添加一个`go.mod`文件。