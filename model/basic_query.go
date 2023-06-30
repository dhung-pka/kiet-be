package model

var TypeModel = map[string]interface{}{
	"profile":    Profile{},
	"credential": Credential{},
	"district":   District{},
	"province":   Province{},
}

var TypeListModel = map[string]interface{}{
	"[]profile":    []Profile{},
	"[]credential": []Credential{},
	"[]district":   []District{},
	"[]province":   []Province{},
}

type BasicQueryRequest struct {
	Data      interface{}    `json:"data"`
	ModelType string         `json:"modelType"`
	Option    OPTION_REQUEST `json:"option"`
}

type OPTION_REQUEST string

const (
	INSERT OPTION_REQUEST = "INSERT"
	UPDATE OPTION_REQUEST = "UPDATE"
	DELETE OPTION_REQUEST = "DELETE"
)

var OptionRequest = map[OPTION_REQUEST]interface{}{
	INSERT: INSERT,
	UPDATE: UPDATE,
	DELETE: DELETE,
}
