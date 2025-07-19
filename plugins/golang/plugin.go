package main;
// pxnWeld-GoLang
// * build

import(
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



const Version = "{{{VERSION}}}";

const WEIGHT_BUILD = 50;



type PluginGoLang struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginGoLang{};
}



func (plugin *PluginGoLang) GetName() string {
	return "GoLang";
}

func (plugin *PluginGoLang) GetWeight(stage string) uint8 {
	if stage == "build" { return WEIGHT_BUILD; }
	return 0;
}



func (plugin *PluginGoLang) Run(workspace *Workspace.Workspace, stage string) error {
	print("Hello! GoLang works\n");
	return nil;
}
