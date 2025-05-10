'use client'

import { Field, PageLayout } from "@/components";
import { loginAction } from "@/lib/auth";
import { useActionState } from "react";


export default function Login() {
  const [, action, pending] = useActionState(
    loginAction,
    { message: null, errors: {} }
  )

  return (
    <PageLayout>
      <form className="w-full flex flex-col gap-4" action={action}>
        <Field id="username" label="E-mail">
          <input type="text" id="username" name="username" placeholder="Enter e-mail" inputMode="email" autoComplete="username" disabled={pending} />
        </Field>
        <Field id="password" label="Password" required>
          <input type="password" id="password" name="password" placeholder="Enter password" autoComplete="current-password" disabled={pending} />
        </Field>
        <button type="submit" disabled={pending}>Login</button>
      </form>
    </PageLayout>
  )
}
