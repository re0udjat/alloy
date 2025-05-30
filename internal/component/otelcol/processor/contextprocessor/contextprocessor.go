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

// Arguments configures the otelcol.processor.contextprocessor component
type Arguments struct {
	ActionCfgs []ActionConfig `alloy:"actions,block"`

	// Output configures where to send processed data. Required.
	Output *otelcol.ConsumerArguments `alloy:"output,block"`

	// DebugMetrics configures component internal metrics. Optional.
	DebugMetrics otelcolCfg.DebugMetricsArguments `alloy:"debug_metrics,block,optional"`
}

var (
	_ processor.Arguments = Arguments{}
)

func (args *Arguments) SetToDefault() {
	args.DebugMetrics.SetToDefault()
}

func init() {
	component.Register(component.Registration{
		Name:      "otelcol.processor.contextprocessor",
		Stability: featuregate.StabilityGenerallyAvailable,
		Args:      Arguments{},
		Exports:   otelcol.ConsumerExports{},

		Build: func(opts component.Options, args component.Arguments) (component.Component, error) {
			fact := ctxp.NewFactory()
			return processor.New(opts, fact, args.(Arguments))
		},
	})
}

// Validate implements syntax.Validator
func (args *Arguments) Validate() error {
	for _, actionCfg := range args.ActionCfgs {
		if actionCfg.Key == "" {
			return fmt.Errorf("key must not be empty")
		}

		if actionCfg.Value == "" {
			return fmt.Errorf("value must not be empty")
		}

		if actionCfg.FromAttribute == "" {
			return fmt.Errorf("from_attribute must not be empty")
		}

		if actionCfg.Action != Insert && actionCfg.Action != Upsert && actionCfg.Action != Update && actionCfg.Action != Delete {
			return fmt.Errorf("action must be one of the following values: insert, upsert, update or delete")
		}
	}

	return nil
}

// Convert implements processor.Arguments
func (args Arguments) Convert() (otelcomponent.Config, error) {
	var otelActionCfgs []ctxp.ActionConfig
	for _, actionCfg := range args.ActionCfgs {
		otelActionCfg, err := actionCfg.Convert()
		if err == nil {
			otelActionCfgs = append(otelActionCfgs, *otelActionCfg)
		}
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
