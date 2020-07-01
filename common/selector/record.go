package selector

import (
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/olekukonko/tablewriter"
	"github.com/roccowong95/airtable-cli/common"
)

// SelectFromRecords 不对选项进行排序.
func SelectFromRecords(recs []common.Record, header []string) string {
	w := newMultilineWriter()
	table := tablewriter.NewWriter(w)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding(" ")
	table.SetNoWhiteSpace(true)

	var items []item
	for _, rec := range recs {
		items = append(items, item{key: rec.AirtableID, data: rec.Fields})
	}

	var allLines [][]string
	line2idx := make(map[string]int)
	line2id := make(map[string]string)
	for idx, item := range items {
		var line []string
		for _, k := range header {
			v, ok := item.data[k]
			if !ok {
				line = append(line, "")
				continue
			}
			if f, ok := v.(float64); ok {
				line = append(line, strconv.FormatFloat(f, 'e', 2, 64))
			} else if i, ok := v.(int); ok {
				line = append(line, strconv.Itoa(i))
			} else if s, ok := v.(string); ok {
				line = append(line, s)
			} else {
				line = append(line, "obj")
			}
		}
		allLines = append(allLines, line)
		k := strings.Join(line, "")
		line2idx[k] = idx
	}

	table.AppendBulk(allLines)
	table.Render()

	lines := w.Dump()

	for i := 0; i < len(lines); i++ {
		idx := line2idx[strings.Join(allLines[i], "")]
		items[idx].line = lines[i]
		line2id[lines[i]] = items[idx].key
	}

	input := prompt.Input("> ", getCompleter(items))
	return line2id[input]
}
