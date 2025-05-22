package contextprocessor

import (
	"github.com/mitchellh/mapstructure"
	ctxp "github.com/re0udjat/alloy-custom-components/contextprocessor"
)

func (actionCfg ActionConfig) Convert() (*ctxp.ActionConfig, error) {
	var otelActionCfg ctxp.ActionConfig

	err := mapstructure.Decode(map[string]interface{}{
		"key":            actionCfg.Key,
		"value":          actionCfg.Value,
		"action":         actionCfg.Action,
		"from_attribute": actionCfg.FromAttr,
	}, &otelActionCfg)

	if err != nil {
		return nil, err
	}

	return &otelActionCfg, nil
}
