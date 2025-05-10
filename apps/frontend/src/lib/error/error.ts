import { ZodError } from "zod";

interface StandardError<T extends string = string> {
  readonly kind: T
}

interface StandardApiErrorData {
  message: string
}

export type ValidationMessages<E> = {
  [K in keyof E]?: string[]
}

const VALIDATION_ERROR = 'VALIDATION_ERROR' as const
const API_RESPONSE_ERROR = 'API_RESPONSE_ERROR' as const

export class ValidationError<E> implements StandardError<typeof VALIDATION_ERROR> {
  constructor(public errors: ValidationMessages<E>) {
  }

  public readonly kind = VALIDATION_ERROR

  static fromError<E>({ formErrors }: ZodError<E>) {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    return new ValidationError<E>(formErrors.fieldErrors as any)
  }
}

export class ApiResponseError implements StandardError<typeof API_RESPONSE_ERROR> {
  constructor(public error: StandardApiErrorData) {
  }

  public readonly kind = API_RESPONSE_ERROR

  static async fromResponse(response: Response) {
    if (response.headers.get('content-type')?.startsWith('application/json')) {
      return new ApiResponseError(await response.json())
    }
    return new ApiResponseError({
      message: await response.text()
    })
  }
}

export const StandardErrors = {
  VALIDATION_ERROR: VALIDATION_ERROR,
  API_RESPONSE_ERROR: API_RESPONSE_ERROR
}
