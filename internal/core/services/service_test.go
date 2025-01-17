package services

import (
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/a-korkin/ecommerce/internal/core/adapters/db"
)

var categoryService *CategoryService
var productService *ProductService

func migrate() {
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	migrationDir := filepath.Join(dir, "../../../migrations")
	if err := goose.Up(categoryService.DB.DB, migrationDir); err != nil {
		log.Fatal(err)
	}
}

func prepareData() {
	sql := `
insert into public.categories(id, title, code)
values
	('688e64d3-c722-48e5-be96-850e419df2d6', 'category@1', 'cat@1'),
	('996be659-81f0-457c-8682-800abcfd64c2', 'category@2', 'cat@2'),
	('efa8b389-a3bd-4e06-84dd-4960a0dfc55b', 'category@3', 'cat@3');

insert into public.products(id, title, category, price)
values
	('5c0d6b4f-2d94-4e91-b69f-78f3832a810d', 'product@1', '688e64d3-c722-48e5-be96-850e419df2d6', 712.62),
	('fd3310fd-2101-445f-ad3d-216fda4bd8a2', 'product@2', '688e64d3-c722-48e5-be96-850e419df2d6', 86.21),
	('7ba8e565-82d9-4918-977c-85e62bc32e2c', 'product@3', '688e64d3-c722-48e5-be96-850e419df2d6', 23.31),
	('6d49ce02-5e08-4d95-9451-c40bb44966e1', 'product@4', '996be659-81f0-457c-8682-800abcfd64c2', 73.25),
	('85169049-293c-43e0-a0d9-327eeab730d4', 'product@5', 'efa8b389-a3bd-4e06-84dd-4960a0dfc55b', 66.50),
	('6022261d-c88f-4551-8419-7319eb3ce18f', 'product@6', '996be659-81f0-457c-8682-800abcfd64c2', 51.51),
	('c4b129bd-43cb-4922-85ba-210dbd120ac3', 'product@7', 'efa8b389-a3bd-4e06-84dd-4960a0dfc55b', 12.07),
	('66724666-13ed-47f7-b042-eab7694e7499', 'product@8', '996be659-81f0-457c-8682-800abcfd64c2', 37.88),
	('74958436-3427-4608-94b8-854c5db62e97', 'product@9', 'efa8b389-a3bd-4e06-84dd-4960a0dfc55b', 3.63);
	`
	_, err := categoryService.DB.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func starting() {
	conn, err := db.NewDBConnection(
		"postgres",
		`
host=localhost port=5432 user=postgres password=admin
dbname=ecommerce_testdb sslmode=disable`)
	if err != nil {
		log.Fatal(err)
	}
	categoryService = NewCategoryService(conn.DB)
	productService = NewProductService(conn.DB)
	migrate()
	prepareData()
}

func dropData() {
	sql := `
delete from public.products;
delete from public.categories;
	`
	_, err := categoryService.DB.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func shutdown() {
	dropData()
	if err := categoryService.DB.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	starting()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
