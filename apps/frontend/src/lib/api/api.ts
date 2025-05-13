import { ApiResponseError } from '../error'
import { err, ok, Result } from '../result'

type ApiCallMethod =
  | 'get'
  | 'post'
  | 'put'
  | 'patch'
  | 'delete'
  | 'options'
  | 'head'

type BodySupplier = {
  parse: () => Promise<BodyInit | null | undefined>
  contentType?: string
}

type ResponseParser<T> = {
  accept?: string
  parse: (response: Response) => Promise<Result<ApiResponseError, T>>
}

type ApiCallOptions = {
  headers?: Headers
  query?: URLSearchParams
  baseUrl?: string
}

const config = {
  fetch,
  baseUrl: 'http://localhost:4000',
}

async function call<T>(
  method: ApiCallMethod,
  path: string,
  { contentType, parse: reqParser }: BodySupplier,
  { accept, parse: resParser }: ResponseParser<T>,
  options: ApiCallOptions = {},
) {
  const headers = options.headers ?? new Headers()

  let url = (options.baseUrl ?? config.baseUrl) + path
  if (options.query) url += '?' + options.query.toString()
  if (contentType) headers.set('content-type', contentType)
  if (accept) headers.set('accept', accept)

  try {
    const response = await config.fetch(url, {
      method,
      headers,
      body: await reqParser(),
    })
    const responseContentType = response.headers?.get('content-type')
    if (
      accept &&
      responseContentType &&
      !responseContentType.startsWith(accept)
    ) {
      const responseText = await response.text()
      return err(
        new ApiResponseError({
          message: `Response content-type mismatch. (Expected: ${accept}, Received: ${responseContentType}) Body: ${responseText}`,
        }),
      )
    }

    const [error, result] = await resParser(response)

    if (error != null) return err(error)

    return ok(result)
  } catch (error) {
    if (typeof error === 'object' && error !== null && 'message' in error) {
      return err(new ApiResponseError({ message: String(error.message) }))
    }
    return err(new ApiResponseError({ message: String(error) }))
  }
}

export const parsers = {
  request: {
    noContent: { parse: async () => undefined } satisfies BodySupplier,
    json: (body: unknown) =>
      ({
        contentType: 'application/json',
        parse: async () => JSON.stringify(body),
      }) satisfies BodySupplier,
  },
  response: {
    noContent: {
      parse: async () => ok(undefined),
    } satisfies ResponseParser<undefined>,
    json: <T>(): ResponseParser<T> => ({
      accept: 'application/json',
      parse: async (response: Response) => {
        try {
          if (!response.ok)
            return err(new ApiResponseError(await response.json()))
          return ok<T>(await response.json())
        } catch (error) {
          return err(new ApiResponseError({ message: String(error) }))
        }
      },
    }),
  },
}

export const api = {
  config,
  call,
  get<T>(path: string, res: ResponseParser<T>, options?: ApiCallOptions) {
    return call('get', path, parsers.request.noContent, res, options)
  },
  post<T>(
    path: string,
    req: BodySupplier,
    res: ResponseParser<T>,
    options?: ApiCallOptions,
  ) {
    return call('post', path, req, res, options)
  },
  put<T>(
    path: string,
    req: BodySupplier,
    res: ResponseParser<T>,
    options?: ApiCallOptions,
  ) {
    return call('put', path, req, res, options)
  },
  patch<T>(
    path: string,
    req: BodySupplier,
    res: ResponseParser<T>,
    options?: ApiCallOptions,
  ) {
    return call('patch', path, req, res, options)
  },
  delete<T>(
    path: string,
    req: BodySupplier,
    res: ResponseParser<T>,
    options?: ApiCallOptions,
  ) {
    return call('delete', path, req, res, options)
  },
}
