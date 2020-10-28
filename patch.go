package patchy

type Patch struct {
	RepositoryName string
	Diff           string
	TargetBranch   string
	OriginBranch   string
}

type Patcher interface {
	Patch(PatchRequest) (Patch, error)
}

type PatchRequest struct {
	Repository   string
	OriginBranch string
	TargetBranch string
}
