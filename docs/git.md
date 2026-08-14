# Git / GitHub 戦略（本リポの上書き）

共通ルールは [charter/git-rule.md](./charter/git-rule.md) と [charter/versioning-rule.md](./charter/versioning-rule.md)。  
このファイルは **slz 固有の上書き**だけ書く。矛盾する場合は本ファイルを優先する。

CLI / パッケージ配信のため、charter の `staging` / `production`（Web の CD）は使わない。後で Homebrew 等から配るときは `release` をプロダクション相当とする。

最初は `develop` が正。`main` / `release` は成長に応じて足す。

CI は **v0.2.0 から**（v0.1.0 は手元スモークのみ）。GitHub Actions の導入タイミングは実装時に決める。

----

以上
