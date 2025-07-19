package workspace;

import(
	Log      "log"
	Fmt      "fmt"
	FilePath "path/filepath"
	PxnFS    "github.com/PoiXson/pxnGoCommon/utils/fs"
	Configs  "github.com/PoiXson/pxnWelder/weld/configs"
);



type Workspace struct {
	Path string
	Yml  *Configs.WeldYml
}



func New(path string) (*Workspace, error) {
//TODO: make temp workspace dir here?
	// load weld.yml
	file := FilePath.Join(path, "weld.yml");
	if !PxnFS.IsFile(file) {
		return nil, Fmt.Errorf("File not found: %s", file); }
	Log.Printf("Loading file: %s", file);
	yml, err := PxnFS.LoadConfig[Configs.WeldYml](file);
	if err != nil { return nil, err; }
	return &Workspace{
		Path: path,
		Yml:  yml,
	}, nil;
}
