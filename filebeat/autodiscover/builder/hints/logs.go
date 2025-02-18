// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package hints

import (
	"fmt"
	"regexp"

	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/elastic/go-ucfg"

	"github.com/k0ffee/beats/v7/filebeat/fileset"
	"github.com/k0ffee/beats/v7/filebeat/harvester"
	"github.com/k0ffee/beats/v7/libbeat/autodiscover"
	"github.com/k0ffee/beats/v7/libbeat/autodiscover/builder"
	"github.com/k0ffee/beats/v7/libbeat/autodiscover/template"
	"github.com/k0ffee/beats/v7/libbeat/beat"
	"github.com/k0ffee/beats/v7/libbeat/common/bus"
	"github.com/k0ffee/beats/v7/libbeat/logp"
)

func init() {
	autodiscover.Registry.AddBuilder("hints", NewLogHints)
}

const (
	multiline    = "multiline"
	includeLines = "include_lines"
	excludeLines = "exclude_lines"
	processors   = "processors"
	json         = "json"
	pipeline     = "pipeline"
)

// validModuleNames to sanitize user input
var validModuleNames = regexp.MustCompile("[^a-zA-Z0-9\\_\\-]+")

type logHints struct {
	config   *config
	registry *fileset.ModuleRegistry
	log      *logp.Logger
}

// NewLogHints builds a log hints builder
func NewLogHints(cfg *conf.C) (autodiscover.Builder, error) {
	config := defaultConfig()
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("unable to unpack hints config due to error: %w", err)
	}

	moduleRegistry, err := fileset.NewModuleRegistry(nil, beat.Info{}, false)
	if err != nil {
		return nil, err
	}

	return &logHints{&config, moduleRegistry, logp.NewLogger("hints.builder")}, nil
}

// Create config based on input hints in the bus event
func (l *logHints) CreateConfig(event bus.Event, options ...ucfg.Option) []*conf.C {
	var hints mapstr.M
	if hintsIfc, found := event["hints"]; found {
		hints, _ = hintsIfc.(mapstr.M)
	}

	// Hint must be explicitly enabled when default_config sets enabled=false.
	if !l.config.DefaultConfig.Enabled() && !builder.IsEnabled(hints, l.config.Key) ||
		builder.IsDisabled(hints, l.config.Key) {
		l.log.Debugw("Hints config is not enabled.", "autodiscover.event", event)
		return nil
	}

	if inputConfig := l.getInputsConfigs(hints); inputConfig != nil {
		var configs []*conf.C
		for _, cfg := range inputConfig {
			if config, err := conf.NewConfigFrom(cfg); err == nil {
				configs = append(configs, config)
			} else {
				l.log.Warnw("Failed to create config from input.", "error", err)
			}
		}
		l.log.Debugf("Generated %d input configs from hint.", len(configs))
		// Apply information in event to the template to generate the final config
		return template.ApplyConfigTemplate(event, configs)
	}

	var configs []*conf.C
	inputs := l.getInputs(hints)
	for _, h := range inputs {
		// Clone original config, enable it if disabled
		config, _ := conf.NewConfigFrom(l.config.DefaultConfig)
		config.Remove("enabled", -1)

		tempCfg := mapstr.M{}
		mline := l.getMultiline(h)
		if len(mline) != 0 {
			tempCfg.Put(multiline, mline)
		}
		if ilines := l.getIncludeLines(h); len(ilines) != 0 {
			tempCfg.Put(includeLines, ilines)
		}
		if elines := l.getExcludeLines(h); len(elines) != 0 {
			tempCfg.Put(excludeLines, elines)
		}

		if procs := l.getProcessors(h); len(procs) != 0 {
			tempCfg.Put(processors, procs)
		}

		if pip := l.getPipeline(h); len(pip) != 0 {
			tempCfg.Put(pipeline, pip)
		}

		if jsonOpts := l.getJSONOptions(h); len(jsonOpts) != 0 {
			tempCfg.Put(json, jsonOpts)
		}
		// Merge config template with the configs from the annotations
		if err := config.Merge(tempCfg); err != nil {
			logp.Debug("hints.builder", "config merge failed with error: %v", err)
			continue
		}

		module := l.getModule(hints)
		if module != "" {
			moduleConf := map[string]interface{}{
				"module": module,
			}

			filesets := l.getFilesets(hints, module)
			for fileset, cfg := range filesets {
				filesetConf, _ := conf.NewConfigFrom(config)

				if inputType, _ := filesetConf.String("type", -1); inputType == harvester.ContainerType {
					filesetConf.SetString("stream", -1, cfg.Stream)
				} else {
					filesetConf.SetString("containers.stream", -1, cfg.Stream)
				}

				moduleConf[fileset+".enabled"] = cfg.Enabled
				moduleConf[fileset+".input"] = filesetConf

				logp.Debug("hints.builder", "generated config %+v", moduleConf)
			}
			config, _ = conf.NewConfigFrom(moduleConf)
		}
		logp.Debug("hints.builder", "generated config %+v", config)
		configs = append(configs, config)
	}

	// Apply information in event to the template to generate the final config
	return template.ApplyConfigTemplate(event, configs)
}

