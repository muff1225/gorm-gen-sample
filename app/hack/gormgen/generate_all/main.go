package main

import (
	"os"

	"github.com/muff1225/gorm-gen-sample/app/interface/db"
	"gorm.io/gen"
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
	all := g.GenerateAllTable() // database to table model.

	g.ApplyBasic(all...)

	// Generate the code
	g.Execute()
}
