package view

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/xagero/go-helper/helper"
)

const (
	TextAlignLeft  = "text_align_left"
	TextAlignRight = "text_align_right"
)

type columnConfig struct {
	title     string // @todo title in header
	padding   int    // @todo column padding
	textAlign string
}

type Table interface {
	SetHeading(heading string) Table
	AddRow(values ...interface{}) Table
	Render()
}

type table struct {

	// Config
	padding int
	column  map[int]*columnConfig

	// Data
	heading string
	header  []string
	rows    [][]string

	// Processed
	columnWidth []int
	totalWidth  int
}

func Construct(headers ...interface{}) Table {
	// Create new table
	table := new(table)

	// Configure
	table.heading = ""
	table.padding = 4
	table.header = make([]string, len(headers))
	table.column = make(map[int]*columnConfig)

	// Setup header column
	for i, col := range headers {
		// @todo remove me
		table.header[i] = fmt.Sprint(col)

		cc := new(columnConfig)
		cc.title = fmt.Sprint(col)
		cc.textAlign = TextAlignRight

		table.column[i] = cc
	}

	return table
}

// SetHeading Set table heading
func (table *table) SetHeading(heading string) Table {
	table.heading = heading
	return table
}

// AddRow insert new row into Table
func (table *table) AddRow(values ...interface{}) Table {

	counter := 0 // column
	for _, val := range values {
		counter = helper.Max(strings.Count(fmt.Sprint(val), "\n"), counter)
	}

	for i := 0; i <= counter; i++ {
		row := make([]string, len(table.header))
		for j, val := range values {
			if j >= len(table.header) {
				break
			}
			v := strings.Split(fmt.Sprint(val), "\n")
			row[j] = offset(v, i)
		}
		table.rows = append(table.rows, row)
	}

	return table
}

// Render print the Table
func (table *table) Render() {
	format := strings.Repeat("%s", len(table.header)) + "\n"

	table.process()

	table.printHeading()
	table.printTableBetweenRowLine()
	table.printHeader(format)
	table.printTableBetweenRowLine()

	for _, row := range table.rows {
		table.printRow(format, row)
	}

	table.printTableBetweenRowLine()
}

func (table *table) printTableBetweenRowLine() {
	fmt.Println(strings.Repeat("-", table.totalWidth))
}

func (table *table) printHeading() {
	if helper.IsNotBlank(table.heading) {
		fmt.Println("  " + table.heading)
	}
}

func (table *table) printHeader(format string) {
	values := table.applyWidth(table.header, table.columnWidth)
	fmt.Printf(format, values...)
}

func (table *table) printRow(format string, row []string) {
	values := table.applyWidth(row, table.columnWidth)
	fmt.Printf(format, values...)
}

func (table *table) applyWidth(row []string, widths []int) []interface{} {
	out := make([]interface{}, len(row))
	for i, s := range row {
		if table.column[i].textAlign == TextAlignLeft {
			out[i] = s + table.length(s, widths[i])
		} else if table.column[i].textAlign == TextAlignRight {
			out[i] = table.length(s, widths[i]) + s
		}
	}
	return out
}

func (table *table) length(s string, w int) string {
	count := w - utf8.RuneCountInString(s)
	if count <= 0 {
		return ""
	}
	return strings.Repeat(" ", count)
}

func offset(arr []string, idx int) string {
	if idx >= len(arr) {
		return ""
	}
	return arr[idx]
}

func (table *table) process() {

	table.totalWidth = 0
	table.columnWidth = make([]int, len(table.header))

	for _, row := range table.rows {
		for i, v := range row {
			w := utf8.RuneCountInString(v) + table.padding
			if w > table.columnWidth[i] {
				table.columnWidth[i] = w
			}
		}
	}

	for i, v := range table.header {
		w := utf8.RuneCountInString(v) + table.padding
		if w > table.columnWidth[i] {
			table.columnWidth[i] = w
		}
	}

	for _, w := range table.columnWidth {
		table.totalWidth += w
	}
}
