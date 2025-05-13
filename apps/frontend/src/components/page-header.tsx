'use server'

import { withAuthData } from '@/lib/auth'

export async function PageHeader() {
  const auth = await withAuthData()
  return (
    <header className="w-full flex flex-col md:flex-row items-center justify-between gap-4 bg-foreground/10 py-4 px-8 rounded-full">
      <h1 className="text-2xl font-bold">TSM</h1>
      {auth ? <span>{auth.data.name}</span> : <a href="/login">Login</a>}
    </header>
  )
}
