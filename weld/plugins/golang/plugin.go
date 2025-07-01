package main;

import(
	WeldPlug "github.com/PoiXson/pxnWelder/weld/plugin"
);



const Version = "{{{VERSION}}}";



type PluginGoLang struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginGoLang{};
}



func (plugin *PluginGoLang) Run(stage WeldPlug.WeldStage) error {
	print("Hello! GoLang works\n");
	return nil;
}
