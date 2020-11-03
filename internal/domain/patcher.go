package domain

type Patcher interface {
	Patch(PatchRequest) (Patch, error)
}
