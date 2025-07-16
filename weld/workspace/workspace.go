package workspace;

import(
	Configs "github.com/PoiXson/pxnWelder/weld/configs"
);



type Workspace struct {
	Path string
	Yaml *Configs.WeldYml
}



func NewWorkspace(path string) *Workspace {
//TODO: make temp workspace dir here?
//TODO: load weld.yml
yaml := &Configs.WeldYml{};
	return &Workspace{
		Path: path,
		Yaml: yaml,
	};
}



























































































