## Prismaについて
### テーブルやカラムを変更するとき
1. schema.prismaをいじる
2. prisma migrate dev　←これでmigrateファイルの作成 & node_modules配下のファイルを更新

### 変更を取り入れるとき
1. pullしてくる。
2. prisma migrate deploy　←これでmigrateファイルの実行
3. prisma generate ←これでnode_modules配下のファイルを更新