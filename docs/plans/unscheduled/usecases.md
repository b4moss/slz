# ユースケース（版未定）

- **状態**: 意図スタブ
- **マイルストーン**: `unscheduled`

コマンド名は決めた。詳細なサブコマンド・対話は、注記があるもの以外は後日決定する。版が決まったら `plans/vX.Y.Z/` へ移し、実装後は `docs/specs/` へ移す。

版が付いたもの:

- v0.2.0: `setup` / `config` / `--help` / `mac dim` → [v0.2.0/usecases.md](../v0.2.0/usecases.md)
- v0.3.0: `slz mac dim` 復帰時の常駐解除、空 `config.yaml` の `<blank>` 表示 → [v0.3.0/revisit.md](../v0.3.0/revisit.md)
- v0.3.0 以降のその他: [v0.3.0/revisit.md](../v0.3.0/revisit.md) で再検討

## `slz mac dim`

v0.2.0。正本は [v0.2.0/usecases.md](../v0.2.0/usecases.md)。dim 復帰時の `caffeinate` 解除は [v0.3.0](../v0.3.0/revisit.md)。

## `slz mac awake`

利用者は、画面は消さず、Mac のスリープだけ止める。`mac dim` の対。caffeinate 相当（暗くしない）。版未定。v0.3.0 には入れない。

```bash
slz mac awake
```

CLI のまま。メニューバー、画面オーバーレイ、タイトルバーの文字色は持たない（移植可能な手段がない）。それらが要るときは別配布の [`slz awake indicator`](#slz-awake-indicator)。

目印は、awake 中だけ端末のウィンドウ／タブの **タイトル文字列だけ** を上書きする（例: `slz awake`）。終了時は元のタイトルに戻す。タイトルに ANSI 色は入れない。iTerm2 専用のタブ色も使わない。前面占有だけでは足りない（`slz mac awake &` や Ctrl-Z で見えなくなる）。

timeout、通知センター、`--stop` などの追加フラグは持たない。`caffeinate` を前面で待つか裏で持つかは実装時に決める。

## `slz awake indicator`

（PO の綴り `inidicator` は誤記。名前は `slz awake indicator`。）

CLI バイナリには含めない。別配布。版未定。v0.3.0 には入れない。

macOS のスタンドアロン GUI アプリ。仕事は **awake がオンであることだけを示す**（メニューバー＝ツールバー／ステータス項目、および／または画面オーバーレイ）。`mac awake` が動いているときに使う。Amphetamine / Caffeine 型の付属であり、slz ディスパッチャそのものではない。

ふたつの製品は分ける。

1. `slz` CLI — `mac awake` と端末タイトル文字列
2. `slz awake indicator` — 任意の付属 `.app`

CLI と別配布。起動の結び方は未決。

## `slz mac dns-flush`

利用者は、Mac の DNS キャッシュを捨てる。対話は不要。

```bash
slz mac dns-flush
```

## `slz shell reload`

利用者は、現在のシェルをリロードする。

```bash
slz shell reload
```

`slz` は子プロセスのため、親シェルを直接 `exec` / `source` できない。実装時は親シェル側の仕組み（関数や `eval`）が必要になる。

## `slz port`

利用者は、指定ポートを使っているプロセスを調べる。既定は出すだけ。`--kill` のときだけプロセスを終了する（一覧 → 確認。`-y` で省略可）。

```bash
slz port 8080
slz port 8080 --kill
```

## `slz notify`

利用者は、macOS の通知を出す。`--` 以降が本文。

```bash
make && slz notify -- 終わった
slz notify -- done
```

## `slz cp`

利用者は、不可視ファイルを含めてコピーする。対象は slz が解決する（`mv` のコピー版）。

```bash
slz cp . ../
```

## `slz docker rm`

利用者は、Docker の container / volume / image、またはそれら全部を削除する。

```bash
slz docker rm
```

`prune` という語は使わない。サブコマンドと対話の振る舞いは後日決定する。

## `slz docker stop`

利用者は、動いている Docker container を止める。何を止めるか（対話）は後日決定する。

```bash
slz docker stop
```

## `slz s3 *`

利用者は、S3 互換ストレージに対して操作する（create / list / mount / sync）。

```bash
slz s3 …
```

裏は `rclone` が必要。詳細な振る舞いは後日決定する。

## `slz gcloud`

利用者は、GCP のサービスを立ち上げ・変更・破棄する。

```bash
slz gcloud
```

裏は `gcloud` が必要。詳細な振る舞いは後日決定する。

## `slz git remote add origin gh`

利用者は、ローカルリポジトリに GitHub 上の **既存 repo** を `origin` として足す。`gh` で候補を出して選ぶ。新規作成はしない。

```bash
slz git remote add origin gh
```

未決:

- `origin` が既にあるとき（エラー / `set-url` / 確認）
- SSH か HTTPS か（`gh` の既定に寄せてよいことが多い）
- 自分の repo だけか、org も出すか

## 未決（要検討）

### `slz git gone`

remote 側で消えた tracking branch をまとめて消す。`git rm-branches`（glob）とは入力が違う別ユースケース。確認は必須。採用してよいが、まだ版には入れない。

### `slz git wip` / git-fire 相当

ローカルの一時 commit と、push まで含む緊急退避は別物。一つの `wip` に push を足すと、日常の退避が remote に出る。分けるなら `slz git wip`（ローカル）と `slz git fire`（push まで）。

### `slz gh pr`

今の branch の PR をブラウザで開く、に閉じるなら採用してよい。`gh pr create` は `gh` 本体がすでに対話なので、slz が被せる価値は薄い。作る・一覧まで入れるかは後日。

## 見送り

- httpd の再起動 — 一旦見送り

----

以上
