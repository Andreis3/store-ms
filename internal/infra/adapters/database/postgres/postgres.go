package postgres

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/stores-ms/internal/infra/common/configs"
)

type Postgres struct {
	db *pgxpool.Pool
}

func NewPostgresDB(conf configs.Conf) *Postgres {
	var singleton sync.Once
	var pool *pgxpool.Pool
	singleton.Do(func() {
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName)
		maxConns, _ := strconv.Atoi(conf.MaxConnections)
		minConns, _ := strconv.Atoi(conf.MinConnections)
		maxConnLifetime, _ := strconv.Atoi(conf.MaxConnLifetime)
		maxConnIdleTime, _ := strconv.Atoi(conf.MaxConnIdleTime)

		connConfig, err := pgxpool.ParseConfig(connStr)

		if err != nil {
			panic(err)
		}

		connConfig.MinConns = int32(minConns)
		connConfig.MaxConns = int32(maxConns)
		connConfig.MaxConnLifetime = time.Duration(maxConnLifetime) * time.Minute
		connConfig.MaxConnIdleTime = time.Duration(maxConnIdleTime) * time.Minute
		connConfig.HealthCheckPeriod = 10 * time.Minute

		connConfig.ConnConfig.RuntimeParams["application_name"] = "store-ms"

		pool, err = pgxpool.NewWithConfig(context.Background(), connConfig)

		if err != nil {
			panic(err)
		}
	})
	return &Postgres{db: pool}
}
func (p *Postgres) InstanceDB() any {
	return p.db
}

func (p *Postgres) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	return p.db.Exec(ctx, sql, arguments...)
}
func (p *Postgres) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return p.db.Query(ctx, sql, args...)
}
func (p *Postgres) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return p.db.QueryRow(ctx, sql, args...)
}
func (p *Postgres) Close() {
	p.db.Close()
}
