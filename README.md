# prac-go-sacloud-simplemq-api

さくらのクラウド シンプルMQとそのGo版APIライブラリを触ってみる

- [シンプルMQ | さくらのクラウド](https://cloud.sakura.ad.jp/products/simplemq/)
- [シンプルMQ | さくらのクラウド マニュアル](https://manual.sakura.ad.jp/cloud/appliance/simplemq/index.html)
- [シンプルMQ API](https://manual.sakura.ad.jp/api/cloud/simplemq/)
- [シンプルMQ さくらのクラウドAPI](https://manual.sakura.ad.jp/api/cloud/simplemq/sacloud/)

## ここにあるもの

> [!CAUTION]
> APIの認識を間違えたまま作ってしまっています。
>
> - receiveはデキューであり、削除はack
>
> 以下の内容とCLIは間違っています。(要修正)

https://github.com/sacloud/simplemq-api-go を使ってエンキュー・デキュー・ピーク・長さ取得を行うCLIを簡単に作った。\
ビルドは`make build`。各CLIの実行にはconfig.yamlの作成が必要。

### `enqueue`

キューに値を入れる

```
$ ./enqueue hoge fuga
[
  {
    "id": "019bbd01-4582-7ee1-bed8-030888a46053",
    "content": "hoge",
    "created_at": 1768402535810,
    "updated_at": 1768402535810,
    "expires_at": 1769612135810
  },
  {
    "id": "019bbd01-45c0-7312-ba3f-4bb346bf36ec",
    "content": "fuga",
    "created_at": 1768402535872,
    "updated_at": 1768402535872,
    "expires_at": 1769612135872
  }
]
```

### `dequeue`

キューから値を取り出す(参照 + 削除)

```
$ ./dequeue
[
  {
    "id": "019bbd01-4582-7ee1-bed8-030888a46053",
    "content": "hoge",
    "created_at": 1768402535810,
    "updated_at": 1768402581920,
    "expires_at": 1769612135810,
    "acquired_at": 1768402581920,
    "visibility_timeout_at": 1768402611920
  }
]
```

### `peek`

次にデキューされる値を出力する(参照のみ)

```
$ ./peek
[
  {
    "id": "019bbd01-45c0-7312-ba3f-4bb346bf36ec",
    "content": "fuga",
    "created_at": 1768402535872,
    "updated_at": 1768402629889,
    "expires_at": 1769612135872,
    "acquired_at": 1768402629889,
    "visibility_timeout_at": 1768402659889
  }
]
```

### `length`

キューの長さを取得する

```
$ ./length
1
```
