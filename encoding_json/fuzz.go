package encoding_json

import (
    "encoding/json"
)

func Fuzz(input []byte) {
    var objmap map[string]*json.RawMessage
    json.Unmarshal(input, &objmap)
}
