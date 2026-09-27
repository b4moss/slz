# `slz setup` テスト仕様

設定ディレクトリ用意の検証。

関連仕様: [specs/setup.md](../specs/setup.md)

前提:

- 入口は `go run ./cmd/slz`（または同等のビルド成果物 `slz`）
- 設定の置き場は XDG。`XDG_CONFIG_HOME` があればその下の `slz/`、なければ `~/.config/slz/`

---

### Setup

- XDG の設定ディレクトリを用意する
- 無ければディレクトリ、空の `config.yaml`、`commands/` を作る
- 既にあれば何もしない（冪等）
- `doctor` は走らせない

#### テスト：正常系
- 設定ディレクトリが無いとき、`slz setup` がディレクトリ・空の `config.yaml`・`commands/` を作り、終了コード 0 である
- 既にあるとき、`slz setup` は既存の `config.yaml` を上書きせず、終了コード 0 である
- `XDG_CONFIG_HOME` が設定されていれば、その下の `slz/` に作る（ホームの `~/.config/slz` には作らない）

#### テスト: 異常系
- `config.yaml` と同じ名前のディレクトリが既にあるなど、作れないときは終了コードが 0 でない
- `slz setup` は `doctor-warned` を作らない
- `slz setup extra` のように余剰な位置引数があるとき、終了コードが 0 でない

----

以上
