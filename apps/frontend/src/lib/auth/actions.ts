"use server"

import { cookies } from "next/headers";
import { ValidationMessages } from "../error";
import { login } from "./api";
import { AuthLoginPayload, parseAuthLoginPayload } from "./schema";
import { redirect } from "next/navigation";

export async function loginAction(_: unknown, form: FormData): Promise<{ message: string | null, errors: ValidationMessages<AuthLoginPayload> }> {
  const [validationError, data] = parseAuthLoginPayload(form)
  if (validationError !== null) return { message: null, errors: validationError.errors }

  const [apiError, result] = await login(data)
  if (apiError !== null) return { message: apiError.error.message, errors: {} }

  const cookiesStorage = await cookies()

  cookiesStorage.set('session', result.token, {
    sameSite: 'lax',
    path: '/',
    secure: true,
    httpOnly: true,
  })

  redirect('/')
}
