package web

import (
	"bytes"
	"context"
	"store/pkg/repo"
)

type TableRow map[string]string

type AdminTable struct {
	headings []string
	data     [][]string
	DataMap  *repo.QueryToMapResult
}

func NewAdminTable() *AdminTable {
	return &AdminTable{
		headings: []string{"one", "two", "three", "four"},
		data: [][]string{
			{"oneData", "twoData", "threeData", "fourData"},
			{"oneData", "twoData", "threeData", "fourData"},
			{"oneData", "twoData", "threeData", "fourData"},
		},
	}
}

func (table *AdminTable) Render() string {
	var buf bytes.Buffer
	_ = AdminTbl(table.DataMap).Render(context.Background(), &buf)
	return string(buf.String())
}
