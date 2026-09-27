# CLI 入口

ルートコマンド・バージョン・help の現行仕様。

## 入口

- バイナリ入口: `cmd/slz`
- ディスパッチャ: cobra（`internal/cli`）
- バージョン文字列はハードコード（`internal/cli.Version` = `0.2.0`）
- `--version` を使う。`-v` は対象外
- 既定の `completion` サブコマンドは無効（`CompletionOptions.DisableDefaultCmd`）

```text
slz version 0.2.0
```

- 標準出力に出す。標準エラーには出さない
- 引数なしの `slz` はバージョン行を出さない（終了コード 0）
- 未知のフラグ・サブコマンドは終了コード非 0

## Help

- cobra 既定の help（`slz --help`）
- 現行サブコマンド: `setup` / `config` / `mac`
- `slz help <topic>` で未知トピックは終了コード非 0
- help は設定ディレクトリを作らない。外部コマンドも起動しない
- 対話ライブラリ（`huh` 等）は依存に含めない

検証: [tests/cli.md](../tests/cli.md)

----

以上
