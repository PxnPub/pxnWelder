package main;

import(
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



const Version = "{{{VERSION}}}";



type PluginGoLang struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginGoLang{};
}



func (plugin *PluginGoLang) Run(workspace *Workspace.Workspace,
		stage string) error {
	print("Hello! GoLang works\n");
	return nil;
}
