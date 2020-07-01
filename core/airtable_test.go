package core_test

import (
	"testing"

	"github.com/roccowong95/airtable-cli/common/conf"
	"github.com/roccowong95/airtable-cli/core"
)

func TestCreate(t *testing.T) {
	c, err := core.NewAirtableCore(conf.AppConf{
		Table:  "Records",
		APIKey: "keyNRcjX5yDeDcBXU",
		BaseID: "applajgx3arovUxz4",
	})
	if nil != err {
		t.Fatal(err)
	}
	err = c.Add(map[string]interface{}{"Amount": 113})
	if nil != err {
		t.Fatal(err)
	}
}
