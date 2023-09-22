package db_conn

import (
	"context"
	"fmt"
	"net/url"
	"savannah/cmd/config"
	"savannah/internal/repository/db"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DatabaseConn(conf config.Conf) (db.Store, error) {
	dsnConn := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(conf.Database.User, conf.Database.Password),
		Host:     fmt.Sprintf("%s:%v", conf.Database.Address, conf.Database.Port),
		Path:     conf.Database.Name,
		RawQuery: "sslmode=disable",
	}
	poolConfig, err := pgxpool.ParseConfig(dsnConn.String())
	if err != nil {
		return nil, err
	}
	poolConfig.HealthCheckPeriod = 10 * time.Second
	poolConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return conn.Ping(ctx)
	}
	pool, errx := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if errx != nil {
		return nil, errx
	}
	return db.NewStore(pool), nil
}
