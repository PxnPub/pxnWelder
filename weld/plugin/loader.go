package plugin;

import(
	OS       "os"
	Fmt      "fmt"
	Plugin   "plugin"
	FilePath "path/filepath"
	PxnFS    "github.com/PoiXson/pxnGoCommon/utils/fs"
	Log      "github.com/PoiXson/pxnGoLog/logger"
);



func LoadPath(path string) []WeldPlugin {
	if len(path) < 2      { Log.Panicf("Invalid plugins path"); }
	if !PxnFS.IsDir(path) { Log.Panicf("Plugins path not found: %s", path); }
	// single file
	{
		info, err := OS.Stat(path);
		if err != nil { Log.Panic(err); }
		if info.Mode().IsRegular() {
			plugins := make([]WeldPlugin, 1);
			plug, err := LoadPlugin(path);
			if err != nil { Log.Panicf("%s, in WeldPlugin::LoadPath()", err); }
			plugins[0] = plug;
			return plugins;
		}
	}
	// scan dir
	{
		plugins := make([]WeldPlugin, 0);
		if err := FilePath.Walk(path, func(filepath string,
				info OS.FileInfo, err error) error {
			if err != nil { Log.Panic(err); }
			if !info.IsDir() && FilePath.Ext(filepath) == ".so" {
				plug, err := LoadPlugin(filepath);
				if err != nil { Log.Panicf("%s, in WeldPlugin::LoadPath()", err); }
				plugins = append(plugins, plug);
			}
			return nil;
		}); err != nil { Log.Panicf("%s, in WeldPlugin::LoadPath()", err); }
		return plugins;
	}
}

func LoadPlugin(file string) (WeldPlugin, error) {
	Log.Printf("Loading Plugin: %s", file);
	plugin_so, err := Plugin.Open(file);
	if err != nil { return nil, Fmt.Errorf("%v, in LoadPlugin()", err); }
	plugin_const, err := plugin_so.Lookup("NewPlugin");
	if err != nil { return nil, Fmt.Errorf("%v, in LoadPlugin(), failed to find NewPlugin()", err); }
	plugin_construct, ok := plugin_const.(func() WeldPlugin);
	if !ok { return nil, Fmt.Errorf("%s ::NewPlugin() has incorrect signature", file); }
	return plugin_construct(), nil;
}
