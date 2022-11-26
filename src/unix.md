# frequently
## pip

```
-i https://pypi.douban.com/simple/
-i https://pypi.tuna.tsinghua.edu.cn/simple
```

## npm

```
--registry=https://registry.npm.taobao.org
```

## cargo

在 `$HOME/.cargo/config` 中添加如下内容：
```
[source.crates-io]
replace-with = 'ustc'

# 如果所处的环境中不允许使用 git 协议，可以把上述地址改为：

[source.ustc]
registry = "https://mirrors.ustc.edu.cn/crates.io-index"
# registry = "git://mirrors.ustc.edu.cn/crates.io-index"


```
## rustup (install rust)

```sh
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

## nvm

```sh
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.2/install.sh | bash
```

## ubuntu

  * [ ] [ubuntu-source](./sources/ubuntu-source.md)



## git
  * `git log -p main.c`
    看单个文件的git历史记录
