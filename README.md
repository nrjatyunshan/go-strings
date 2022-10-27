# 获取 Go 中的字符串

https://www.huaweicloud.com/zhishi/vss-017.html

文章里提到两种切割方法

## 进行反汇编，解析汇编指令，确定字符串起始位置和处理长度

文章中提到的反汇编可以参考,但无法直接是用.
在 linux 中指令更复杂,表示字符串的方式更多.

文章中根据三条 mov 指令即可确认字符串位置和长度.

现实中的场景更复杂, 第一种场景类似文章中的描述 lea mov lea 三条指令可以确定字符串位置和长度.

```txt
  48ae30:       48 8d 05 12 9d 01 00    lea    0x19d12(%rip),%rax        # 4a4b49 <go.string.*+0x29b1>
  48ae37:       bb 13 00 00 00          mov    $0x13,%ebx
  48ae3c:       48 8d 4c 24 28          lea    0x28(%rsp),%rcx
```

第二种场景. 这 4 条指令实际上是把一个字符串作为参数放到栈上. 4c27f8 保存了一个类似 Go 字符串的结构体, 这个地址指向的内存里保存了字符串的地址和长度.

```txt
  48ae00:       48 8d 15 39 76 00 00    lea    0x7639(%rip),%rdx        # 492440 <type.*+0x7440>
  48ae07:       48 89 54 24 28          mov    %rdx,0x28(%rsp)
  48ae0c:       48 8d 15 e5 79 03 00    lea    0x379e5(%rip),%rdx        # 4c27f8 <runtime.defaultGOROOT.str+0x70>
  48ae13:       48 89 54 24 30          mov    %rdx,0x30(%rsp)
```

难点: 如何在没有反汇编工具的情况下,通过分析 ELF 字节流确定出上述两种场景

## 直接根据Go语言字符串组织原则进行切割

测试发现这个规则可能会被违反.

```golang
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
}

```

5a5b0a5d 违法了上述规则. 多个字节的字符串也会出现 0a 填充的情况.

```txt
objdump -s hello --start-address=0x4973f8 --stop-address=0x497428

hello:     file format elf64-x86-64

Contents of section .rodata:
 4973f8 2528292b 2c2d2e2f 30353a3c 3d3f434c  %()+,-./05:<=?CL
 497408 4d505355 5a5b0a5d 60686d73 7b7d202b  MPSUZ[.]`hms{} +
 497418 20402050 205b2822 29202928 290a2c20   @ P [(") )().,
```

难点: 确定 0a 的产生原因,尝试绕过.


