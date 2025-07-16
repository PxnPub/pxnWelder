package pages;
// pxnWelder CI Frontend

import(
	Gorilla "github.com/gorilla/mux"
	PxnWeb  "github.com/PoiXson/pxnGoCommon/net/web"
	WebLink "github.com/PoiXson/pxnWelder/frontend/weblink"
);

import _ "embed";



//go:embed menu.tpl
var TPL_Menu []byte;



type Pages struct {
	Link     *WebLink.WebLink
	Builder  *PxnWeb.Builder
	PageHome *PageHome
}



func New(link *WebLink.WebLink, router *Gorilla.Router) *Pages {
	builder := PxnWeb.NewBuilder().
		WithIncludes().
		WithBootstrap().
		WithBootsIcons().
		WithTooltips().
		AddRawTPL(TPL_Menu).
		SetTitle("WelderCI").
		AddFilesCSS("/static/welder.css").
		SetFavIcon("/static/welder.ico");
	pages := Pages{
		Link:    link,
		Builder: builder,
	};
	pages.PageHome = pages.NewPageHome();
	PxnWeb.AddStaticRoute(router);
	router.HandleFunc("/", pages.PageHome.RenderWeb);
	return &pages;
}
