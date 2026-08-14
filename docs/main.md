# slz

`slz`（Super Laziness）は、**たまによく使う**コマンドを簡便に実行する CLI。ユーザーが自分のコマンドを登録することもできる。

## 目的

- よくあるが冗長な操作を、`slz <verb>` で短くする
- 破壊的な操作は、対象を解決して一覧し、確認してから実行する
- 公式の体験は Go が持つ。ユーザーの雑な短縮はシェルで足せる

例:

```bash
slz mv . ../
slz git rm-branches 'dev-*'
```

## ネーミング

- **正式名称**: slz（Super Laziness）
- **識別子**: `slz`

## フィロソフィー

- laziness は美徳。ただし危険な省略は標準化して安全側に倒す
- 対話は足りない項目だけ。全部ウィザードにはしない
- `fzf` / `gum` / `gh` / `aws` / `gcloud` を必須にしない。配るものは `slz` バイナリ 1 個
- パッケージマネージャにはならない

## 技術スタック

- Go **1.26** / cobra（ディスパッチャとビルトイン）
  - `go.mod` の言語バージョンは 1.26（2026-08 時点の安定メジャー）
  - パッチ（例: 1.26.6）は Dev Container のイメージや `toolchain` で追う。メジャー上げは意図して決める
- shell（ユーザーが追加するコマンド）
- 設定: YAML
- 開発: Dev Container（Linux。mac コマンドの確認はホスト）
- ライセンス: MIT

## 対応 OS

| 対象 | 方針 |
|------|------|
| macOS | 初版から対応 |
| Linux | いずれ対応 |
| Windows | 要望が出たら検討 |

## 開発環境

ローカルに Go を直接入れず、**Dev Container**（Linux）で開発する。イメージの Go は 1.26 系。コンパイルはコンテナ内。v0.1.0 の受け入れは [plans/v0.1.0/scaffold.md](./plans/v0.1.0/scaffold.md)。

mac コマンド（`mac dim` など）はコンテナでは動かさない。`GOOS=darwin` でクロスビルドし、**ホストの macOS で確認**する。コンテナ内や未対応 OS で mac コマンドを実行したらエラーでよい。

CLI の骨格は **cobra**（v0.1.0 から）。バージョン文字列は当面ハードコードする（`slz --version` → `slz version 0.1.0`）。

テスト仕様書は版ごとに `docs/tests/`（v0.1.0 は [tests/v0.1.0.md](./tests/v0.1.0.md)、v0.2.0 は [tests/v0.2.0.md](./tests/v0.2.0.md)）。自動テストの実装と CI は **v0.2.0 から**。CI は `develop` / `dev-v*` 向け PR のみ。

## 役割分担

Go は **ディスパッチャ兼、公式コマンドの実装**。シェルは **ユーザー拡張専用**。

| | ビルトイン | ユーザー追加 |
|--|--|--|
| 置き場 | Go のコード | 設定ディレクトリの `commands/` |
| 対話・色 | Go で持つ | 自分でやる（`read` で十分） |
| 依存 | `slz` だけ | そのスクリプトが呼ぶコマンド |

名前が衝突したら **ビルトインが勝つ**。同じ名前のユーザーコマンドがある場合は警告する。きれいな対話が必要になったら、ユーザースクリプトのまま伸ばさず **ビルトインへ昇格** する。

CLI なので、薄い DDD の層や CRUD Trait にそのまま準拠しなくてよい（[charter](./charter/README.md)）。思想は守る。

- 1 ユースケース 1 コマンド
- 入口（引数・フラグ）と、対象の解決・実行を分ける
- 永続化が無いので Repository / CRUD は置かない

詳細は [roadmap.md](./roadmap.md)。

## ディレクトリ構成

単一モジュール `github.com/b4moss/slz`。cobra はリポジトリ直下に置かず、`cmd/` はバイナリ入口、実装は `internal/`。

```text
slz/
├── cmd/slz/main.go              # 入口だけ。cobra を起動する
├── internal/
│   ├── cli/                     # cobra（入口・フラグ）
│   │   ├── root.go              # ディスパッチャ、初回 doctor warning
│   │   ├── mv.go                # slz mv
│   │   ├── cp.go
│   │   ├── git.go               # slz git …
│   │   ├── docker.go
│   │   ├── doctor.go
│   │   └── …
│   ├── prompt/                  # 確認・入力・TTY / -y / --dry-run
│   ├── config/                  # XDG、config.yaml、commands/ 探索
│   ├── doctor/                  # PATH 検査、スタンプ
│   ├── mv/                      # 対象解決と実行
│   ├── gitrm/                   # git rm-branches
│   └── usercmd/                 # ユーザースクリプトの exec
├── .devcontainer/
│   └── devcontainer.json        # Go 1.26 系
├── docs/
├── scripts/                     # ruleset 用
├── Makefile
├── go.mod
└── README.md
```

