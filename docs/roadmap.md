# slz ロードマップ

詳細計画は [plans/](./plans/README.md)。現行仕様は [specs/](./specs/README.md)。開発ルールは [charter/](./charter/README.md)。PO メモは [wishlist.md](./wishlist.md)。

## 現状

現行リリースは **v0.2.0**（`internal/cli` の `Version`、タグ `v0.1.0` / `v0.2.0`）。プロダクト方針のハブは [main.md](./main.md)。

| 版 | 内容 | 状態 | 正本 |
| --- | --- | --- | --- |
| v0.1.0 | スキャフォールド、Dev Container、`slz --version` | 出荷済 | [specs/cli.md](./specs/cli.md) |
| v0.2.0 | `setup` / `config` / `--help` / `mac dim`、自動テストと CI | 出荷済 | [specs/](./specs/README.md) |
| v0.3.0 | `slz mac dim` 復帰時の常駐解除、空 `config.yaml` の `<blank>` 表示。その他は再検討 | 意図スタブ | [plans/v0.3.0/revisit.md](./plans/v0.3.0/revisit.md) |

版未定のユースケース: [plans/unscheduled/usecases.md](./plans/unscheduled/usecases.md)  
その他の将来: [plans/unscheduled/future-intents.md](./plans/unscheduled/future-intents.md)

----

以上
