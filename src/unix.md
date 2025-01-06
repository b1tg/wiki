# frequently
## pip

```
-i https://pypi.douban.com/simple/
```

```
-i https://pypi.tuna.tsinghua.edu.cn/simple
```

## go

启用 Go Modules 功能
```
go env -w GO111MODULE=on
```

七牛

```
go env -w  GOPROXY=https://goproxy.cn,direct
```

阿里云

```
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
```

goproxy
```
go env -w  GOPROXY=https://goproxy.io,direct
```

查看：

go env | grep GOPROXY

安装：

```
curl https://dl.google.com/go/go1.23.4.linux-amd64.tar.gz -O
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```


## npm

```
--registry=https://registry.npm.taobao.org
```

## cargo

在 `$HOME/.cargo/config` 中添加如下内容：


```
# ~/.cargo/config
[source.crates-io]
replace-with = 'rsproxy-sparse'
[source.rsproxy]
registry = "https://rsproxy.cn/crates.io-index"
[source.rsproxy-sparse]
registry = "sparse+https://rsproxy.cn/index/"
[registries.rsproxy]
index = "https://rsproxy.cn/crates.io-index"
[net]
git-fetch-with-cli = true

```

或者：

```
[source.crates-io]
replace-with = 'ustc'

# 如果所处的环境中不允许使用 git 协议，可以把上述地址改为：

[source.ustc]
registry = "https://mirrors.ustc.edu.cn/crates.io-index"
# registry = "git://mirrors.ustc.edu.cn/crates.io-index"


```


## install docker

```
curl -fsSL https://get.docker.com | bash -s docker
```



## rustup (install rust)

```sh
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

使用代理：

```
export RUSTUP_DIST_SERVER="https://rsproxy.cn"
export RUSTUP_UPDATE_ROOT="https://rsproxy.cn/rustup"
curl --proto '=https' --tlsv1.2 -sSf https://rsproxy.cn/rustup-init.sh | sh
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
