import { z } from "zod"
import { UserData } from "../user"

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
