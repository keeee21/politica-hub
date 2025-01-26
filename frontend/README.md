# politica-hub

## setup

``bash compose.sh up``

## minio

基本的には docker-comopse up をすれば起動する。
起動したあとは

- http://localhost:8900/login 　にアクセス
- docker-compose の MINIO_ROOT_USER、MINIO_ROOT_PASSWORD を入力してログイン
- fiels.localhost というバケットが作成されている
- 自分がアップロードした画像を探す

というような手順で画像が実際にアップロードされているかどうかの確認はできる。
