# v0.2.0 最初のユースケース

- **状態**: 方針確定
- **マイルストーン**: `v0.2.0`

スキャフォールドの上に、次だけ実装する。`huh` は入れない（cobra の help と非対話で足りる）。

実装に入る直前にテスト仕様書 `docs/tests/v0.2.0.md` を書く（`--version` の自動テストも含む）。CI もこの版から入れる。

mac コマンドは Dev Container では動かさない。`GOOS=darwin` でクロスビルドし、**ホストの macOS で確認**する。未対応 OS ではエラー。

## 含むもの

### `slz setup`

利用者は、slz の利用に必要な設定ディレクトリを用意する。

```bash
slz setup
```

置く場所は [main.md](../../main.md) の XDG。無ければ次を作る。既にあれば何もしない（冪等）。`doctor` は走らせない。

- 設定ディレクトリ
- 空の `config.yaml`
- `commands/` ディレクトリ

### `slz config`

利用者は、設定を参照する。

```bash
slz config
```

設定ファイルの **中身** を標準出力に出す。無ければエラー（`slz setup` を案内）。`$EDITOR` で開く・キーの get/set は後回し。

### `slz --help`

利用者は、使えるコマンドの一覧と短い説明を見る。cobra 既定の help。

```bash
slz --help
```

### `slz mac dim`

利用者は、Mac をスリープさせずに画面だけ暗くする。

```bash
slz mac dim
```

`caffeinate` でスリープを抑え、`pmset displaysleepnow` で画面を消す。プロセスは **前面で待つ**（Ctrl+C で `caffeinate` が終わり、通常のスリープに戻る）。バックグラウンド化や専用の止めコマンドは後回し。macOS 以外ではエラー。

## 含まないもの

- `mac awake` / `dns-flush` など、dim 以外の mac コマンド
- `huh`、設定エディタ、doctor
- v0.3.0 以降に回すユースケース全般

----

以上
