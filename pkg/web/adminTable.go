package web

import (
	"bytes"
	"context"
	"store/pkg/repo"
)

type TableRow map[string]string

type AdminTable struct {
	entityName string
	DataMap    *repo.QueryToMapResult
	router     *AppRouter
}

func NewAdminTable(entityName string, router *AppRouter) *AdminTable {
	return &AdminTable{entityName: entityName, router: router}
}

func (table *AdminTable) Render() string {
	var buf bytes.Buffer
	_ = AdminTbl(table.entityName, table.DataMap, table.router).Render(context.Background(), &buf)
	return string(buf.String())
}
