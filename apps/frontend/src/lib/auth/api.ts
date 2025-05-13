'use server'

import { cookies } from 'next/headers'
import { api, parsers } from '../api'
import type { AuthData, AuthLoginPayload, AuthLoginResult } from './schema'

export async function withAuth() {
  const store = await cookies()
  const authCookie = store.get('session')

  if (!authCookie) return [null, set] as const

  return [
    { token: authCookie.value, refreshToken: '' } satisfies AuthLoginResult,
    set,
  ] as const

  function set({ token }: AuthLoginResult) {
    store.set('session', token, {
      sameSite: 'lax',
      path: '/',
      secure: true,
      httpOnly: true,
    })
  }
}

export async function withAuthHeader(headers = new Headers()) {
  const [auth] = await withAuth()
  if (!auth) return headers
  headers.set('authorization', auth.token) // TODO: use bearer token
  return headers
}

export async function withAuthData(): Promise<AuthData | null> {
  const [err, data] = await info()

  if (err != null) return null

  return data
}

export async function login(data: AuthLoginPayload) {
  return await api.post(
    '/auth',
    parsers.request.json(data),
    parsers.response.json<AuthLoginResult>(),
  )
}

export async function info() {
  const headers = await withAuthHeader()
  return await api.get('/auth', parsers.response.json<AuthData>(), { headers })
}
