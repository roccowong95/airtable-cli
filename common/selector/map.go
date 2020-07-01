package selector

import (
	"sort"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/olekukonko/tablewriter"
)

type item struct {
	key  string
	data map[string]interface{}
	line string
}

func getCompleter(items []item) prompt.Completer {
	var suggests []prompt.Suggest
	for _, item := range items {
		suggests = append(suggests, prompt.Suggest{
			// Text:        item.key,
			// Description: item.line,
			Text: item.line,
			// Description: item.key,
		})
	}
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterContains(suggests, d.GetWordBeforeCursor(), true)
	}
}

func SelectFromMap(data map[string]map[string]interface{}) string {
	var header []string
	for _, line := range data {
		if len(header) == 0 {
			for k := range line {
				header = append(header, k)
			}
			sort.Sort(strs(header))
			break
		}
	}
	return SelectFromMapWithHeader(data, header)
}

func SelectFromMapWithHeader(data map[string]map[string]interface{}, header []string) string {
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
	for key, line := range data {
		items = append(items, item{key: key, data: line})
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

	// table.SetHeader(header)
	table.AppendBulk(allLines)
	table.Render()

	lines := w.Dump()

	for i := 0; i < len(lines); i++ {
		idx := line2idx[strings.Join(allLines[i], "")]
		items[idx].line = lines[i]
		line2id[lines[i]] = items[idx].key
	}

	sort.Sort(itemArr(items))

	input := prompt.Input("> ", getCompleter(items))
	return line2id[input]
}
