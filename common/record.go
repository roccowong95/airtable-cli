package common

type Record struct {
	AirtableID string                 `json:"id,omitempty"`
	Fields     map[string]interface{} `json:"fields"`
}
