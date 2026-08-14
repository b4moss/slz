# 版未定・将来

- **状態**: 意図スタブ
- **マイルストーン**: `unscheduled`

v0.2.0 までに入れない、基盤・配布まわり。ユースケース本体は [usecases.md](./usecases.md)。v0.3.0 で再検討するものも含む。版が決まったら `plans/vX.Y.Z/` へ移す。

- キャッシュの実装（置く場所は [main.md](../../main.md) で `$XDG_CACHE_HOME/slz`）
- `slz doctor` と初回起動 warning
- `slz mv` / `slz git rm-branches`
- ユーザースクリプト向け `slz prompt`
- `commands/` のディレクトリネスト
- Linux 対応（初版は macOS のみ）
- Windows 対応（要望が出たら）
- Homebrew 等への配布（そのとき `release` ブランチをプロダクション相当にする。[git-rule.md](../../charter/git-rule.md)）

----

以上
