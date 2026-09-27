# `slz mac` テスト仕様

`mac dim` の検証。

関連仕様: [specs/mac.md](../specs/mac.md)

前提:

- 入口は `go run ./cmd/slz`（または同等のビルド成果物 `slz`）
- `mac dim` の外部コマンドはテストではモックしてよい。実機確認はホストの macOS（仕様書の外）

---

### Dim

- macOS ではスリープを抑えつつ画面だけ暗くする
- `caffeinate` を前面で待ち、`pmset displaysleepnow` で画面を消す
- macOS 以外ではエラーにする
- バックグラウンド化しない

#### テスト：正常系
- darwin では `Start caffeinate` → `Run pmset displaysleepnow` → `Wait caffeinate` の順で呼ぶ（実プロセスはモックしてよい）
- darwin で外部コマンドが成功すれば終了コード 0 である
- `caffeinate` はバックグラウンド化せず、呼び出し側が完了を待つ

#### テスト: 異常系
- darwin 以外では終了コードが 0 でなく、`caffeinate` / `pmset` を起動しない
- darwin で `caffeinate` が失敗したとき、終了コードが 0 でなく、`pmset` を起動しない
- `slz mac dim extra` のように余剰な位置引数があるとき、終了コードが 0 でない
- 未知の `slz mac nosuch` の終了コードが 0 でない

----

以上
