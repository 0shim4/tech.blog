---
title: "Goのモジュール作成"
description: "Goのモジュールを作成して実行する記事"
date: 2022-06-24T00:00:00-00:00
lastmod: 2022-06-24T00:00:00-00:00
featured_image: ""
tags: []
disable_share: false
draft: false
---

# はじめに

モジュールの作成方法と作成したモジュールの実行方法を紹介します。

本記事ではプログラムを開始するhelloモジュールと、mainモジュール内で実行するgreetingsモジュールを作成します。

# greetings モジュールの作成

## greetings モジュールを配置するディレクトリを作成

```bash
mkdir greetings
cd greetings
```

## greetings モジュールの初期化

[`go mod init`](https://go.dev/ref/mod#go-mod-init)コマンドでカレントディレクトリをルートとする新しいモジュールの初期化をします。

```bash
go mod init example.com/greetings
# go: creating new go.mod: module example.com/greetings
```

## greetings モジュールの実装

Go言語におけるアクセス制限に`public`や`private`のような修飾子はなく、頭文字が大文字の場合にアクセス可・小文字の場合にはアクセス不可となります。

greetings/greetings.goファイルを作成して以下の様に編集します。

```
作業ディレクトリ
└ greetings
  ├ go.mod
  └ greetings.go
```

```go
package greetings

import "fmt"

// Helloは指名された人への挨拶を返します。
func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

# プログラムを開始する hello モジュールの作成

## hello モジュールを配置するディレクトリを作成

作業ディレクトリ下にhelloディレクトリを作成します。

```
作業ディレクトリ
├ greetings
│ ├ go.mod
│ └ greetings.go
└ hello
```

```bash
cd ../
mkdir hello
cd hello
```

## hello モジュールの初期化

こちらも同じく[`go mod init`](https://go.dev/ref/mod#go-mod-init)コマンドを実行します。

```bash
go mod init example.com/hello
# go: creating new go.mod: module example.com/hello
```

## mainモジュールの実装

hello/main.goファイルを作成して以下の様に編集します。

※プログラムを開始するファイルのためファイル名をmain.goとしています。

```
作業ディレクトリ
├ greetings
│ ├ go.mod
│ └ greetings.go
└ hello
  ├ go.mod
  └ main.go
```

```go
package main

import {
  "fmt"

  "example.com/greetings"
}

func main() {
  message := greetings.Hello("Gladys")
  fmt.Println(message)
}
```

## ローカルに実装された greetings モジュールの参照

example.com/greetingsは公開されていないため、importされたときにローカルの../greetingsを参照するようgo.modを変更する[`go mod edit`](https://go.dev/ref/mod#go-work-edit)コマンドを実行します。

```bash
go mod edit -replace example.com/greetings=../greetings
```

## モジュールの依存関係を整理

[`go mod tidy`](https://go.dev/ref/mod#go-mod-tidy)コマンドを実行します。ソースコードに不足していたり、使われていないモジュールを自動的に追加・削除してくれるコマンドです。
ちなみに、tidyは「几帳面」という意味だそうです。

```bash
go mod tidy
# go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000
```

## 実行

```bash
go run main.go
# Hi, Gladys. Welcome!
```

# 参考

[Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)