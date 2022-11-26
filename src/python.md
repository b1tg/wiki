# run direct | python main.py

```py
import sys

def main():
   filepath = sys.argv[1]

if __name__ == '__main__':
    main() 
```


# md5 

```py
import hashlib
def md5(fname):
    hash_md5 = hashlib.md5()
    with open(fname, "rb") as f:
        for chunk in iter(lambda: f.read(4096), b""):
            hash_md5.update(chunk)
    return hash_md5.hexdigest()
```

# file

```py
outF = open("myOutFile.txt", "w") # write (create new)
outF = open("myOutFile.txt", "a") # append

# write

outF.write("abc")
outF.write("\n")
outF.writelines(["aa","bb"]) # 写很多行

# read
outF.read(1024) # read 1024 bytes
outF.read() # read all
outF.readlines() # TODO

# read line by line
for cnt, line in enumerate(fp):
    print("Line {}: {}".format(cnt, line))

# close to clean up
outF.close() # need to close


# use context open
with open(out_filename, 'w') as out_file:
    out_file.write("xx")

```

# error

```py
try:
  print("Hello")
except NameError:
  print("Variable x is not defined")
except:
  print("Something else went wrong")
else:
  print("Nothing went wrong") 
finally:
  print("The 'try except' is finished") 


```


# dir (directory)

```py
import os
for dir in os.listdir():
   pass

```

# subcommand

一个运行命令的「框架」，带超时，同时捕获 stdout and stderr

- `subcommand.run` 的第一个参数必须是数组
- timeout单位是秒
- stdout=PIPE, stderr=STDOUT的作用是把stderr重定向到stdout，stdout可以通过后续拿到
 
```py
import os
import subprocess
from subprocess import *
for line in outfile:
    try:
        i = i+1
        print("runing#{} {}".format(i,line))
        sr = subprocess.run(["rabin2", "-I", line[:-1]], timeout=4, stdout=PIPE, stderr=STDOUT)
        print(sr.stdout.decode("utf-8"))
    except subprocess.TimeoutExpired:
        print("timeout")
    except:
        print("unknow error")
```

# python2 subprocess

https://stackoverflow.com/a/4760517


期望信息：stdout，stderr，status
```
import subprocess
p = subprocess.Popen(['ls', '-a'], stdout=subprocess.PIPE, 
                                   stderr=subprocess.PIPE)
out, err = p.communicate()
status = p.returncode

```
