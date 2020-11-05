package domain

type PatchUsecase interface {
	Patch(PatchRequest) (Patch, error)
}
