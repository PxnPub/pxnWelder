package app;
// pxnWelder CI Service

import(
//	Log     "log"
	Flag    "flag"
//	Flagz   "github.com/PoiXson/pxnGoCommon/utils/flagz"
	PxnServ "github.com/PoiXson/pxnGoCommon/service"
//	PxnFS   "github.com/PoiXson/pxnGoCommon/utils/fs"
//	Configs "github.com/PoiXson/pxnWelder/ci/configs"
);



type AppWelderCI struct {
	service *PxnServ.Service
//	config  *Configs.CfgWelderCI
}



func New() PxnServ.AppFace {
	return &AppWelderCI{};
}

func (app *AppWelderCI) Main() {
	app.service = PxnServ.New();
	app.service.Start();
	app.flags_and_configs(DefaultConfigFile);



}



func (app *AppWelderCI) flags_and_configs(file string) {
	Flag.Parse();
}
