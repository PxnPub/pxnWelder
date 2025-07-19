package app;
// weld tool

import(
//	OS        "os"
	Log       "log"
	Fmt       "fmt"
	Flag      "flag"
	Flagz     "github.com/PoiXson/pxnGoCommon/utils/flagz"
//	PxnUtils  "github.com/PoiXson/pxnGoCommon/utils"
	PxnFS     "github.com/PoiXson/pxnGoCommon/utils/fs"
	PxnServ   "github.com/PoiXson/pxnGoCommon/service"
	Configs   "github.com/PoiXson/pxnWelder/weld/configs"
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Worker    "github.com/PoiXson/pxnWelder/weld/worker"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
	_         "github.com/PoiXson/pxnGoLog/logger/pxnlog"
);



type AppWeldTool struct {
//	service   *PxnServ.Service
	config    *Configs.CfgWeldTool
	IsDebug   bool
	IsVerbose bool
	// paths
	PathPlugins   string
	PathTemplates string
}



func New() PxnServ.AppFace {
	return &AppWeldTool{};
}

func (app *AppWeldTool) Main() {
	app.load_config(DefaultConfigFile);
	app.load_flags();
	// defaults
	if app.PathPlugins   == "" { app.PathPlugins   = DefaultPathPlugins;   }
	if app.PathTemplates == "" { app.PathTemplates = DefaultPathTemplates; }
	// load plugins
	plugins := WeldPlug.LoadPath(app.PathPlugins);
	// load first workspace
	workspace, err := Workspace.New("/zcode/may/tools/pxnWelder/weld");
	if err != nil { Log.Panic(err); }
	//LOOP_ARGS:
	for _, action := range Flag.Args() {
		action_count := Worker.Plugins_Run(plugins, workspace, action);
		if action_count == 0 {
			Fmt.Printf("Nothing to do: %s\n", action);
		} else {
			s := "";
			if action_count != 1 { s = "s"; }
			Fmt.Printf("Performed %d action%s: %s\n", action_count, s, action);
		}
		print("\n");
	}
	print("\nFinished!\n");
}



func (app *AppWeldTool) load_config(file string) {
	if !PxnFS.IsFile(file) {
		Log.Printf("Config file not found, will be ignored: %s", file);
		return;
	}
	cfg, err := PxnFS.LoadConfig[Configs.CfgWeldTool](file);
	if err != nil { Log.Panicf("%s, when loading config %s", err, file); }
	if app.PathPlugins   == "" { app.PathPlugins   = cfg.PathPlugins;   }
	if app.PathTemplates == "" { app.PathTemplates = cfg.PathTemplates; }
	app.config = cfg;
}

func (app *AppWeldTool) load_flags() {
	var flag_debug   bool; var flag_d bool;
	var flag_verbose bool; var flag_v bool;
	var flag_path_plugins   string;
	var flag_path_templates string;
	Flagz.Bool(&flag_debug,   "debug"  ); Flagz.Bool(&flag_d, "d");
	Flagz.Bool(&flag_verbose, "verbose"); Flagz.Bool(&flag_v, "v");
	Flagz.String(&flag_path_plugins,   "path-plugins",   "");
	Flagz.String(&flag_path_templates, "path-templates", "");
	Flag.Parse();
	if flag_debug   || flag_d { app.IsDebug   = true; }
	if flag_verbose || flag_v { app.IsVerbose = true; }
	if flag_path_plugins   != "" { app.PathPlugins   = flag_path_plugins;   }
	if flag_path_templates != "" { app.PathTemplates = flag_path_templates; }
}
