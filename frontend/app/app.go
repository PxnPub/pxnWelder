package app;
// pxnWelder CI Frontend

import(
	Log     "log"
	Flag    "flag"
	Flagz   "github.com/PoiXson/pxnGoCommon/utils/flagz"
	PxnFS   "github.com/PoiXson/pxnGoCommon/utils/fs"
	PxnUtil "github.com/PoiXson/pxnGoCommon/utils"
	PxnWeb  "github.com/PoiXson/pxnGoCommon/net/web"
	PxnServ "github.com/PoiXson/pxnGoCommon/service"
	WebLink "github.com/PoiXson/pxnWelder/frontend/weblink"
	Configs "github.com/PoiXson/pxnWelder/frontend/configs"
	Pages   "github.com/PoiXson/pxnWelder/frontend/pages"
);



type AppFront struct {
	service *PxnServ.Service
	websvr  *PxnWeb.WebServer
	pages   *Pages.Pages
	link    *WebLink.WebLink
	config  *Configs.CfgFront
}



func New() PxnServ.AppFace {
	return &AppFront{};
}

func (app *AppFront) Main() {
	app.service = PxnServ.New();
	app.service.Start();
	app.flags_and_configs(DefaultConfigFile);
	// rpc
	app.link = WebLink.New(app.service, app.config.BrokerAddr);
	// web server
	app.websvr = PxnWeb.NewWebServer(
		app.service,
		app.config.BindWeb,
		app.config.Proxied,
	);
	router := app.websvr.WithGorilla();
	app.pages = Pages.New(app.link, router);
	// start things
	if err := app.link  .Start(); err != nil { Log.Panic(err); }
	if err := app.websvr.Start(); err != nil { Log.Panic(err); }
	// delay rpc close
	app.service.AddStopHook(func() { go func() {
		PxnUtil.SleepCn(5);
		app.link.Close();
	}(); });
	app.service.WaitUntilEnd();
}



func (app *AppFront) flags_and_configs(file string) {
	var flag_broker  string;
	var flag_bindweb string;
	var flag_proxied bool;
	Flagz.String(&flag_broker,  "broker", "");
	Flagz.String(&flag_bindweb, "bind",   "");
	Flagz.Bool  (&flag_proxied, "proxied"   );
	Flag.Parse();
	// load config
	cfg, err := PxnFS.LoadConfig[Configs.CfgFront](file);
	if err != nil { Log.Panicf("%s, when loading config %s", err, file); }
	// remote rpc
	if flag_broker    != "" { cfg.BrokerAddr = flag_broker;          }
	if cfg.BrokerAddr == "" { cfg.BrokerAddr = DefaultBrokerAddress; }
	// bind web
	if flag_bindweb != "" { cfg.BindWeb = flag_bindweb;          }
	if cfg.BindWeb  == "" { cfg.BindWeb = PxnWeb.DefaultBindWeb; }
	if flag_proxied       { app.config.Proxied = true;           }
	app.config = cfg;
}
