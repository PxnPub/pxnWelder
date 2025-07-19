package main;
// pxnWeld-Deploy
// * deploy

import(
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



const Version = "{{{VERSION}}}";

const WEIGHT_DEPLOY = 50;



type PluginDeploy struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginDeploy{};
}



func (plugin *PluginDeploy) GetName() string {
	return "Deploy";
}

func (plugin *PluginDeploy) GetWeight(stage string) uint8 {
	if stage == "deploy" { return WEIGHT_DEPLOY; }
	return 0;
}



func (plugin *PluginDeploy) Run(workspace *Workspace.Workspace, stage string) error {
	print("Hello! Deploy works\n");
	return nil;
}
