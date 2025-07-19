package worker;

import(
//	Fmt       "fmt"
	Math      "math"
	PxnUtils  "github.com/PoiXson/pxnGoCommon/utils"
	WeldPlug  "github.com/PoiXson/pxnWelder/weld/plugin"
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



func Plugins_Run(plugins []WeldPlug.WeldPlugin,
		workspace *Workspace.Workspace, stage string) uint32 {
	plugs := make([]WeldPlug.WeldPlugin, len(plugins));
	copy(plugs, plugins);
	var action_count  uint32 = 0;
	var weight_current uint8 = 0;
	LOOP_RUN:
	for {
		// find next weight
		var weight_next uint8 = Math.MaxUint8;
		var plugin_next WeldPlug.WeldPlugin = nil;
		var plugin_index int;
		for index, plugin := range plugs {
			w := plugin.GetWeight(stage);
			if           0 < w &&
			weight_current < w &&
			weight_next    > w {
				weight_next = w;
				plugin_next = plugin;
				plugin_index = index;
			}
		}
		if plugin_next == nil { break LOOP_RUN; }
		weight_current = weight_next;

//Fmt.Printf("Next Weight: %d %s\n", weight_next, plugin_next.GetName());

done := true; //TODO: run()
		// plugin finished stage
		if done {
			plugs = PxnUtils.RemoveIndex(plugs, plugin_index);
			action_count++;
		}
	}
	return action_count;
}
