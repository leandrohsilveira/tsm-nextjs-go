'use client'

import { Field, PageLayout } from "@/components";
import { loginAction } from "@/lib/auth";
import { useActionState } from "react";


export default function Login() {
  const [state, action, pending] = useActionState(
    loginAction,
    { message: null, errors: {} }
  )

  return (
    <PageLayout>
      <h3 className="font-bold text-2xl text-center w-full">Sign-in</h3>
      <form className="w-full flex flex-col gap-4" action={action}>
        <Field id="username" label="E-mail" errors={state.errors.username} required>
          <input type="text" id="username" name="username" placeholder="Enter e-mail" inputMode="email" autoComplete="username" disabled={pending} />
        </Field>
        <Field id="password" label="Password" errors={state.errors.password} required>
          <input type="password" id="password" name="password" placeholder="Enter password" autoComplete="current-password" disabled={pending} />
        </Field>
        <button type="submit" disabled={pending}>Login</button>
      </form>
    </PageLayout>
  )
}
