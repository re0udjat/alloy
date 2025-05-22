package contextprocessor

import (
	"fmt"

	"github.com/grafana/alloy/internal/component"
	"github.com/grafana/alloy/internal/component/otelcol"
	otelcolCfg "github.com/grafana/alloy/internal/component/otelcol/config"
	"github.com/grafana/alloy/internal/component/otelcol/processor"
	"github.com/grafana/alloy/internal/featuregate"
	ctxp "github.com/re0udjat/alloy-custom-components/contextprocessor"
	otelcomponent "go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pipeline"
)

const (
	INSERT ActionType = "insert"
	UPDATE ActionType = "update"
	UPSERT ActionType = "upsert"
	DELETE ActionType = "delete"
)

type ActionType string

type ActionConfig struct {
	Key      string     `alloy:"key,attr"`
	Value    string     `alloy:"value,attr"`
	Action   ActionType `alloy:"action,attr"`
	FromAttr string     `alloy:"from_attribute,attr"`
}

type Arguments struct {
	ActionCfgs []ActionConfig `alloy:"actions,block"`

	// Output configures where to send processed data. Required.
	Output *otelcol.ConsumerArguments `alloy:"output,block"`

	// DebugMetrics configures component internal metrics. Optional.
	DebugMetrics otelcolCfg.DebugMetricsArguments `alloy:"debug_metrics,block,optional"`
}

func init() {
	component.Register(component.Registration{
		Name:      "otelcol.processor.contextprocessor",
		Stability: featuregate.StabilityPublicPreview,
		Args:      Arguments{},
		Exports:   otelcol.ConsumerExports{},
		Build: func(opts component.Options, args component.Arguments) (component.Component, error) {
			fact := ctxp.NewFactory()
			return processor.New(opts, fact, args.(Arguments))
		},
	})
}

var (
	_ processor.Arguments = Arguments{}
)

// SetToDefault implements syntax.Defaulter.
func (args *Arguments) SetToDefault() {
	args.DebugMetrics.SetToDefault()
}

// Validate implements syntax.Validator.
func (args *Arguments) Validate() error {
	for _, actionCfg := range args.ActionCfgs {
		if len(actionCfg.Key) == 0 {
			return fmt.Errorf("key must be not empty")
		}
		if len(actionCfg.Value) == 0 {
			return fmt.Errorf("value must be not empty")
		}
		if len(actionCfg.FromAttr) == 0 {
			return fmt.Errorf("from_attribute must be not empty")
		}
		if actionCfg.Action != INSERT && actionCfg.Action != UPDATE && actionCfg.Action != UPSERT && actionCfg.Action != DELETE {
			return fmt.Errorf("action must be one of the following: insert, update, upsert or delete")
		}
	}

	return nil
}

// Convert implements processor.Arguments.
func (args Arguments) Convert() (otelcomponent.Config, error) {
	var otelActionCfgs []ctxp.ActionConfig
	for _, actionCfg := range args.ActionCfgs {
		newActionCfg, err := actionCfg.Convert()
		if err != nil {
			return nil, err
		}
		otelActionCfgs = append(otelActionCfgs, *newActionCfg)
	}

	return &ctxp.Config{
		ActionsConfig: otelActionCfgs,
	}, nil
}

// Extensions implements processor.Arguments.
func (args Arguments) Extensions() map[otelcomponent.ID]otelcomponent.Component {
	return nil
}

// Exporters implements processor.Arguments.
func (args Arguments) Exporters() map[pipeline.Signal]map[otelcomponent.ID]otelcomponent.Component {
	return nil
}

// NextConsumers implements processor.Arguments.
func (args Arguments) NextConsumers() *otelcol.ConsumerArguments {
	return args.Output
}

// DebugMetricsConfig implements processor.Arguments.
func (args Arguments) DebugMetricsConfig() otelcolCfg.DebugMetricsArguments {
	return args.DebugMetrics
}
