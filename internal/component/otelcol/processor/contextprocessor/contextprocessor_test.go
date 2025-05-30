package contextprocessor

import (
	"testing"

	"github.com/grafana/alloy/syntax"
	"github.com/stretchr/testify/require"
)

func TestBadAlloyConfigKey(t *testing.T) {
	exampleBadAlloyConfigKey := `
		actions {
			key = ""
			action = "insert"
			from_attribute = "X-Scope-OrgId"
			value = "test"
		}
		output {
			// no-op: will be overridden by test code
		}
	`

	var args Arguments
	require.Error(t, syntax.Unmarshal([]byte(exampleBadAlloyConfigKey), &args), "key must not be empty")
}

func TestBadAlloyConfigValue(t *testing.T) {
	exampleBadAlloyConfigValue := `
		actions {
			key = "tenant"
			action = "insert"
			from_attribute = "x-scope-orgid"
			value = ""
		}		
		output {
			// no-op: will be overridden by test code
		}
	`

	var args Arguments
	require.Error(t, syntax.Unmarshal([]byte(exampleBadAlloyConfigValue), &args), "value must not be empty")
}

func TestBadAlloyConfigFromAttribute(t *testing.T) {
	exampleBadAlloyConfigFromAttribute := `
		actions {
			key = "tenant"
			action = ""
			from_attribute = "x-scope-orgid"
			value = "test"
		}
		output {
			// no-op: will be overridden by test code
		}
	`

	var args Arguments
	require.Error(t, syntax.Unmarshal([]byte(exampleBadAlloyConfigFromAttribute), &args), "action must not be empty")
}

func TestBadAlloyConfigAction(t *testing.T) {
	exampleBadAlloyConfigAction := `
		actions {
			key = "X-Scope-OrgId"
			action = "deleted"
			from_attribute = "x-scope-orgid"
			value = "test"
		}
		output {
			// no-op: will be overridden by test code
		}
	`

	var args Arguments
	require.Error(t, syntax.Unmarshal([]byte(exampleBadAlloyConfigAction), &args), "action")
}
