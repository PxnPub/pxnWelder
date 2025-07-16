package app;
// pxnWelder Badge

import(
	Log     "log"
	Flag    "flag"
//	Flagz   "github.com/PoiXson/pxnGoCommon/utils/flagz"
	PxnFS   "github.com/PoiXson/pxnGoCommon/utils/fs"
	PxnServ "github.com/PoiXson/pxnGoCommon/service"
	Configs "github.com/PoiXson/pxnWelder/badge/configs"
);



type AppBadge struct {
	service *PxnServ.Service
	config  *Configs.CfgBadge
}



func New() PxnServ.AppFace {
	return &AppBadge{};
}

func (app *AppBadge) Main() {
	app.service = PxnServ.New();
	app.service.Start();
	app.flags_and_configs(DefaultConfigFile);



}



func (app *AppBadge) flags_and_configs(file string) {
	Flag.Parse();
	// load config
	cfg, err := PxnFS.LoadConfig[Configs.CfgBadge](file);
	if err != nil { Log.Panicf("%s, when loading config %s", err, file); }
	app.config = cfg;
}
