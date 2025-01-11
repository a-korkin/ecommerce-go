package ports

type DBConnection interface {
	NewDBConnection(driver, conn string) (*DBConnection, error)
	CloseDBConnection() error
}
