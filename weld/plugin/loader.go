package plugin;

import(
	Log    "log"
	Plugin "plugin"
);



func Load(file string) WeldPlugin {
	Log.Printf("Loading Plugin: %s", file);
	plugin_so, err := Plugin.Open(file);
	if err != nil { Log.Panic(err); }
	plugin_const, err := plugin_so.Lookup("NewPlugin");
	if err != nil { Log.Panic(err); }
	plugin_construct := plugin_const.(func() WeldPlugin);
	return plugin_construct();
}
