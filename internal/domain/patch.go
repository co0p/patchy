package domain

type Patch struct {
	RepositoryName string
	Diff           string
	TargetBranch   string
	OriginBranch   string
}

type PatchRequest struct {
	Repository   string
	OriginBranch string
	TargetBranch string
}

func (r PatchRequest) Valid() bool {
	return len(r.Repository) > 0 && len(r.OriginBranch) > 0 && len(r.TargetBranch) > 0
}
