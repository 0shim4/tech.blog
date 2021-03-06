---
title: "Go の型"
description: "Goの型についてまとめた記事"
date: 2022-06-19T00:00:00-00:00
lastmod: 2022-06-20T00:00:00-00:00
featured_image: ""
tags: []
disable_share: false
draft: false
---

goは静的型付け言語なのでコンパイル時に型の整合性が検証される。暗黙の型変換も行われないため意図しない型変換でのトラブルも遭遇しにくい。また、値から型が明白なら型の記述を省ける。

# 組み込み型の一覧

[A Tour of Go Basic types](https://go-tour-jp.appspot.com/basics/11)

|種類|型名|
|:-|:-|
|論理型|bool|
|文字列型|string|
|符号付き整数型|int, int8, int16, int32, int64, rune|
|符号なし整数型|uint, uint8, uint16, uint32, uint64, uintptr, byte|
|浮動小数点型|float32, float64|
|複素数型|complex64, complex128|
|エラー型|Error|

# コンポジット型

コンポジット型は、複数のデータを一つの集合としてあらわす型。
構造体(struct)、配列(array)、スライス(slice)、マップ(map)、チャネル(channel, chan)が存在する。

## 構造体

構造体は0個以上の変数を集合させたデータ構造を持つ。構造体内で漸減された変数はフィールドと呼ばれ、それぞれのフィールドは異なるデータを持つことができる。また、フィールドの方には組み込み型以外も適用できる。

```go
// フィールドを持たない構造体
var empty struct{}

// フィールドを持つ構造体
var point struct {
  ID struct
  x, y, int
}
```

## 配列

配列は同じ型のデータを集めて並べたデータ構造を持つ。

```go
// ゼロ値で初期化
var array [5]int

// 5つの要素を持つ配列を定義
arrayLiteral := [5]int{1, 2, 3, 4, 5}

// 配列のインデックスと値を指定
// インデックスの指定が無い箇所はゼロ値
arrayIndex := [...]int{2:1, 5:5, 7:13}
```

配列の型とデータの数は一度決めたら固定される。

## スライス

スライスは配列と同様に同じ型のデータを集めて並べたデータ構造を持つ。

スライスの定義

```go
// ゼロ値で初期化
var slice []int

// 5つの要素を持つスライスを定義
sliceLiteral := []int{1, 2, 3, 4, 5}
```

キーと値には別々の型を指定できる。キーの型には比較演算子による比較ができるものを指定する必要がある。

# 参考文献

[エキスパートたちのGo言語　一流のコードから応用力を学ぶ Software Design plus](https://gihyo.jp/book/2022/978-4-297-12519-6)