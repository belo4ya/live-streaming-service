package conf

type Bootstrap struct {
	Server *Server
	Data   *Data
}

type Server struct {
	Addr string
}

type Data struct {
	Database *Database
}

type Database struct {
	Dsn string
}
