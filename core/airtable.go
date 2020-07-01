package core

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/fabioberger/airtable-go"
	"github.com/roccowong95/airtable-cli/common"
	"github.com/roccowong95/airtable-cli/common/conf"
	"github.com/roccowong95/airtable-cli/common/selector"
)

type airtableCore struct {
	conf conf.AppConf

	cli   *airtable.Client
	table string
}

func NewAirtableCore(c conf.AppConf) (CliCore, error) {
	var ret airtableCore
	cli, err := airtable.New(c.APIKey, c.BaseID)
	if nil != err {
		return nil, err
	}

	ret.conf = c
	ret.cli = cli
	ret.table = c.Table
	return &ret, nil
}

func (c *airtableCore) Add(args []string) error {
	if len(args)%2 != 0 {
		return fmt.Errorf("len of args is odd(%d)", len(args))
	}
	fields := make(map[string]struct{})
	var rec common.Record
	rec.Fields = make(map[string]interface{})
	for i := 0; i < len(args); i += 2 {
		alias := args[i]
		fields[alias] = struct{}{}
		val := args[i+1]
		field, ok := c.conf.Fields[alias]
		if !ok {
			return fmt.Errorf("alias %s is not found in conf", alias)
		}
		if field.Type == conf.FieldNum {
			f, err := strconv.ParseFloat(val, 64)
			if nil != err {
				return fmt.Errorf("parse float for field %s failed: %v", alias, err)
			}
			if field.Negative {
				f = -f
			}
			rec.Fields[field.Name] = f
			continue
		}
		rec.Fields[field.Name] = val
	}

	var allAlias []string
	for alias := range c.conf.Fields {
		allAlias = append(allAlias, alias)
	}
	sort.Strings(allAlias)
	for _, alias := range allAlias {
		field := c.conf.Fields[alias]
		_, ok := fields[alias]

		// 是必须字段, 且不是link, 没有提供
		if field.Must && !ok && field.Type != conf.FieldLink {
			return fmt.Errorf("field '%s' is a must but not provided", field.Name)
		}

		if !(!ok && field.Type == conf.FieldLink) {
			continue
		}
		if len(field.LinkTable) == 0 {
			return fmt.Errorf("field %s is FieldLink, but LinkTable is not provided", alias)
		}

		var rc []common.Record
		err := c.cli.ListRecords(field.LinkTable, &rc, airtable.ListParameters{View: field.LinkView})
		if nil != err {
			return fmt.Errorf("list field in table %s failed: %v", field.LinkTable, err)
		}

		m := make(map[string]map[string]interface{})
		for _, r := range rc {
			m[r.AirtableID] = r.Fields
		}
		fmt.Printf("Please select %s:\n", field.Name)
		// rec.Fields[field.Name] = []string{selector.SelectFromMapWithHeader(m, field.LinkHeader)}
		rec.Fields[field.Name] = []string{selector.SelectFromRecords(rc, field.LinkHeader)}
	}
	return c.cli.CreateRecord(c.table, rec)
}

func (c *airtableCore) List() error {
	// c.cli.ListRecords()
	return nil
}
