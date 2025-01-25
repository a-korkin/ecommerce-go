package handlers

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/a-korkin/ecommerce/internal/core/adapters/db"
	"github.com/pressly/goose/v3"
)

var router *Router
var connection *db.PostgresConnection

func migrate() {
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	migrationDir := filepath.Join(dir, "../../../migrations")
	if err := goose.Up(connection.DB.DB, migrationDir); err != nil {
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

insert into public.users(id, last_name, first_name)
values
	('4636a25d-02ee-4eb8-9757-efd677677076', 'Ivanov', 'Ivan'),
	('5e782875-4d9c-4641-be3c-afddeb05c083', 'Petrov', 'Petr'),
	('d3f729cb-43c0-40c4-9084-74fb2b0bd408', 'Sidorov', 'Sidr');
`
	if _, err := connection.DB.Exec(sql); err != nil {
		log.Fatal(err)
	}
}

// var users []*models.User

// func initUsers() {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	testFilesDir := filepath.Join(dir, "../../../test")
// 	usersDataFile := filepath.Join(testFilesDir, "users.json")
//
// 	file, err := os.Open(usersDataFile)
//
// 	if err = json.NewDecoder(file).Decode(&users); err != nil {
// 		log.Fatal(err)
// 	}
// }

func start() {
	var err error
	connection, err = db.NewDBConnection(
		"postgres",
		`
host=localhost port=5432 user=postgres password=admin
dbname=ecommerce_testdb sslmode=disable`)
	if err != nil {
		log.Fatal(err)
	}
	router = NewRouter(connection.DB, nil, "")
	migrate()
	prepareData()
	// initUsers()
}

func shutdown() {
	sql := `
delete from public.products;
delete from public.categories;
delete from public.users;`
	_, err := connection.DB.DB.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}

	if err := connection.CloseDBConnection(); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	log.Printf("start main testing...")
	start()
	exitCode := m.Run()
	shutdown()
	log.Printf("stop main testing...")
	os.Exit(exitCode)
}

func TestServeHTTP(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got: %v, want: %v",
			status, http.StatusOK)
	}
	bodyBytes, err := io.ReadAll(rr.Body)
	if err != nil {
		log.Fatal(err)
	}
	want := "hello from main router\n"
	body := string(bodyBytes)
	if body != want {
		t.Errorf("Wrong response body, got: %v, want: %v", body, want)
	}
}
