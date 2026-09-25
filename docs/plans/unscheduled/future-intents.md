# 版未定・将来

- **状態**: 意図スタブ
- **マイルストーン**: `unscheduled`

出荷済に入れていない基盤・配布まわり。ユースケース本体は [usecases.md](./usecases.md)。版が決まったら `plans/vX.Y.Z/` へ移す。

- キャッシュの実装（置く場所は [main.md](../../main.md) で `$XDG_CACHE_HOME/slz`）
- `slz doctor` と初回起動 warning
- `slz mv` / `slz git rm-branches`
- ユーザースクリプト向け `slz prompt`
- `commands/` のディレクトリネスト
- Linux 対応（現行は macOS のみを対象）
- Windows 対応（要望が出たら）
- Homebrew 等への配布（そのとき `release` ブランチをプロダクション相当にする。[git-rule.md](../../charter/git-rule.md)）

----

以上
