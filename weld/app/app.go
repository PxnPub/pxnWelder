package app;
// weld tool

import(
//	OS       "os"
//	Log      "log"
//	Fmt      "fmt"
	Flag     "flag"
	Flagz    "github.com/PoiXson/pxnGoCommon/utils/flagz"
//	Utils    "github.com/PoiXson/pxnGoCommon/utils"
//	PxnFS    "github.com/PoiXson/pxnGoCommon/utils/fs"
	PxnServ  "github.com/PoiXson/pxnGoCommon/service"
//	Configs  "github.com/PoiXson/pxnWelder/weld/configs"
	WeldPlug "github.com/PoiXson/pxnWelder/weld/plugin"
);



type AppWeldTool struct {
//	service *PxnServ.Service
//	config  *Configs.CfgWeldTool
	IsDebug bool
}



func New() PxnServ.AppFace {
	return &AppWeldTool{};
}

func (app *AppWeldTool) Main() {
	app.flags_and_configs(DefaultConfigFile);

//	LOOP_ARGS:
//	for i, arg := range Flag.Args() {
//		Fmt.Printf(" Arg: %d %s\n", i, arg);
//	}

//	app.service = PxnServ.New();
//	app.service.Start();

// test load a plugin
plugin := WeldPlug.Load("../plugins/rpm/weld-rpm.so");
workspace := Work.NewWorkspace();
stage := "clean";
if err := plugin.Run(workspace, stage); err != nil { panic(err); }

}



func (app *AppWeldTool) flags_and_configs(file string) {
	var flag_debug bool; var flag_d bool;
	Flagz.Bool(&flag_debug, "debug");
	Flagz.Bool(&flag_d,     "d"    );
	Flag.Parse();
	if flag_d || flag_debug { app.IsDebug = true; }
//	// load config
//	cfg, err := UtilsFS.LoadConfig[Configs.CfgWeldTool](file);
//	if err != nil { Log.Panicf("%s, when loading config %s", err, file); }
//	app.config = cfg;
}
