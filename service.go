package libredis

// Service ...
type Service interface {
	ClassManager

	GetClassManager() ClassManager

	GetSourceManager() SourceManager
}
