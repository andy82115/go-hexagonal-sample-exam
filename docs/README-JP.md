# インストール方法
1. ダウンロード Docker & Docker-Compose。 （postgresql container）
2. ダウンロード Task 
```
//MacOs
brew install go-task/tap/go-task
```
3. ダウンロード Air。 （auto reload）
```
//MacOs
`brew install air`
```
4. .envsampleをコピーし、名前を.envに変更する。必要に応じて以下のパラメータを調整してください。
```
DB_USER="enter_db_user"

# Use Url at localstack request
AWS_URL="INPUT YOUR AWS APIGATEWAY BASE URL HERE"

TOKEN_SAVE_SECRET_AT_AWS="false"
```
5. TOKEN_SAVE_SECRET_AT_AWS=「true 」の場合、Awsプロジェクトを実行する (任意)
[aws project for this](https://github.com/andy82115/aws-localstack-apigateway-sample-exam)
6. postgresqlのドッカーを起動する。
```
//with task
task docker-start
//manual
docker-compose up -d
```
7. goプロジェクトを立ち上げる
```
//with task
task dev
//manual
//build binary and run it
```
8. swaggerを使ってテストする (任意)
```
http://127.0.0.1:8080/api-docs/index.html
```
9. 必要に応じてDockerでDBデータをチェックする (任意)
10. 完成~~  :tada: :tada: :tada:

# お知らせ
1. 役割は初回登録API設定時に決定 (premium/normal)
2. ユーザー情報の更新と削除ができるのはプレミアムユーザーのみです。

# swaggerを利用する場合

1. swagger -> http://127.0.0.1:8080/api-docs/index.html


# 前提条件

1. Task //Golang terminal supporter

2. docker & docker-compose

3. Air //Golang hot reload tool

4. dotenv