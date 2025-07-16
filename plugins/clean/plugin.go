package main;

import(
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



const Version = "{{{VERSION}}}";



type PluginClean struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginClean{};
}



func (plugin *PluginClean) Run(workspace *Workspace.Workspace,
		stage string) error {
	print("Hello! Clean works\n");
	return nil;
}
