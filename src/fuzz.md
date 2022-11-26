
# winafl


# honggfuzz

google开源的fuzz，多进程多线程

git clone https://github.com/google/honggfuzz

sudo apt install binutils-dev libunwind-dev
make
sudo make install

hfuzz-gcc

hfuzz-gcc imgRead.c -o imgRead

- honggfuzz 主fuzz程序

```sh
Examples:
 Run the binary over a mutated file chosen from the directory. Disable fuzzing feedback (static mode):
  honggfuzz -i input_dir -x -- /usr/bin/djpeg ___FILE___ # -x 没有插桩、反馈
 As above, provide input over STDIN:
  honggfuzz -i input_dir -x -s -- /usr/bin/djpeg
 Use compile-time instrumentation (-fsanitize-coverage=trace-pc-guard,...):
  honggfuzz -i input_dir -- /usr/bin/djpeg ___FILE___ # 有插桩，需要前提是用hfuzz-gcc编译过的？
 Use persistent mode w/o instrumentation:
  honggfuzz -i input_dir -P -x -- /usr/bin/djpeg_persistent_mode # -P persistent fuzzing？？ 什么意思
 Use persistent mode and compile-time (-fsanitize-coverage=trace-pc-guard,...) instrumentation:
  honggfuzz -i input_dir -P -- /usr/bin/djpeg_persistent_mode
```

honggfuzz -i input/ -- ./imgRead __FILE__


- persistent fuzz [文档](https://github.com/google/honggfuzz/blob/master/docs/PersistentFuzzing.md)
对单个api进行fuzz，有两种方式：
1. 定义`int LLVMFuzzerTestOneInput(uint8_t *buf, size_t len)`函数，在这里面调用api
2. 在程序中获取输入（比如循环读用户输入）的地方用`HF_ITER(&buf, &len);`来获取fuzzer喂来的数据

这种模式集中了fuzz的火力，速度当然非常快。



# afl fuzz tcpdump
```sh
git clone https://github.com/the-tcpdump-group/tcpdump
cd tcpdump
./configure  # failed

# install libpcap
git clone https://github.com/the-tcpdump-group/libpcap

sudo apt install flex bison -y
cd libpcap && ./configure 

# 这里为什么要给两次fsanitize，值还不一样？？？
CC=afl-clang CFLAGS="-g -fsanitize=address -fsanitize=undefined -fno-omit-frame-pointer" LDFLAGS="-g -fsanitize-address -fsanitize=undefined -fno-omit-frame-pointer" ./configure

make
sudo make install

# finish install libpcap
cd ..
CC=afl-clang CFLAGS="-g -fsanitize=address -fsanitize=undefined -fno-omit-frame-pointer" LDFLAGS="-g -fsanitize-address -fsanitize=undefined -fno-omit-frame-pointer" ./configure
make
sudo make install
```

tcpdump 能处理.pcap文件，这是fuzz的点，很容易就值想到他能抓包而束手无策


精简corpus，tests文件夹下往往有样本可用
```
afl-cmin -i tests/ -o testsmain -m none -- ./tcpdump -vv -ee -nnr @@
```

开搞
```
afl-fuzz -i testsmin/ -o tcpdumpfuzz -m none -- ./tcpdump -vv -ee -nnr @@
```

## 如果用honggfuzz

https://youtu.be/9jqg7T3Ltn4

1. 编译阶段
CC=hfuzz 。。。

2. ~精简~, 不精简还是hgfuzz自带？？

honggfuzz -i tests -- ./tcpdump -vv -ee -nnr ___FILE___
