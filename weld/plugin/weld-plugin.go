package plugin;

import(
	Work "github.com/PoiXson/pxnWelder/weld/worker"
);



//type WeldStage uint8;
//const (
//	Stage_Clean  WeldStage = iota
//	Stage_Config
//	Stage_Build
//	Stage_Package
//);



type WeldPlugin interface {
	Run(work *Work.Workspace, stage string) error
}
