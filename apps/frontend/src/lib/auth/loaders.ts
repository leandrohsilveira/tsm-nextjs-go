'use server'

import { cookies } from 'next/headers'
import { info } from './api'
import type { AuthData, AuthLoginResult } from './schema'

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
  const headers = await withAuthHeader()

  const [err, data] = await info(headers)

  if (err != null) return null

  return data
}
