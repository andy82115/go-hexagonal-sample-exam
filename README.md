[日本語](./docs/README-JP.md)

# How to install
1. Download Docker & Docker-Compose for postgresql container
2. Download Task 
```
//MacOs
brew install go-task/tap/go-task
```
3. Download Air for auto reload
```
//MacOs
`brew install air`
```
4. Copy .envsample and change name to .env . Adjust below param if needed
```
DB_USER="enter_db_user"

# Use Url at localstack request
AWS_URL="INPUT YOUR AWS APIGATEWAY BASE URL HERE"

TOKEN_SAVE_SECRET_AT_AWS="false"
```
5. Run Aws project if TOKEN_SAVE_SECRET_AT_AWS="true" (Optional)
[aws project for this](https://github.com/andy82115/aws-localstack-apigateway-sample-exam)
6. Start the docker of postgresql 
```
//with task
task docker-start
//manual
docker-compose up -d
```
7. Start the go project
```
//with task
task dev
//manual
//build binary and run it
```
8. Use swagger to test it (Optional)
```
http://127.0.0.1:8080/api-docs/index.html
```
9. Check db data in Docker if needed (Optional)
10. finished~~  :tada: :tada: :tada:

# Notice
1. The role will be decided when the first time Register API set (premium/normal)
2. Only Premium user can update user info and delete it 

# Connect swagger

1. swagger -> http://127.0.0.1:8080/api-docs/index.html


# Prerequisites

1. Task //Golang terminal supporter

2. docker & docker-compose

3. Air //Golang hot reload tool

4. dotenv