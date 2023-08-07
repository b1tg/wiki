
![](assets/2023-08-07-10-13-43.png)



Q: 如何修改内存？

这里的d是修改的宽度

```
ed 0xe750f50 42424242
```


Q: ghidra rebase program

To rebase a binary in Ghidra via the GUI: Window -> Memory Map -> Set Image Base button (far right house icon).
https://twitter.com/_argp/status/1167359147211251713


Q: ghidra 高亮变量

鼠标中键
https://github.com/NationalSecurityAgency/ghidra/issues/25



Q: 搜索内存：

0:013>  s -a 0 L?80000000 "defines Annotatio"
0ac44ce3  64 65 66 69 6e 65 73 20-41 6e 6e 6f 74 61 74 69  defines Annotati
0e63a0e2  64 65 66 69 6e 65 73 20-41 6e 6e 6f 74 61 74 69  defines Annotati
0e90414c  64 65 66 69 6e 65 73 20-41 6e 6e 6f 74 61 74 69  defines Annotati


Q: gflags

"D:\Windows Kits\10\Debuggers\x86\gflags.exe"  -i FoxitPDFReader.exe +hpa

实验：在win10上，搜索内存反查，失败，heap指令不生效就算用老版+hpa也不行


Q: 格式化打印

.printf "filename: %mu \n", poi(@esp+4)

- 宽字符串 %mu
- 字符串  %ma

和断点结合，需要对冒号和反斜杠进行转义

bp kernel32!CreateFileW ".printf \"filename: %mu \\n\", poi(@esp+4);kv;gc;"


Q: 条件断点之字符串

实例：打开某个特定文件名时断下
```
bp kernel32!CreateFileW "as /mu $FileName poi(@esp+0x4);.block{r @$t0=$scmp(@\"$FileName\", @\"G:\\clean\\2023-07\\2023-07-21-mesos\\mesos-master\\links_highlights_annots.pdf\");.if(0!=@$t0){gc;}}"
```



Q: 使用 drrun.exe 收集

```
G:\clean\2023-07\2023-07-cov-ghidra>D:\working\wafl-play\DynamoRIO-Windows-8.0.18752\bin64\drrun.exe -t drcov -- test_cov.exe

```






