# `slz config`

設定ファイルの中身を標準出力に出す。

```bash
slz config
```

## 振る舞い

- `$XDG_CONFIG_HOME/slz/config.yaml`（未設定なら `~/.config/slz/config.yaml`）を読む
- ファイルがあれば中身をそのまま出す（空ファイルなら空でよい）
- 無ければ終了コード非 0。標準エラーに `slz setup` を案内する
- `$EDITOR` で開かない。キーの get/set はしない
- 余剰な位置引数はエラー

空ファイル時の `<blank>` 表示は未実装（[plans/v0.3.0/revisit.md](../plans/v0.3.0/revisit.md)）。

検証: [tests/v0.2.0.md](../tests/v0.2.0.md)

----

以上
