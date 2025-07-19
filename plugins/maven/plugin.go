package main;
// pxnWeld-Maven
// * config
// * build

import(
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



const Version = "{{{VERSION}}}";

const WEIGHT_MAVEN_CONFIG = 50;
const WEIGHT_MAVEN_BUILD  = 50;



type PluginMaven struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginMaven{};
}



func (plugin *PluginMaven) GetName() string {
	return "Maven";
}

func (plugin *PluginMaven) GetWeight(stage string) uint8 {
	switch (stage) {
		case "config": return WEIGHT_MAVEN_CONFIG;
		case "build":  return WEIGHT_MAVEN_BUILD;
	}
	return 0;
}



func (plugin *PluginMaven) Run(workspace *Workspace.Workspace, stage string) error {
	print("Hello! Maven works\n");
	return nil;
}
