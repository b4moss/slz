# slz

`slz`（Super Laziness）は、**たまによく使う**コマンドを簡便に実行する CLI。
ユーザーコマンドの登録は方針としてあるが、**現行（v0.2.0）では未実装**（`setup` が `commands/` を作るだけ）。

## 目的

- よくあるが冗長な操作を、`slz <verb>` で短くする
- 破壊的な操作は、対象を解決して一覧し、確認してから実行する（方針。出荷済コマンドに破壊的なものは無い）
- 公式の体験は Go が持つ。ユーザーの雑な短縮はシェルで足せる（方針）

将来の例（未実装）:

```bash
slz mv . ../
slz git rm-branches 'dev-*'
```

出荷済の振る舞いの正本は [specs/](./specs/README.md)。

## ネーミング

- **正式名称**: slz（Super Laziness）
- **識別子**: `slz`

## フィロソフィー

- laziness は美徳。ただし危険な省略は標準化して安全側に倒す
- 対話は足りない項目だけ。全部ウィザードにはしない（方針）
- `fzf` / `gum` / `gh` / `aws` / `gcloud` を必須にしない。配るものは `slz` バイナリ 1 個
- パッケージマネージャにはならない

## 技術スタック

- Go **1.25** / cobra（ディスパッチャとビルトイン）
  - `go.mod` の言語バージョンは 1.25
  - パッチは Dev Container のイメージや `toolchain` で追う。メジャー上げは意図して決める
- 設定ファイル: `config.yaml`（現行は空ファイルの作成と中身の表示のみ。YAML としての解釈はしない）
- 開発: Dev Container（Linux。mac コマンドの確認はホスト）
- ライセンス: MIT
- ユーザー拡張用シェル（`commands/` の exec）は [未実装](./plans/unscheduled/future-intents.md)

## 対応 OS

| 対象 | 現状 |
|------|------|
| macOS | 対応（`mac dim` 含む） |
| Linux | 対応（`setup` / `config` / `--version` 等。`mac` はエラー） |
| Windows | 要望が出たら検討 |

## 開発環境

ローカルに Go を直接入れず、**Dev Container**（Linux）で開発する。イメージの Go は 1.25 系。コンパイルはコンテナ内。

mac コマンド（`mac dim` など）はコンテナでは動かさない。`GOOS=darwin` でクロスビルドし、**ホストの macOS で確認**する。コンテナ内や未対応 OS で mac コマンドを実行したらエラーでよい。

CLI の骨格は **cobra**。バージョン文字列はハードコード（`slz --version` → `slz version 0.2.0`）。詳細は [specs/cli.md](./specs/cli.md)。

テスト仕様は `docs/tests/` を specs と同じドメイン切りで置く。自動テストと CI あり。CI は `develop` / `dev-v*` 向け PR のみ（`.github/workflows/test.yml`）。

## 役割分担

Go は **ディスパッチャ兼、公式コマンドの実装**。シェルによるユーザー拡張は **方針**（現行は未実装）。

| | ビルトイン（現行） | ユーザー追加（方針・未実装） |
|--|--|--|
| 置き場 | Go のコード | 設定ディレクトリの `commands/` |
| 対話・色 | 出荷済は非対話 | 自分でやる（`read` で十分）想定 |
| 依存 | `slz` だけ（`mac dim` は PATH の `caffeinate` / `pmset`） | そのスクリプトが呼ぶコマンド |

方針: 名前が衝突したら **ビルトインが勝つ**。同じ名前のユーザーコマンドがある場合は警告する。きれいな対話が必要になったら、ユーザースクリプトのまま伸ばさず **ビルトインへ昇格** する。

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
│   │   ├── root.go
│   │   ├── setup.go
│   │   ├── config.go
│   │   ├── mac.go
│   │   └── …
│   ├── config/                  # XDG_CONFIG_HOME、config.yaml、commands/
│   ├── mac/                     # mac dim
│   ├── prompt/                  # （将来・空）確認・入力
│   ├── doctor/                  # （将来・空）PATH 検査
│   ├── mv/                      # （将来・空）
│   ├── gitrm/                   # （将来・空）
│   └── usercmd/                 # （将来・空）ユーザースクリプト exec
├── .devcontainer/
│   └── devcontainer.json        # Go 1.25 系
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
| 対象解決と実行 | `internal/mac` など（ドメイン名。`controllers` / `services` フォルダは作らない） |
| 共用 | `internal/config`（出荷済）。`prompt` / `doctor` / `usercmd` 等はプレースホルダのみ |

