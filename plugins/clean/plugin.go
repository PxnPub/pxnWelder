package main;

import(
	WeldPlug "github.com/PoiXson/pxnWelder/weld/plugin"
	Work     "github.com/PoiXson/pxnWelder/weld/work"
);



const Version = "{{{VERSION}}}";



type PluginClean struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginClean{};
}



func (plugin *PluginClean) Run(workspace *Work.Workspace,
		stage string) error {
	print("Hello! Clean works\n");
	return nil;
}
