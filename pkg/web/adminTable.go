package web

import (
	"bytes"
	"context"
	"fmt"
	"strings"
)

type TableRow map[string]string

type AdminTable struct {
	headings []string
	rows     []TableRow
}

func NewAdminTable() *AdminTable {
	return &AdminTable{
		headings: []string{"one", "two", "three", "four", ""},
	}
}

func (table *AdminTable) AddRow(row TableRow) {
	table.rows = append(table.rows, row)
}
func (table *AdminTable) Render() string {
	var buf bytes.Buffer
	_ = AdminTbl(table.headings).Render(context.Background(), &buf)
	return string(buf.String())
	// return table.tmpl()
}
func (table *AdminTable) renderHeading() string {
	var cols []string
	for _, col := range table.headings {
		cols = append(
			cols,
			fmt.Sprintf(`<th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">%s</th>`, col),
		)
	}
	return strings.Join(cols, "\n")
}
func (table *AdminTable) renderDataRows() string {
	var trs []string
	for _, row := range table.rows {
		var tds []string
		for _, col := range table.headings {
			tds = append(
				tds,
				fmt.Sprintf(`<td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">%s</td> `, row[col]),
			)
		}
		trs = append(trs, fmt.Sprintf(`<tr>%s</tr>`, strings.Join(tds, "\n")))
	}
	return strings.Join(trs, "\n")
}

func (table *AdminTable) tmpl() string {
	return fmt.Sprintf(`
  <div class="px-4 sm:px-6 lg:px-8">
    <div class="sm:flex sm:items-center">
      <div class="sm:flex-auto">
        <h1 class="text-base font-semibold leading-6 text-gray-900">Users</h1>
        <p class="mt-2 text-sm text-gray-700">A list of all the users in your account including their name, title, email and role.</p>
      </div>
      <div class="mt-4 sm:ml-16 sm:mt-0 sm:flex-none">
        <button type="button" class="block rounded-md bg-indigo-600 px-3 py-2 text-center text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Add user</button>
      </div>
    </div>
    <div class="mt-8 flow-root">
      <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
          <table class="min-w-full divide-y divide-gray-300">
            <thead>
              <tr>
                %s
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr v-for="person in people" :key="person.email">
                <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">{{ person.name }}</td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ person.title }}</td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ person.email }}</td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ person.role }}</td>
                <td class="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-0">
                  <a href="#" class="text-indigo-600 hover:text-indigo-900"
                    >Edit<span class="sr-only">, {{ person.name }}</span></a
                  >
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
  `,
		table.renderHeading(),
	)
}
