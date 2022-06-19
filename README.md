# Tech blog

```sh
git clone --recursive https://github.com/0shima/tech.blog

hugo server -D

hugo # ビルド
```

# ディレクトリ構成

```
.
├── archetypes
├── config.toml
├── content
├── data
├── layouts
├── static
└── themes
```

## archetypes

`hugo new` でコンテンツファイルを作ったときの Markdown のフォーマットを指定できる。

```
++++
title: ファイル名と同じ
date: new コマンドを実行した日時
draft: true
++++
```

セクション名ごとに作成もできる。

```
├── ...
├── archetypes
│   ├── news.md
│   └── product.md
├── content
│   ├── news
│   │   └── 20191004_news.md
│   └── product
│       └── service_a.md
├── ...
```

`hugo new news/20191005_news` で作成したファイルには news.md テンプレートが適用される。

## assets

`hugo new site` したときには生成されない。 `Hugo Pipes` で処理するファイルを保存するためのディレクトリ。

## config.toml

[HUGOの設定ファイル](https://gohugo.io/getting-started/configuration/#all-configuration-settings)。TOML・YAML・JSON 形式で指定できる。

## config

`hugo new site` したときには生成されない。HUGOの設定ファイルが複数必要な場合このconfigディレクトリに設定ファイルを置く。

# content

記事ファイルを置くディレクトリ。contentの下にサブディレクトリを作ってブログのセクションを分けることができる。

```
├── ...
├── content
│   │   ├── news
│   │   └── 20191004_news.md
│   ├── product
│   └── blog
├── ...
```

## data

サイトの全ページから参照したいデータを記述したファイルを置くディレクトリ。config.tomlの [Params]に書くと{{ .Site.Params.variable }} で参照できる。

データベースとしてサイトから参照したい情報を設定ファイルに保持すると設定ファイルが煩雑化するので、dataの下に TOML・YAML・JSON 形式のファイルを配置してData Template として扱うことができる。 呼び出しは `{{ .Site.Data.ファイル名 }}` と記述する。

例えば次のようなJSONファイルを `./data/hitsong2019.json` に置くと、記事ファイルで参照できる。

```json
{
  "books": [
    {
      "Title": "Go言語によるWebアプリケーション開発",
      "Author": "Mat Ryer, 鵜飼 文敏他"
    },
    {
      "Title": "Go言語による並行処理",
      "Artist": "Katherine Cox-Buday, 山口 能迪"
    }
  ]
}
```

```md
<ul>
  {{ range  .Site.Data.hitsong2019.hits }}
    <li>{{ .Title }} - {{ .Artist }} </li>
  {{ end }}
</ul>
```

## layouts

themes においたテーマファイルを一部修正したいときやレイアウトパーツを追加したいときに使う。 修正したいテンプレートファイルと同じ構成・ファイル名でlayoutsの下に配置すると、そのレイアウトファイルだけ独自に変更できる。

例えば、記事ページ（single.html）のレイアウトを変更したい場合は次のように行う。

まず、テーマ/layouts/_default/single.html という構成でファイルが配置されている。プロジェクト直下でも同じように `./layouts/_default/single.html` に配置することで記事ページのレイアウトを上書きできます。

```
├── ...
├── layouts
│   └── _default
│       └── single.html # 上書きするレイアウトファイル
├── static
└── themes
    └── theme-a
        ├── ...
        ├── layouts
        │   ├── _default
        │   │   ├── baseof.html
        │   │   ├── list.html
        │   │   └── single.html # 上書きしたいレイアウトファイル
        │   ├── partials
        │   ├── 404.html
        │   └── index.html
        ├── ...
```

## static

サイト内の静的ファイル（スタイルシート、JavaScript や画像ファイルなど）を配置するディレクトリ。 記事に画像を貼る場合もこのディレクトリの下に配置する。

staticディレクトリに配置したディレクトリ・ファイルは、サーバーを立ち上げたときのルートディレクトリ（hugo コマンドでビルドしたときの publicディレクトリ）に同じ構成で配置される。

例えば、次のような構成でstaticディレクトリに配置する。

```
├── ...
├── static
│   ├─── img
│   │   └── image.png
│   └── me.png
├── ...
```

static 次のファイルは、それぞれ次のようにアクセスできます。


https://{your-site-url}/img/image.png

https://{your-site-url}/me.png