package plugin;

import(
	Workspace "github.com/PoiXson/pxnWelder/weld/workspace"
);



//type WeldStage uint8;
//const (
//	Stage_Clean  WeldStage = iota
//	Stage_Config
//	Stage_Build
//	Stage_Package
//);



type WeldPlugin interface {
	GetName() string
	GetWeight(stage string) uint8
	Run(workspace *Workspace.Workspace, stage string) error
}
