package template_repo

//
//go:generate gormgen -structs Template -input .
type Template struct {
	Id                 int32  //
	Name               string // 项目模板名
	ControllerTemplate string //
	ServiceTemplate    string //
	ModelTemplate      string //
	RouterTemplate     string //
	PojoTemplate       string //
	ProjectNamespace   string //
}
