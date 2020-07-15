package storage

import (
	"context"
	"fmt"

	"github.com/ekozlova94/internal/ctxutils"
	"github.com/ekozlova94/internal/model"
)

func Save(data model.AccessTime, ctx context.Context) error {
	db := ctxutils.DbFromContext(ctx)
	if _, err := db.Exec("insert into access_time(ip, time) values ($1, $2)", data.Ip, data.Time); err != nil {
		return fmt.Errorf("got wrong while saving time: %w", err)
	}
	return nil
}

func Get(ip string, isAsc bool, ctx context.Context) (*model.AccessTime, error) {
	db := ctxutils.DbFromContext(ctx)
	var order = "desc"
	if isAsc {
		order = "asc"
	}
	result, err := db.Query("select time from access_time where ip = $1 order by id "+order+" limit 1", ip)
	if err != nil {
		return nil, fmt.Errorf("got wrong while saving time: %w", err)
	}
	//noinspection GoUnhandledErrorResult
	defer result.Close()
	for result.Next() {
		var m model.AccessTime
		err = result.Scan(&m.Time)
		if err != nil {
			return nil, fmt.Errorf("got wrong while saving time: %w", err)
		}
		return &m, nil
	}
	return nil, nil
}
