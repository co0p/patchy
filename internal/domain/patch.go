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
