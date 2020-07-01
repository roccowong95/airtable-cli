package conf

import "fmt"

type Conf struct {
	APIKey   string
	BaseID   string
	AppConfs []AppConf
}

type AppConf struct {
	Name   string
	APIKey string
	BaseID string
	Table  string
	View   string
	// Fields is the mapping of fields, key is alias.
	Fields map[string]Field
}

type FieldType string

const (
	FieldStr  FieldType = "str"
	FieldNum  FieldType = "num"
	FieldLink FieldType = "link"
)

type Field struct {
	Name       string
	Type       FieldType
	LinkTable  string
	Must       bool
	Negative   bool
	LinkHeader []string
	LinkView   string
}

func (c *Conf) Fixup() error {
	for i := range c.AppConfs {
		if len(c.AppConfs[i].APIKey) == 0 {
			c.AppConfs[i].APIKey = c.APIKey
		}
		if len(c.AppConfs[i].APIKey) == 0 {
			return fmt.Errorf("empty apikey for appconfs[%d]", i)
		}

		if len(c.AppConfs[i].BaseID) == 0 {
			return fmt.Errorf("empty baseid for appconfs[%d]", i)
		}
		if len(c.AppConfs[i].BaseID) == 0 {
			c.AppConfs[i].BaseID = c.BaseID
		}

		if len(c.AppConfs[i].Name) == 0 {
			return fmt.Errorf("empty name for appconfs[%d]", i)
		}
		if len(c.AppConfs[i].Table) == 0 {
			return fmt.Errorf("empty table for appconfs[%d]", i)
		}
		if len(c.AppConfs[i].View) == 0 {
			return fmt.Errorf("empty view for appconfs[%d]", i)
		}
	}
	return nil
}

func (c *Conf) GetApp(name string) *AppConf {
	for _, conf := range c.AppConfs {
		if conf.Name == name {
			return &conf
		}
	}
	return nil
}
