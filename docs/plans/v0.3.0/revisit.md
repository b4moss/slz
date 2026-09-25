# v0.3.0 以降

- **状態**: 意図スタブ
- **マイルストーン**: `v0.3.0`

## 含むもの

### `slz mac dim` 復帰時、常駐を解除する

現行の `slz mac dim`（[specs/mac.md](../../specs/mac.md)）は `caffeinate` を前面で待ち、止め方は Ctrl+C。

利用者が復帰したとき（画面が起きる / マシンに戻る）に常駐（`caffeinate`）を自動で解除する。通常のスリープに戻すのに Ctrl+C だけに頼らない。

`mac awake` は対象外。dim 本体の再実装もしない。

### `config.yaml` が空なら、`<blank>` という表示を出す

現行の `slz config`（[specs/config.md](../../specs/config.md)）はファイルの中身を出す。無ければ `slz setup` を案内する。

ファイルはあるが空のとき、リテラル `<blank>` を出す。何も出さない空の端末にはしない。

## その他

候補は [unscheduled/usecases.md](../unscheduled/usecases.md) と [unscheduled/future-intents.md](../unscheduled/future-intents.md)。版割当は未定。

----

以上
