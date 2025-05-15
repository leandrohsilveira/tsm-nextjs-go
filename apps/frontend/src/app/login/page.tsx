import { PageLayout } from '@/components'
import { LoginForm } from '@/lib/auth/components'

export default function Login() {
  return (
    <PageLayout hideHeader>
      <h3 className="font-bold text-2xl text-center w-full">Sign-in</h3>
      <LoginForm />
    </PageLayout>
  )
}
