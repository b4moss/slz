# CLI 入口 テスト仕様

ルートコマンド・バージョン・help の検証。SemVer フォルダは使わない。

関連仕様: [specs/cli.md](../specs/cli.md)

前提:

- 入口は `go run ./cmd/slz`（または同等のビルド成果物 `slz`）
- バージョン文字列はハードコードの `0.2.0`
- cobra 既定の `--version` を使う。`-v` は対象外
- `huh` は使わない

---

### Execute

- cobra のルートコマンドを起動する
- `--version` が付いていればバージョンを標準出力に出して終了する
- `--help` が付いていれば cobra 既定の help を出して終了する

#### テスト：正常系
- `slz --version` の標準出力が `slz version 0.2.0`（末尾改行あり）である
- `slz --version` の終了コードが 0 である
- `slz --version` のとき標準エラーにバージョン文字列を出さない

#### テスト: 異常系
- 未知のフラグ（例: `slz --no-such-flag`）の終了コードが 0 でない
- 未知のサブコマンド（例: `slz nosuch`）の終了コードが 0 でない
- 引数なしの `slz` はバージョン行 `slz version 0.2.0` を標準出力に出さない

---

### Help

- cobra 既定の help を出す
- 使えるコマンドの一覧に、現行のサブコマンドが含まれる

#### テスト：正常系
- `slz --help` の終了コードが 0 である
- `slz --help` の標準出力に `setup` が含まれる
- `slz --help` の標準出力に `config` が含まれる
- `slz --help` の標準出力に `mac` が含まれる

#### テスト: 異常系
- `slz help nosuch` の終了コードが 0 でない
- `slz --help` は設定ディレクトリを作成しない
- `slz --help` は外部コマンド（`caffeinate` / `pmset`）を起動しない

----

以上
