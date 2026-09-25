# `slz setup`

設定ディレクトリを用意する。

```bash
slz setup
```

## 振る舞い

- 置き場は [main.md](../main.md) の XDG（`XDG_CONFIG_HOME/slz`、未設定なら `~/.config/slz`）
- 無ければ次を作る。既にあれば何もしない（冪等）
  - 設定ディレクトリ
  - 空の `config.yaml`
  - `commands/` ディレクトリ
- `doctor` は走らせない（`doctor-warned` も作らない）
- 余剰な位置引数はエラー
- `config.yaml` と同名のディレクトリがあるなど作れないときは終了コード非 0

検証: [tests/v0.2.0.md](../tests/v0.2.0.md)

----

以上
