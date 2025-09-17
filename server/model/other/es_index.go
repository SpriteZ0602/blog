package other

import "encoding/json"

type Data struct {
	ID  *string
	Doc json.RawMessage
}

type ESIndexResponse struct {
	Data []Data `json:"data"`
}
