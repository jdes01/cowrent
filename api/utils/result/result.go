package result

type Result[T any] struct {
	Ok  T
	Err error
}

func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

func (r Result[T]) IsErr() bool {
	return r.Err != nil
}

type NextFunction[T any, U any] func(T) Result[U]

func IfOkRun[T any, U any](result Result[T], fn func(T) Result[U]) Result[U] {
	if result.IsErr() {
		return Result[U]{Err: result.Err}
	}
	return fn(result.Ok)
}

func NewResult[T any](val T, err error) Result[T] {
	return Result[T]{val, err}
}
