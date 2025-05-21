package contextprocessor

import (
	"github.com/grafana/alloy/internal/component"
	"github.com/grafana/alloy/internal/component/otelcol"
	"github.com/grafana/alloy/internal/component/otelcol/processor"
	"github.com/grafana/alloy/internal/featuregate"
	"github.com/re0udjat/alloy-custom-components/contextprocessor"
)

type Arguments struct {
}

func init() {
	component.Register(component.Registration{
		Name:      "otelcol.processor.contextprocessor",
		Stability: featuregate.StabilityPublicPreview,
		Args:      Arguments{},
		Exports:   otelcol.ConsumerExports{},

		Build: func(opts component.Options, args component.Arguments) (component.Component, error) {
			return processor.New(opts, contextprocessor.NewFactory(), args.(Arguments))
		},
	})
}
