"use server"

import { ValidationMessages } from "../error";
import { login } from "./login";
import type { AuthLoginPayload } from "./schema";

export async function loginAction(_: unknown, form: FormData): Promise<{ message: string | null, errors: ValidationMessages<AuthLoginPayload> }> {
  const [error, data] = await login(form)

  if (error !== null) switch (error.kind) {
    case "VALIDATION_ERROR":
      console.error("Validation error", error.errors)
      return { message: null, errors: error.errors }
    case "API_RESPONSE_ERROR":
      console.error("API Response error", error.error)
      return { message: error.error.message, errors: {} }
  }

  console.log("Login sucessful", data)

  return { message: null, errors: {} }
}
