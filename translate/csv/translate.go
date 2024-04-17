package csv

import (
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/Miemiemiemieqiang/translator/core"
	"github.com/Miemiemiemieqiang/translator/translate"
)

type Translate struct {
}

func (t Translate) Config() core.Config {
	return core.NewConfig(core.WithType(translate.CSV))
}

func (t Translate) Predicate(bytes []byte) bool {
	return false
}

func (t Translate) Read(bs []byte) (*core.Core, error) {
	reader := csv.NewReader(bytes.NewReader(bs))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	results := make([]map[string]string, 0)
	header := records[0]
	for _, record := range records[1:] {
		result := make(map[string]string)
		for i, cell := range record {
			result[header[i]] = cell
		}
		results = append(results, result)
	}

	return core.NewCore(results), nil
}

func (t Translate) Write(c *core.Core) ([]byte, error) {
	if c.Recurse {
		return nil, errors.New("csv can not write recurse data")
	}
	data := c.Data
	if data == nil {
		return nil, errors.New("data is nil")
	}

	results, err := parseCsv(data)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	defer writer.Flush()
	err = writer.WriteAll(results)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func parseCsv(data interface{}) ([][]string, error) {
	switch i := data.(type) {
	case map[string]interface{}:
		return simpleObject(i), nil
	case []map[string]interface{}:
		return simpleArray(i), nil
	}
	return nil, errors.New("data type is not simple object")
}

func simpleObject(i interface{}) [][]string {
	o := i.(map[string]interface{})
	rs := make([][]string, 0)
	header := make([]string, 0)
	for k := range o {
		header = append(header, k)
	}
	rs = append(rs, header)
	record := make([]string, 0)
	for _, k := range header {
		record = append(record, o[k].(string))
	}
	rs = append(rs, record)
	return rs
}

func simpleArray(i interface{}) [][]string {
	os := i.([]map[string]interface{})
	rs := make([][]string, 0)
	header := make([]string, 0)
	for _, result := range os {
		for k := range result {
			header = append(header, k)
		}
		break
	}
	rs = append(rs, header)
	for _, result := range os {
		record := make([]string, 0)
		for _, k := range header {
			record = append(record, result[k].(string))
		}
		rs = append(rs, record)
	}
	return rs
}

func NewTranslate() *Translate {
	return &Translate{}
}

var (
	_ core.LangBid = Translate{}
)
