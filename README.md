# go-fsize

系统里那些零碎查询（文件大小、环境变量、随机串），终端里顺手就查了。

查看文件大小，零依赖。

## 用法

```bash
go run . somefile.bin          # 1.5 KiB 形式
go run . -b somefile.bin       # 直接打印字节数
go run . a.txt b.txt           # 可以一次看多个
```

按 1024 进制算（KiB/MiB），默认就给可读形式，`-b` 退回原始字节数。
