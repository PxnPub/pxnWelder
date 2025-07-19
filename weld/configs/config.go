package configs;
// pxnWelder tool config

import(
	_ "gopkg.in/yaml.v2"
);



type CfgWeldTool struct {
	PathPlugins   string `yaml:"Path-Plugins"`
	PathTemplates string `yaml:"Path-Templates"`
}
