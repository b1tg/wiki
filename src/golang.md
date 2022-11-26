<!-- ---
title: Golang
subtitle: Golang
date: 2021-02-24
bigimg: [{src: "/primitives/img/unsplash-patrick-tomasso.jpg", desc: "Path"}]
--- -->

# string to int

Atoi (string to int) and Itoa (int to string).

```go
i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)
```

# go build x86 on x64

```sh
GOARCH=386 GOOS=windows go build
```