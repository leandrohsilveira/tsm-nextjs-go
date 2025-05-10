"use server"

import { ValidationMessages } from "../error";
import { login } from "./api";
import { AuthLoginPayload, parseAuthLoginPayload } from "./schema";

export async function loginAction(_: unknown, form: FormData): Promise<{ message: string | null, errors: ValidationMessages<AuthLoginPayload> }> {
  const [validationError, data] = parseAuthLoginPayload(form)
  if (validationError !== null) return { message: null, errors: validationError.errors }

  const [apiError, result] = await login(data)
  if (apiError !== null) return { message: apiError.error.message, errors: {} }

  console.log("Login sucessful", result)

  return { message: null, errors: {} }
}
