# gorm-gen-sample
このプロジェクトはQiitaの記事[GoのORMが大進化 2024年はgormのGenをはじめよう]()の補足用プロジェクトです。
gormのGenが動く環境を提供しています。

## 動かすには
### 実行に必要な環境
1. Docker

### 動かし方
プロジェクトのルートで以下実行してください。
docker-composeで必要なコンテナが起動します。
```bash
./runserver.sh
Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
[+] Running 3/3
 ⠿ Network gorm-gen-sample_default    Created                                                      0.1s
 ⠿ Container gorm-gen-sample-mysql-1  Started                                                      0.7s
 ⠿ Container gorm-gen-sample-app-1    Started                                                      1.1s
aa3fc46c431f:/go/src/github.com/muff1225/gorm-gen-sample/app# 
```

#### 初回起動時
マイグレーションを実行してください。
必要なテーブルをMySQLコンテナにマイグレーションします。
```bash
aa3fc46c431f:/go/src/github.com/muff1225/gorm-gen-sample/app# make migrate-up
1/u initial (73.843173ms)
```

#### 実行
ルート配下のmain.goにサンプルコードがあります。

```bash
aa3fc46c431f:/go/src/github.com/muff1225/gorm-gen-sample/app# go run main.go 

2023/12/18 21:35:28 /go/src/github.com/muff1225/gorm-gen-sample/app/interface/db/query/users.gen.go:342
[10.283ms] [rows:1] INSERT INTO `users` (`name`,`deleted_at`) VALUES ('sample',NULL)

2023/12/18 21:35:28 /go/src/github.com/muff1225/gorm-gen-sample/app/interface/db/query/users.gen.go:359
[2.101ms] [rows:1] SELECT * FROM `users` WHERE `users`.`name` = 'sample' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
&{1 sample 2023-12-18 21:35:28.305 +0000 UTC 2023-12-18 21:35:28.305 +0000 UTC {0001-01-01 00:00:00 +0000 UTC false} {0 0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC {0001-01-01 00:00:00 +0000 UTC false}}}

2023/12/18 21:35:28 /go/src/github.com/muff1225/gorm-gen-sample/app/interface/db/query/users.gen.go:383
[1.336ms] [rows:1] SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` DESC
[0xc0000d2600]

2023/12/18 21:35:28 /go/src/github.com/muff1225/gorm-gen-sample/app/interface/db/query/purchase_orders.gen.go:264
[46.592ms] [rows:1] INSERT INTO `purchase_orders` (`user_id`,`deleted_at`) VALUES (1,NULL)

2023/12/18 21:35:28 /go/src/github.com/muff1225/gorm-gen-sample/app/interface/db/query/purchase_orders.gen.go:305
[3.729ms] [rows:1] SELECT `purchase_orders`.`id`,`purchase_orders`.`user_id`,`purchase_orders`.`created_at`,`purchase_orders`.`updated_at`,`purchase_orders`.`deleted_at` FROM `purchase_orders` INNER JOIN `users` ON `users`.`id` = `purchase_orders`.`user_id` WHERE `users`.`name` = 'sample' AND `purchase_orders`.`deleted_at` IS NULL
&{1 1 2023-12-18 21:35:28.323 +0000 UTC 2023-12-18 21:35:28.323 +0000 UTC {0001-01-01 00:00:00 +0000 UTC false}}
aa3fc46c431f:/go/src/github.com/muff1225/gorm-gen-sample/app# 
```