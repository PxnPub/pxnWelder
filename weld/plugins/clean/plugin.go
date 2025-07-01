package main;

import(
	WeldPlug "github.com/PoiXson/pxnWelder/weld/plugin"
);



const Version = "0.1.1";



type PluginClean struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginClean{};
}



func (plugin *PluginClean) Run(stage WeldPlug.WeldStage) error {
	print("Hello! Clean works\n");
	return nil;
}
