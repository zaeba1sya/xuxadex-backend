package db

import (
	"fmt"
	"time"

	"gitlab.com/xyxa.gg/backend-mvp-main/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBClient struct {
	client      *sqlx.DB
	cfg         *config.Config
	isConnected bool
}

func New(cfg *config.Config) *DBClient {
	return &DBClient{
		isConnected: false,
		cfg:         cfg,
		client:      nil,
	}
}

func (d *DBClient) Connect() error {
	var db *sqlx.DB
	var sslMode string = "disable"
	var err error

	if d.cfg.DB.SSL {
		sslMode = "require"
	}

	switch d.cfg.DB.Dialect {
	case "postgres":
		db, err = sqlx.Open("postgres", fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			d.cfg.DB.Host, d.cfg.DB.Port, d.cfg.DB.User, d.cfg.DB.Password, d.cfg.DB.Name, sslMode,
		))
	default:
		return fmt.Errorf("unknown db type %s", d.cfg.DB.Dialect)
	}

	if err != nil {
		return err
	}

	db.SetConnMaxIdleTime(time.Duration(d.cfg.DB.Pool.ConnMaxLifetime))
	db.SetMaxIdleConns(d.cfg.DB.Pool.MaxIdleConns)
	db.SetMaxOpenConns(d.cfg.DB.Pool.MaxOpenConns)

	d.isConnected = true
	d.client = db

	return nil
}

func (d *DBClient) MustConnect() {
	for {
		if err := d.Connect(); err != nil {
			fmt.Println("Failed to connect to Database. Retrying in 3 seconds...")
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}
}

func (d *DBClient) GetIsConnected() bool {
	return d.isConnected
}

func (d *DBClient) GetClient() *sqlx.DB {
	return d.client
}

func (d *DBClient) Release() error {
	if d.GetIsConnected() {
		return d.client.Close()
	}

	return nil
}
