import { z } from 'zod'
import { UserData } from '../user'
import { err, ok } from '../result'
import { ValidationError } from '../error'

export interface AuthData {
  data: UserData
}

export interface AuthLoginResult {
  token: string
  refreshToken: string
}

export const AuthLoginPayload = z.object({
  username: z.string().email(),
  password: z.string().min(6),
})

export type AuthLoginPayload = z.infer<typeof AuthLoginPayload>

export function parseAuthLoginPayload(form: FormData) {
  const { success, error, data } = AuthLoginPayload.safeParse({
    username: form.get('username'),
    password: form.get('password'),
  })

  if (!success) return err(ValidationError.fromError(error))

  return ok(data)
}