- コードのテストは隣（例: `internal/mac/dim_test.go`、`internal/cli/*_test.go`）。テスト仕様書は `docs/tests/`
- ユーザーコマンド置き場は `$XDG_CONFIG_HOME/slz/commands/`（ディレクトリ作成のみ。exec は未実装）
- 未実装コマンドは版が来てから `internal/cli/` と対応する実装を足す

## 対話（方針）

出荷済コマンド（`--version` / `--help` / `setup` / `config` / `mac dim`）は **非対話**。

将来の破壊的コマンド向けの方針:

- フラグがあれば非対話。なければ TTY 上で、足りない項目だけ聞く。全部をウィザードにはしない
- パイプや CI（非 TTY）では聞かない。未指定ならエラーにする
- 破壊的な操作は **一覧 → 確認** を省略しない。`-y` だけが確認を飛ばす
- 対象の解決はシェルの glob に頼らず `slz` が行う

ビルトインの対話・色付けに `fzf` や `gum` は必須にしない。対話用に Go ライブラリ（例: `huh`）を入れる場合はビルド時依存にする（**現行の `go.mod` には未使用**）。ユーザースクリプト向けの `slz prompt` は後回し。

## 設定の置き場

[XDG Base Directory](https://specifications.freedesktop.org/basedir-spec/latest/) に従う方針。`slz` 専用の環境変数は増やさない。

### 現行（コードが読むもの）

| 変数 | 未設定時 | 用途 |
|------|----------|------|
| `XDG_CONFIG_HOME` | `~/.config` | 設定ディレクトリ（`config.yaml` / `commands/`） |

```text
$XDG_CONFIG_HOME/slz/          # 未設定なら ~/.config/slz/
  config.yaml                  # setup が空で作る。config がその中身を出す
  commands/                    # setup が作る。exec は未実装
```

`slz setup` がこのディレクトリを用意する（[specs/setup.md](./specs/setup.md)）。

### 予定（未使用）

| 変数 | 未設定時 | 用途 |
|------|----------|------|
| `XDG_STATE_HOME` | `~/.local/state` | 初回 doctor warning 済みスタンプ（[未実装](./plans/unscheduled/future-intents.md)） |
| `XDG_CACHE_HOME` | `~/.cache` | キャッシュ（[未実装](./plans/unscheduled/future-intents.md)） |

```text
$XDG_STATE_HOME/slz/           # 未設定なら ~/.local/state/slz/
  doctor-warned                # 予定。setup もコードも作らない
$XDG_CACHE_HOME/slz/           # 未設定なら ~/.cache/slz/
```

ユーザーコマンドは `commands/` 直下の実行ファイルだけ（ファイル名 = サブコマンド名）とする方針。ディレクトリネストは扱わない（[future-intents](./plans/unscheduled/future-intents.md)）。

## 外部コマンド

現行で起動するのは `mac dim` の `caffeinate` / `pmset` のみ（darwin）。

`git` / `gh` / `gcloud` / `aws` の PATH 検査・初回 warning・`slz doctor` は **未実装**（[plans/unscheduled/future-intents.md](./plans/unscheduled/future-intents.md)）。方針としては、無いときは warning、そのコマンド実行時だけエラー、とする。

## 出荷済 / これから / 対象外

- 出荷済（v0.1.0–v0.2.0）: スキャフォールド、`--version` / `--help`、`setup` / `config` / `mac dim` → [specs/](./specs/README.md)
- v0.3.0: `mac dim` 復帰時の常駐解除、空 `config.yaml` の `<blank>` 表示 → [plans/v0.3.0/revisit.md](./plans/v0.3.0/revisit.md)
- 版未定のユースケース → [plans/unscheduled/usecases.md](./plans/unscheduled/usecases.md)
- キャッシュ実装、`slz prompt`、ユーザーコマンド exec、コマンドのディレクトリネスト → [plans/unscheduled/future-intents.md](./plans/unscheduled/future-intents.md)

## ドキュメントの読み方

- 目的・方針: 本ファイル（pillar）
- OKF 索引: [index.md](./index.md)
- 現行の振る舞い: [specs/](./specs/README.md)
- テスト仕様: [tests/](./tests/README.md)
- これから: [roadmap.md](./roadmap.md) / [plans/](./plans/README.md)
- 守るルール: [charter/](./charter/README.md)（git は [git-rule.md](./charter/git-rule.md)）
- PO メモ: [wishlist.md](./wishlist.md)

----

以上
