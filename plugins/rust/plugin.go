package main;
// pxnWeld-Rust
// * config
// * build

import(
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



const Version = "{{{VERSION}}}";

const WEIGHT_RUST_CONFIG = 50;
const WEIGHT_RUST_BUILD  = 50;



type PluginRust struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginRust{};
}



func (plugin *PluginRust) GetName() string {
	return "Rust";
}

func (plugin *PluginRust) GetWeight(stage string) uint8 {
	switch (stage) {
		case "config": return WEIGHT_RUST_CONFIG;
		case "build":  return WEIGHT_RUST_BUILD;
	}
	return 0;
}



func (plugin *PluginRust) Run(workspace *Workspace.Workspace, stage string) error {
	print("Hello! Rust works\n");
	return nil;
}
