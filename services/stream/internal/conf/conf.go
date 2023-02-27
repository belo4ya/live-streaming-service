package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

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

func Load(path string) (config.Config, *Bootstrap, error) {
	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)

	var b Bootstrap
	if err := c.Load(); err != nil {
		return nil, nil, err
	}
	if err := c.Scan(&b); err != nil {
		return nil, nil, err
	}

	return c, &b, nil
}
