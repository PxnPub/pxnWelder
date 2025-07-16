package plugin;

import(
	OS       "os"
	Log      "log"
	FilePath "path/filepath"
	Plugin   "plugin"
);



func LoadPath(path string) []WeldPlugin {
	if len(path) < 2 { Log.Panic("Invalid plugins path"); }
	// single file
	{
		info, err := OS.Stat(path);
		if err != nil { Log.Panic(err); }
		if info.Mode().IsRegular() {
			plugins := make([]WeldPlugin, 1);
			plugins[0] = LoadPlugin(path);
			return plugins;
		}
	}
	// scan dir
	{
		plugins := make([]WeldPlugin, 0);
		err := FilePath.Walk(path, func(filepath string,
				info OS.FileInfo, err error) error {
			if err != nil { Log.Panic(err); }
			if !info.IsDir() && FilePath.Ext(filepath) == ".so" {
				p := LoadPlugin(filepath);
				plugins = append(plugins, p);
			}
			return nil;
		});
		if err != nil { Log.Panic(err); }
		return plugins;
	}
}

func LoadPlugin(file string) WeldPlugin {
	Log.Printf("Loading Plugin: %s", file);
	plugin_so, err := Plugin.Open(file);
	if err != nil { Log.Panic(err); }
	plugin_const, err := plugin_so.Lookup("NewPlugin");
	if err != nil { Log.Panic(err); }
	plugin_construct := plugin_const.(func() WeldPlugin);
	return plugin_construct();
}
