# `slz setup`

設定ディレクトリを用意する。

```bash
slz setup
```

## 振る舞い

- 置き場は `XDG_CONFIG_HOME/slz`（未設定なら `~/.config/slz`）。`XDG_STATE_HOME` / `XDG_CACHE_HOME` は読まない
- 無ければ次を作る。既にあれば何もしない（冪等）
  - 設定ディレクトリ
  - 空の `config.yaml`
  - `commands/` ディレクトリ
- `doctor` は走らせない（`doctor-warned` も作らない）
- 余剰な位置引数はエラー
- `config.yaml` と同名のディレクトリがあるなど作れないときは終了コード非 0

検証: [tests/setup.md](../tests/setup.md)

----

以上
