package plugin;



type WeldStage uint8;
const (
	Stage_Clean  WeldStage = iota
	Stage_Config
	Stage_Build
	Stage_Package
);



type WeldPlugin interface {
	Run(WeldStage) error
}
