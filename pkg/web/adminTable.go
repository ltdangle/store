package web

import (
	"bytes"
	"context"
)

type TableRow map[string]string

type AdminTable struct {
	headings []string
	data     [][]string
}

func NewAdminTable() *AdminTable {
	return &AdminTable{
		headings: []string{"one", "two", "three", "four", ""},
		data: [][]string{
			{"oneData", "twoData", "threeData", "fourData", ""},
			{"oneData", "twoData", "threeData", "fourData", ""},
			{"oneData", "twoData", "threeData", "fourData", ""},
		},
	}
}

func (table *AdminTable) Render() string {
	var buf bytes.Buffer
	_ = AdminTbl(table.headings, table.data).Render(context.Background(), &buf)
	return string(buf.String())
}
