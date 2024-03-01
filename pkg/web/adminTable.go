package web

import (
	"bytes"
	"context"
	"store/pkg/repo"
)

type TableRow map[string]string

type AdminTable struct {
	DataMap  *repo.QueryToMapResult
	router   *AppRouter
}

func NewAdminTable(router *AppRouter) *AdminTable {
	return &AdminTable{router: router}
}

func (table *AdminTable) Render() string {
	var buf bytes.Buffer
	_ = AdminTbl(table.DataMap, table.router).Render(context.Background(), &buf)
	return string(buf.String())
}
