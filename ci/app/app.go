package app;
// pxnWelder CI Service

import(
//	Log     "log"
	Flag    "flag"
//	Flagz   "github.com/PoiXson/pxnGoCommon/utils/flagz"
	Service "github.com/PoiXson/pxnGoCommon/service"
//	UtilsFS "github.com/PoiXson/pxnGoCommon/utils/fs"
//	Configs "github.com/PoiXson/pxnWelder/ci/configs"
);



type AppWelderCI struct {
	service *Service.Service
//	config  *Configs.CfgWelderCI
}



func New() Service.AppFace {
	return &AppWelderCI{};
}

func (app *AppWelderCI) Main() {
	app.service = Service.New();
	app.service.Start();
	app.flags_and_configs(DefaultConfigFile);



}



func (app *AppWelderCI) flags_and_configs(file string) {
	Flag.Parse();
}