func (l *logHints) getMultiline(hints mapstr.M) mapstr.M {
	return builder.GetHintMapStr(hints, l.config.Key, multiline)
}

func (l *logHints) getIncludeLines(hints mapstr.M) []string {
	return builder.GetHintAsList(hints, l.config.Key, includeLines)
}

func (l *logHints) getExcludeLines(hints mapstr.M) []string {
	return builder.GetHintAsList(hints, l.config.Key, excludeLines)
}

func (l *logHints) getModule(hints mapstr.M) string {
	module := builder.GetHintString(hints, l.config.Key, "module")
	// for security, strip module name
	return validModuleNames.ReplaceAllString(module, "")
}

func (l *logHints) getInputsConfigs(hints mapstr.M) []mapstr.M {
	return builder.GetHintAsConfigs(hints, l.config.Key)
}

func (l *logHints) getProcessors(hints mapstr.M) []mapstr.M {
	return builder.GetProcessors(hints, l.config.Key)
}

func (l *logHints) getPipeline(hints mapstr.M) string {
	return builder.GetHintString(hints, l.config.Key, "pipeline")
}

func (l *logHints) getJSONOptions(hints mapstr.M) mapstr.M {
	return builder.GetHintMapStr(hints, l.config.Key, json)
}

type filesetConfig struct {
	Enabled bool
	Stream  string
}

// Return a map containing filesets -> enabled & stream (stdout, stderr, all)
func (l *logHints) getFilesets(hints mapstr.M, module string) map[string]*filesetConfig {
	var configured bool
	filesets := make(map[string]*filesetConfig)

	moduleFilesets, err := l.registry.ModuleAvailableFilesets(module)
	if err != nil {
		logp.Err("Error retrieving module filesets: %+v", err)
		return nil
	}

	for _, fileset := range moduleFilesets {
		filesets[fileset] = &filesetConfig{Enabled: false, Stream: "all"}
	}

	// If a single fileset is given, pass all streams to it
	fileset := builder.GetHintString(hints, l.config.Key, "fileset")
	if fileset != "" {
		if conf, ok := filesets[fileset]; ok {
			conf.Enabled = true
			configured = true
		}
	}

	// If fileset is defined per stream, return all of them
	for _, stream := range []string{"all", "stdout", "stderr"} {
		fileset := builder.GetHintString(hints, l.config.Key, "fileset."+stream)
		if fileset != "" {
			if conf, ok := filesets[fileset]; ok {
				conf.Enabled = true
				conf.Stream = stream
				configured = true
			}
		}
	}

	// No fileset defined, return defaults for the module, all streams to all filesets
	if !configured {
		for _, conf := range filesets {
			conf.Enabled = true
		}
	}

	return filesets
}

func (l *logHints) getInputs(hints mapstr.M) []mapstr.M {
	modules := builder.GetHintsAsList(hints, l.config.Key)
	var output []mapstr.M

	for _, mod := range modules {
		output = append(output, mapstr.M{
			l.config.Key: mod,
		})
	}

	// Generate this so that no hints with completely valid templates work
	if len(output) == 0 {
		output = append(output, mapstr.M{
			l.config.Key: mapstr.M{},
		})
	}

	return output
}
