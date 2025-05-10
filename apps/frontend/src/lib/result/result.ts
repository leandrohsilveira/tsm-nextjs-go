export type OkResult<R> = [null, R]

export type ErrorResult<E> = [E, null]

export type Result<E, R> = ErrorResult<E> | OkResult<R>

export function ok<R>(data: R): OkResult<R> {
  return [null, data]
}

export function err<E>(error: E): ErrorResult<E> {
  return [error, null]
}
