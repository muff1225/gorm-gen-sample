package main

import (
	"context"
	"fmt"
	"os"

	"github.com/muff1225/gorm-gen-sample/app/interface/db"
	"github.com/muff1225/gorm-gen-sample/app/interface/db/model"
	"github.com/muff1225/gorm-gen-sample/app/interface/db/query"
)

func main() {
	// MySQLへ接続
	sqlHandler, err := db.NewSqlHandler(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	if err != nil {
		fmt.Println(err.Error())
	}

	qu := query.Use(sqlHandler.Conn)

	ctx := context.Background()
	userQuery := qu.User
	purchaseOrderQuery := qu.PurchaseOrder

	// Insert
	// 生成されるSQL：INSERT INTO users SET name = 'sample'
	user := &model.User{
		Name: "sample",
	}
	err = qu.Transaction(func(tx *query.Query) error {
		if err := userQuery.WithContext(ctx).Create(user); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Select
	// 生成されるSQL：SELECT * FROM users WHERE name = 'sample' LIMIT 1;
	userResult, err := userQuery.WithContext(ctx).Where(userQuery.Name.Eq("sample")).First()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userResult)

	// Select
	// 生成されるSQL：SELECT * FROM users ORDER BY id DESC;
	userResults, err := userQuery.WithContext(ctx).Order(userQuery.ID.Desc()).Find()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userResults)

	// PurchaseOrderレコードを定義
	purchaseOrder := &model.PurchaseOrder{
		UserID: userResult.ID,
	}

	// Insert
	// 生成されるSQL：INSERT INTO purchase_orders SET user_id = 1;
	err = qu.Transaction(func(tx *query.Query) error {
		if err := purchaseOrderQuery.WithContext(ctx).Create(purchaseOrder); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Select（UserとPurchaseOrderを取得）
	// 生成されるSQL：SELECT * FROM users JOIN purchase_orders ON purchase_orders.user_id = users.id WHERE name = 'sample'
	purcahseOrderResults, err := purchaseOrderQuery.WithContext(ctx).Join(user, userQuery.ID.EqCol(purchaseOrderQuery.UserID)).Where(userQuery.Name.Eq("sample")).Find()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(purcahseOrderResults[0])
}