| 役割 | 置き場 |
|--|--|
| 入口 | `cmd/slz` |
| cobra（引数・フラグ） | `internal/cli`（1 ユースケース 1 ファイル、または 1 サブコマンド群） |
| 対象解決と実行 | `internal/mv` など（ドメイン名。`controllers` / `services` フォルダは作らない） |
| 共用 | `internal/prompt` / `config` / `doctor` |

- コードのテストは隣（`internal/mv/mv_test.go`）。テスト仕様書は `docs/tests/`（実装に入るときに足す）
- ユーザーコマンドはリポジトリではなく `$XDG_CONFIG_HOME/slz/commands/`
- `mac` / `port` / `s3` などは版が来てから `internal/cli/` と対応する実装を足す

v0.1.0 の最小（空なら `.gitkeep`。`--version` 用の入口と `go.mod`、Dev Container は実ファイル）:

```text
cmd/slz/main.go
internal/cli/.gitkeep
internal/prompt/.gitkeep
internal/config/.gitkeep
internal/doctor/.gitkeep
internal/mv/.gitkeep
internal/gitrm/.gitkeep
internal/usercmd/.gitkeep
.devcontainer/devcontainer.json
go.mod
```

## 対話

フラグがあれば非対話。なければ TTY 上で、足りない項目だけ聞く。全部をウィザードにはしない。

- パイプや CI（非 TTY）では聞かない。未指定ならエラーにする
- 破壊的な操作は **一覧 → 確認** を省略しない。`-y` だけが確認を飛ばす
- 対象の解決はシェルの glob に頼らず `slz` が行う

ビルトインの対話・色付けに `fzf` や `gum` は必須にしない。Go 側のライブラリ（例: `huh`）はビルド時の依存。ユーザースクリプト向けの `slz prompt` は後回し。

## 設定の置き場

[XDG Base Directory](https://specifications.freedesktop.org/basedir-spec/latest/) に従う。各変数が設定されていればそちら、なければ既定。`slz` 専用の環境変数は増やさない。

| 変数 | 未設定時 | 用途 |
|------|----------|------|
| `XDG_CONFIG_HOME` | `~/.config` | 設定とユーザーコマンド |
| `XDG_STATE_HOME` | `~/.local/state` | 初回 doctor warning 済みスタンプ |
| `XDG_CACHE_HOME` | `~/.cache` | キャッシュ（将来） |

```text
$XDG_CONFIG_HOME/slz/          # 未設定なら ~/.config/slz/
  config.yaml
  commands/
$XDG_STATE_HOME/slz/           # 未設定なら ~/.local/state/slz/
  doctor-warned
$XDG_CACHE_HOME/slz/           # 未設定なら ~/.cache/slz/
```

設定は YAML。ユーザーコマンドは `commands/` 直下の実行ファイルだけ（ファイル名 = サブコマンド名）。ディレクトリネストは最初は扱わない。

`slz setup` がこのディレクトリを用意する（[v0.2.0](./plans/v0.2.0/usecases.md)）。

## 外部コマンド

`git` / `gh` / `gcloud` / `aws` はバンドルしない。PATH の有無だけ見る。初回 warning と `slz doctor` の実装タイミングは v0.3.0 以降で再検討。方針の詳細は以前どおり（無いときは warning、そのコマンド実行時だけエラー）。

## これから / 対象外

- v0.1.0: スキャフォールド、Dev Container、`slz -v` → [roadmap.md](./roadmap.md)
- v0.2.0: `setup` / `config` / `--help` / `mac dim`
- v0.3.0: `mac dim` 復帰時の常駐解除、空 `config.yaml` の `<blank>` 表示。その他は v0.2.0 完了時に再検討
- 版未定のユースケース → [plans/unscheduled/usecases.md](./plans/unscheduled/usecases.md)
- キャッシュ実装、`slz prompt`、コマンドのディレクトリネスト → [plans/unscheduled/future-intents.md](./plans/unscheduled/future-intents.md)

## ドキュメントの読み方

- 目的・方針: 本ファイル
- これから: [roadmap.md](./roadmap.md) / [plans/](./plans/README.md)
- 守るルール: [charter/](./charter/README.md)（git は [git-rule.md](./charter/git-rule.md)）
- PO メモ: [wishlist.md](./wishlist.md)

----

以上
