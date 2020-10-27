package patchy

type Differ interface {
	Diff(string, string, string) Patch
}

type Patch struct {
	RepositoryName string
	Diff           string
	TargetBranch   string
	OriginBranch   string
}
