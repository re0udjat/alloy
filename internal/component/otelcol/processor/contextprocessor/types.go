package contextprocessor

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	ctxp "github.com/re0udjat/alloy-custom-components/contextprocessor"
)

type ActionType string

const (
	Insert ActionType = "insert"
	Upsert ActionType = "upsert"
	Update ActionType = "update"
	Delete ActionType = "delete"
)

type ActionConfig struct {
	Key           string     `alloy:"key,attr"`
	Action        ActionType `alloy:"action,attr"`
	Value         string     `alloy:"value,attr"`
	FromAttribute string     `alloy:"from_attribute,attr"`
}

func (actionType *ActionType) UnmarshalText(text []byte) error {
	str := ActionType(strings.ToLower(string(text)))
	switch str {
	case Insert, Upsert, Update, Delete:
		*actionType = str
		return nil
	default:
		return fmt.Errorf("unknown action type %v", str)
	}
}

func (actionConfig ActionConfig) Convert() (*ctxp.ActionConfig, error) {
	var otelActionConfig ctxp.ActionConfig

	err := mapstructure.Decode(map[string]interface{}{
		"key":            actionConfig.Key,
		"value":          actionConfig.Value,
		"action":         actionConfig.Action,
		"from_attribute": actionConfig.FromAttribute,
	}, &otelActionConfig)

	if err != nil {
		return nil, err
	}
	return &otelActionConfig, nil
}
