# 介绍

利用 go 语言本身写的一些工具，不涉及第三方模块

# File

文件处理模块

# Pathlib

路径操作工具，类似 python 中的 pathlib，集成了常用方法

```
p := pathlib.Path("D:\\Download")

err := p.Iter(func(path *pathlib.Handler, err error) error {
    fmt.Println(path)
    return err
})

if err != nil {
    log.Fatalln(err)
}
```