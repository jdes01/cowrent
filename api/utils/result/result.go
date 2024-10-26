package result

import (
	"api/utils/logger"
)

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

func (r Result[T]) LogIfErr(extras ...interface{}) Result[T] {
	if r.IsErr() {
		logger.GetLogger().Error(r.Err.Error(), extras)
	}
	return r
}

func (r Result[T]) LogThisIfErr(customMessage string, extras ...interface{}) Result[T] {
	if r.IsErr() {
		logger.GetLogger().Error(customMessage, append(extras, r.Err))
	}
	return r
}

func MapError[T any, MappedType any](result Result[T]) Result[MappedType] {
	var empty MappedType
	return NewResult(empty, result.Err)
}

func (r Result[T]) LogThisIfOk(customMessage string, extras ...interface{}) Result[T] {
	if r.IsOk() {
		logger.GetLogger().Info(customMessage, append(extras, r.Err))
	}
	return r
}
