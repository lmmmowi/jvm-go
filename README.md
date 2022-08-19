# jvm-go

![Language](https://img.shields.io/badge/language-go-9cf.svg)

## 项目介绍
该项目采用Go语言编写一个简化版本的Java虚拟机，主要用于学习JVM的内部实现原理和机制，同时也在过程中增进对Go语言的使用熟练度。

> 代码完全参考 [jvm.go](https://github.com/zxh0/jvm.go) 项目实现。

## 编译运行
```sh
go build github.com/lmmmowi/jvm-go/cmd/java
```

## JVM相关知识

### Class结构

### 字节码指令集
+ 常量指令【18】
    + nop
    + const系列（共15条）
    + push系列（共2条）
+ 加载指令【10】
    + iload系列（共5条）
    + lload系列（共5条）
+ 存储指令【5】
    + istore系列（共5条）
+ 栈操作指令【9】
    + pop系列（共2条）
    + dup系列（共6条）
    + swap
  
> 参考文档：https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-6.html#jvms-6.5