'use server'

import { api, parsers } from '../api'
import type { AuthData, AuthLoginPayload, AuthLoginResult } from './schema'

export async function login(data: AuthLoginPayload) {
  return await api.post(
    '/auth',
    parsers.request.json(data),
    parsers.response.json<AuthLoginResult>(),
  )
}

export async function info(headers: Headers) {
  return await api.get('/auth', parsers.response.json<AuthData>(), { headers })
}
