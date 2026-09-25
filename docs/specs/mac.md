# `slz mac`

macOS 向けヘルパー。現行は `dim` のみ。

## `slz mac dim`

スリープさせずに画面だけ暗くする。

```bash
slz mac dim
```

## 振る舞い

- `caffeinate` を前面で起動し、続けて `pmset displaysleepnow` で画面を消す
- プロセスは前面で待つ（Ctrl+C で `caffeinate` が終わり、通常のスリープに戻る）
- バックグラウンド化しない。専用の止めコマンドは無い
- macOS（`darwin`）以外ではエラー。外部コマンドは起動しない
- 余剰な位置引数・未知の `mac` サブコマンドはエラー

復帰時の常駐解除は未実装（[plans/v0.3.0/revisit.md](../plans/v0.3.0/revisit.md)）。

検証: [tests/v0.2.0.md](../tests/v0.2.0.md)。実機確認はホストの macOS。

----

以上
