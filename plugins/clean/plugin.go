package main;
// pxnWeld-Clean
// * clean

import(
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



const Version = "{{{VERSION}}}";

const WEIGHT_CLEAN = 50;



type PluginClean struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginClean{};
}



func (plugin *PluginClean) GetName() string {
	return "Clean";
}

func (plugin *PluginClean) GetWeight(stage string) uint8 {
	if stage == "clean" { return WEIGHT_CLEAN; }
	return 0;
}



func (plugin *PluginClean) Run(workspace *Workspace.Workspace, stage string) error {
	print("Hello! Clean works\n");
	return nil;
}
