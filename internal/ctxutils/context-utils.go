package ctxutils

import (
	"context"
	"database/sql"
	"net"
)

func NewDbContext(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, dbKey, db)
}

func DbFromContext(ctx context.Context) *sql.DB {
	return ctx.Value(dbKey).(*sql.DB)
}

func NewIpContext(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

func IpFromContext(ctx context.Context) string {
	return ctx.Value(userIPKey).(net.IP).String()
}
