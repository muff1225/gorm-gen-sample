package main

import (
	"os"

	"github.com/muff1225/gorm-gen-sample/app/interface/db"
	"github.com/muff1225/gorm-gen-sample/app/interface/db/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./interface/db/query", // 出力パス
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})

	sqlHandler, err := db.NewSqlHandler(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	if err != nil {
		panic(err)
	}

	g.UseDB(sqlHandler.Conn)

	// Generate the code
	g.Execute()

	// 生成したModelにRelation情報を手動追加（これだけは手動対応が必要）
	allModels := []interface{}{
		g.GenerateModel(
			model.TableNameUser,
			gen.FieldRelateModel(field.HasOne, "PurchaseOrder", model.PurchaseOrder{}, nil),
		),
		g.GenerateModel(
			model.TableNamePurchaseOrder,
		),
	}

	g.ApplyBasic(allModels...)

	// Generate the code
	g.Execute()
}
