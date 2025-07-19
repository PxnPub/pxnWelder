package main;
// pxnWeld-Git
// * vcs

import(
//	Exec      "os/exec"
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



const Version = "{{{VERSION}}}";

const WEIGHT_GIT_VCS = 50;



type PluginGit struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginGit{};
}



func (plugin *PluginGit) GetName() string {
	return "Git";
}

func (plugin *PluginGit) GetWeight(stage string) uint8 {
	if stage == "vcs" { return WEIGHT_GIT_VCS; }
	return 0;
}



func (plugin *PluginGit) Run(workspace *Workspace.Workspace, stage string) error {
	print("Hello! Git works\n");


//	cmd := Exec.Command("go", "build", ".");
//	if err := cmd.Run(); err != nil { return err; }


	return nil;
}
