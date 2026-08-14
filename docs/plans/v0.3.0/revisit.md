# v0.3.0 以降

- **状態**: 意図スタブ
- **マイルストーン**: `v0.3.0`

## 含むもの

### `slz mac dim` 復帰時、常駐を解除する

v0.2.0 の `slz mac dim` は `caffeinate` を前面で待ち、`pmset displaysleepnow` で画面を消す。止め方は Ctrl+C。

v0.3.0 では、利用者が復帰したとき（画面が起きる / マシンに戻る）に常駐（`caffeinate`）を自動で解除する。通常のスリープに戻すのに Ctrl+C だけに頼らない。

`mac awake`（画面を消さずスリープだけ止める）は対象外。dim 本体の再実装もしない。

### `config.yaml` が空なら、`<blank>` という表示を出す

v0.2.0 の `slz config` は `config.yaml` の中身を出す。ファイルが無ければ `slz setup` を案内する。

v0.3.0 では、ファイルはあるが空（出す中身が無い）とき、リテラル `<blank>` を出す。設定が存在して空だと分かるようにする。何も出さない空の端末にはしない。

## その他

**v0.2.0 完了時に再検討する。** いま版を割り当てない。候補は [unscheduled/usecases.md](../unscheduled/usecases.md) と [unscheduled/future-intents.md](../unscheduled/future-intents.md)。

----

以上
