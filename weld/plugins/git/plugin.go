package main;

import(
//	Exec     "os/exec"
	WeldPlug "github.com/PoiXson/pxnWelder/weld/plugin"
);



const Version = "{{{VERSION}}}";



type PluginGit struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginGit{};
}



func (plugin *PluginGit) Run(stage WeldPlug.WeldStage) error {
	print("Hello! Git works\n");


//	cmd := Exec.Command("go", "build", ".");
//	if err := cmd.Run(); err != nil { return err; }


	return nil;
}
