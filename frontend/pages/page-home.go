package pages;
// pxnWelder CI Frontend

import(
	HTTP   "net/http"
	PxnWeb "github.com/PoiXson/pxnGoCommon/net/web"
);

import _ "embed";



//go:embed page-home.tpl
var TPL_Home []byte;



type PageHome struct {
	Pages   *Pages
	Builder *PxnWeb.Builder
}



func (pages *Pages) NewPageHome() *PageHome {
	return &PageHome{
		Pages:   pages,
		Builder: pages.Builder.Clone().
				AddRawTPL(TPL_Home),
	};
}



func (page *PageHome) RenderWeb(out HTTP.ResponseWriter, in *HTTP.Request) {
	out.Header().Set("Content-Type", PxnWeb.Mime_HTML);
	tags := page.Builder.CloneTags();
	tags[PxnWeb.Tag_WithBootstrap] = true;
	page.Builder.TPL.ExecuteTemplate(out, "website", tags);
}
