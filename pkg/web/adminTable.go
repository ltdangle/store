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
}

func NewAdminTable(entityName string) *AdminTable {
	return &AdminTable{entityName: entityName}
}

func (table *AdminTable) Render(router *AppRouter) string {
	var buf bytes.Buffer
	_ = AdminTbl(table.entityName, table.DataMap, router).Render(context.Background(), &buf)
	return string(buf.String())
}
