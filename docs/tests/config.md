# `slz config` テスト仕様

設定ファイル表示の検証。

関連仕様: [specs/config.md](../specs/config.md)

前提:

- 入口は `go run ./cmd/slz`（または同等のビルド成果物 `slz`）
- 設定の置き場は XDG。`XDG_CONFIG_HOME` があればその下の `slz/`、なければ `~/.config/slz/`

---

### Config

- 設定ファイルの中身を標準出力に出す（YAML としてパースしない）
- 無ければエラーにし、`slz setup` を案内する
- `$EDITOR` で開かない。キーの get/set はしない

#### テスト：正常系
- `config.yaml` があるとき、`slz config` の標準出力がそのファイルの中身と一致し、終了コード 0 である
- 空の `config.yaml` のとき、終了コード 0 である（中身は空でよい）
- `XDG_CONFIG_HOME` が設定されていれば、その下の `slz/config.yaml` を読む

#### テスト: 異常系
- 設定ファイルが無いとき、終了コードが 0 でない
- 設定ファイルが無いときの標準エラーに `slz setup` が含まれる
- `slz config extra` のように余剰な位置引数があるとき、終了コードが 0 でない

----

以上
