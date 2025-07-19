package main;
// pxnWeld-Make
// * config
// * build

import(
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



const Version = "{{{VERSION}}}";

const WEIGHT_MAKE_CONFIG = 50;
const WEIGHT_MAKE_BUILD  = 50;



type PluginMake struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginMake{};
}



func (plugin *PluginMake) GetName() string {
	return "Make";
}

func (plugin *PluginMake) GetWeight(stage string) uint8 {
	switch (stage) {
		case "config": return WEIGHT_MAKE_CONFIG;
		case "build":  return WEIGHT_MAKE_BUILD;
	}
	return 0;
}



func (plugin *PluginMake) Run(workspace *Workspace.Workspace, stage string) error {
	print("Hello! Make works\n");
	return nil;
}
