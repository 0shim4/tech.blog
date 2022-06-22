---
date: 2022-06-21 20:30
description: "Goのスライスについてまとめた記事です。"
featured_image: "images/Go1-640x371.jpg"
tags: []
title: "スライスの操作"
disable_share: false
draft: false
---
- [スライスとは](#スライスとは)
- [lenとcapでスライスの長さと容量を取得する](#lenとcapでスライスの長さと容量を取得する)
- [値の初期化](#値の初期化)
- [Slice Tricks](#slice-tricks)
	- [copyでスライスの複製を行う](#copyでスライスの複製を行う)
	- [appendでスライス同士を連結する](#appendでスライス同士を連結する)
	- [append, copyでスライス内の要素を削除する](#append-copyでスライス内の要素を削除する)
	- [スライスを逆順に並べ替える](#スライスを逆順に並べ替える)
	- [スライスの要素を偶数のみでフィルタリングする](#スライスの要素を偶数のみでフィルタリングする)
	- [スライスを任意の要素数に分割する](#スライスを任意の要素数に分割する)
- [参考文献](#参考文献)

# スライスとは

スライスは要素数の指定なく同じ型のデータを並べて初期化するデータ構造。

要素数を指定しないため配列よりもデータの扱いが柔軟。

# lenとcapでスライスの長さと容量を取得する

スライスは配列と同様に同じ型のデータを集めて並べたデータ構造を持つ型で、複数要素のリストを便利に扱える。スライス自体が内部でリストの長さや容量を管理しているため、組み込み関数`len`や`cap`を使って長さや容量を取得できる。また、`append`を使ってデータの追加ができる。

```go
src := []int{1, 2, 3, 4}
fmt.Println(src, len(src), cap(src))
// Output: [1 2 3 4] 4 4

src = append(src, 5)
fmt.Println(src, len(src), cap(src))
// Output: [1 2 3 4 5] 5 8
```

# 値の初期化

スライスのゼロ値は`nil`。

```go
slice := make([]int, 2, 3)
fmt.Println(slice, len(slice), cap(slice))
// Output: [0 0] 2 3

slice = []int{2: 1, 5:5, 7:13}
fmt.Println(slice, len(slice), cap(slice))
// Output: [0 0 1 0 0 5 0 13] 8 8

slice = []int{}
fmt.Println(slice, len(slice), cap(slice))
// Output: [] 0 0
```

# Slice Tricks

公式Wikiの[SliceTricks](https://github.com/golang/go/wiki/SliceTricks)にスライスを扱ううえで有用なテクニックがまとめられています。

## copyでスライスの複製を行う

```go
src := []int{1, 2, 3, 4, 5}

dst := make([] int, len( src))
copy(dst, src)

fmt. Println(dst, len( dst), cap( dst))
// Output: [1 2 3 4 5] 5 5
```

## appendでスライス同士を連結する

```go
src1, src2 := []int{1, 2}, []int{3, 4, 5}

// appendでスライスを連結する
dst := append(src1, src2...)
// [1 2 3 4 5]
```

## append, copyでスライス内の要素を削除する

```go
src := []int{1, 2, 3, 4, 5}

// 3番めの要素（3）を削除する
i := 2
dst := append(src[:i], src[i+1:]...)
// [1 2 4 5]

// append の 代わり に copy を 利用 する 場合
src = []int{ 1, 2, 3, 4, 5}
dst = src[: i + copy( src[ i:], src[ i + 1:])]
// [1 2 4 5]
```
## スライスを逆順に並べ替える

```go
src := []int{1, 2, 3, 4, 5}

// 方法1
for i := len(src)/2 - 1; i >= 0; i-- {
	opp := len(src) - 1 - i
	src[i], src[opp] = src[opp], src[i]
}
fmt.Println(src) // Output: [5 4 3 2 1]

// 方法2
for left, right := 0, len(src)-1; left < right; left, right = left+1, right-1 {
	src[left], src[right] = src[right], src[left]
}
fmt.Println(src)
// Output: [1 2 3 4 5] // 方法 1 で 逆順 に なっ て いる ので 元 に 戻っ た
```

## スライスの要素を偶数のみでフィルタリングする

```go
package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}

	// dst の 要素 は 空 だ が、 src と 同じ ポインタ を 指し て いる
	dst := src[:0]
	for _, v := range src {
		if even(v) {
			dst = append(dst, v)
		}
	}
	fmt.Println(dst)
	// Output: [2 4]

	// 次 の コード により src を ガベージコレクション に 回収 さ せる こと が できる
	for i := len(dst); i < len(src); i++ {
		// 要素 の 型 の ゼロ 値 を 代入 する（ nil など）
		src[i] = 0
	}
}

// 引数 が 偶数 か どう かを 判定 する 関数
func even(n int) bool {
	return n%2 == 0
}
```

## スライスを任意の要素数に分割する

```go
package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	size := 2
	dst := make([][]int, 0, (len(src)+size-1)/size)
	for size < len(src) {
		src, dst = src[size:], append(dst, src[0:size:size])
	}
	dst = append(dst, src)
	fmt.Println(dst)
	// Output: [[1 2] [3 4] [5]]
}
```

# 参考文献

[エキスパートたちのGo言語　一流のコードから応用力を学ぶ Software Design plus](https://gihyo.jp/book/2022/978-4-297-12519-6)
