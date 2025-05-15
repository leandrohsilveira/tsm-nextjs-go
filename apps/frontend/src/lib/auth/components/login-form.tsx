'use client'

import { Field, Message } from '@/components'
import { loginAction } from '@/lib/auth'
import { useActionState } from 'react'

export function LoginForm() {
  const [state, action, pending] = useActionState(loginAction, {
    message: null,
    errors: {},
  })

  return (
    <>
      <Message text={state.message} kind="error" />
      <form className="w-full flex flex-col gap-4" action={action}>
        <Field id="username" label="E-mail" errors={state.errors.username}>
          <input
            type="email"
            id="username"
            name="username"
            placeholder="Enter e-mail"
            autoComplete="username"
            disabled={pending}
            required
          />
        </Field>
        <Field id="password" label="Password" errors={state.errors.password}>
          <input
            type="password"
            id="password"
            name="password"
            placeholder="Enter password"
            autoComplete="current-password"
            disabled={pending}
            required
          />
        </Field>
        <button type="submit" disabled={pending}>
          Login
        </button>
      </form>
    </>
  )
}
