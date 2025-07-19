package main;
// pxnWeld-RPM
// * package

import(
	OS        "os"
	Fmt       "fmt"
	Time      "time"
	IOUtils   "io/ioutil"
	RPM       "github.com/google/rpmpack"
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



const Version = "{{{VERSION}}}";

const WEIGHT_PACKAGE = 50;



type PluginRPM struct {
}



func NewPlugin() WeldPlug.WeldPlugin {
	return &PluginRPM{};
}



func (plugin *PluginRPM) GetName() string {
	return "RPM";
}

func (plugin *PluginRPM) GetWeight(stage string) uint8 {
	if stage == "package" { return WEIGHT_PACKAGE; }
	return 0;
}



func (plugin *PluginRPM) Run(workspace *Workspace.Workspace, stage string) error {
	print("Hello! RPM works\n");
	rpm, err := RPM.NewRPM(RPM.RPMMetaData{
		Name:    "example",
		Version: "1.0.0",
		Compressor: "zstd:1",
		Summary: "asfhgjh",
		Description: "desc...",
		Release: "1",
		Group: "group",
		Licence: "lic",
		Epoch: uint32(Time.Now().Unix()),
		BuildHost: "asd",
	});
	if err != nil { return err; }
	data, err := IOUtils.ReadFile("weld");
	if err != nil { return err; }
Fmt.Printf("\nlen: %d\n", len(data));
	rpm.AddFile(RPM.RPMFile{
		Name: "/usr/local/hello",
		Body: data,
		Type: RPM.GenericFile,
		Owner: "lop",
		Group: "lop",
		Mode: 0555,
	});
	out, err := OS.Create("/zcode/may/tools/pxnWelder/weld/test.rpm");
	if err != nil { return err; }
	defer out.Close();
	if err := rpm.Write(out); err != nil { return err; }
	print("Hello! Done!!!1243\n");
	return nil;
}
