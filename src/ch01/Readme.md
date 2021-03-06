第一章 主要介绍了Go的环境安装
注意的点有以下几点
1、安装的Go的软件包
2、配置Go的环境变量: GoRoot 和 Path 路径

GoPath 在Go 1.12版本正在弱化，后面将会在Go语言中去掉

**为何选择 Golang**
<br>既然有很多其他编程语言可以做同样的工作，如 Python，Ruby，Nodejs 等，为什么要选择 Golang 作为服务端编程语言？

以下是我使用 Go 语言时发现的一些优点：

并发是语言的一部分（译注：并非通过标准库实现），所以编写多线程程序会是一件很容易的事。后续教程将会讨论到，并发是通过 Goroutines 和 channels 机制实现的。
Golang 是一种编译型语言。源代码会编译为二进制机器码。而在解释型语言中没有这个过程，如 Nodejs 中的 JavaScript。
语言规范十分简洁。所有规范都在一个页面展示，你甚至都可以用它来编写你自己的编译器呢。:smile:
Go 编译器支持静态链接。所有 Go 代码都可以静态链接为一个大的二进制文件（译注：相对现在的磁盘空间，其实根本不大），并可以轻松部署到云服务器，而不必担心各种依赖性。