// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package orbattributesprocessor

import (
	"go.opentelemetry.io/collector/config"

	"github.com/ns1labs/orb/otelcollector/components/internal/attraction"
	"github.com/ns1labs/orb/otelcollector/components/internal/filterconfig"
)

// Config specifies the set of attributes to be inserted, updated, upserted and
// deleted and the properties to include/exclude a span from being processed.
// This processor handles all forms of modifications to attributes within a span, log, or metric.
// Prior to any actions being applied, each span is compared against
// the include properties and then the exclude properties if they are specified.
// This determines if a span is to be processed or not.
// The list of actions is applied in order specified in the configuration.
type Config struct {
	config.ProcessorSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct

	filterconfig.MatchConfig `mapstructure:",squash"`

	// Specifies the list of attributes to act on.
	// The set of actions are {INSERT, UPDATE, UPSERT, DELETE, HASH, EXTRACT}.
	// This is a required field.
	attraction.Settings `mapstructure:",squash"`
}

var _ config.Processor = (*Config)(nil)

// Validate checks if the processor configuration is valid
func (cfg *Config) Validate() error {
	return nil
}

func (cfg *Config) AddAttribute(name string, value interface{}) (ok bool) {
	for _, action := range cfg.Settings.Actions {
		if action.Key == name {
			return false
		}
	}
	cfg.Settings.Actions = append(cfg.Settings.Actions, attraction.ActionKeyValue{
		Key:    name,
		Value:  value,
		Action: "insert",
	})
	return ok
}
