package updatepackage

type UpdatePackage struct {
	Name         string
	UpdateMethod string
	PreTasks     []string
	Files        []string
	AfterTasks   []string
}

type Segment struct {
	Index     int
	Path      string
	Signature string
}

type UpdatePackageDeployment struct {
	Segments  []Segment
	Signature string
}