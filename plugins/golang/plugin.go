package main;

import(
	WeldPlug "github.com/PoiXson/pxnWelder/weld/plugin"
	Work     "github.com/PoiXson/pxnWelder/weld/work"
);



const Version = "{{{VERSION}}}";



type PluginGoLang struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginGoLang{};
}



func (plugin *PluginGoLang) Run(workspace *Work.Workspace,
		stage string) error {
	print("Hello! GoLang works\n");
	return nil;
}
