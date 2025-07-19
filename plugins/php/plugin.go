package main;
// pxnWeld-PHP
// * config

import(
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



const Version = "{{{VERSION}}}";

const WEIGHT_PHP_CONFIG = 50;



type PluginPHP struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginPHP{};
}



func (plugin *PluginPHP) GetName() string {
	return "PHP";
}

func (plugin *PluginPHP) GetWeight(stage string) uint8 {
	if stage == "config" { return WEIGHT_PHP_CONFIG; }
	return 0;
}



func (plugin *PluginPHP) Run(workspace *Workspace.Workspace, stage string) error {
	print("Hello! PHP works\n");
	return nil;
}
