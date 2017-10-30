package abstraction

// Task represents task interface of workflow.
type Task interface {
	Execute() error
	ExecuteParallel(value string) error
}
