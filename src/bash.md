
# bash

## 循环做事
```sh
#TLDR

for i in {1..10}; do echo "abc"; done;

# example
for i in {1000..3000}; do for f in example.*; do zzuf -r 0.01 -s $i < "$f" > "$i-$f"; done; done


```

## stdout and stderr pipe to both terminal and file
```sh
sh any.sh 2>&1 | tee output.txt
sh any.sh 2>&1 | tee -a output.txt
```


## 上次命令的返回状态

```sh
echo $?
```
