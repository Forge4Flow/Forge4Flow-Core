// ** React Imports
import { ReactNode, ReactElement, useEffect } from 'react'

// ** Next Import
import { useRouter } from 'next/router'

// ** Hooks Import
import { useAuth4Flow } from '@auth4flow/auth4flow-react'

interface GuestGuardProps {
  children: ReactNode
  fallback: ReactElement | null
}

const GuestGuard = (props: GuestGuardProps) => {
  const { children, fallback } = props
  const auth = useAuth4Flow()
  const router = useRouter()

  useEffect(() => {
    if (!router.isReady) {
      return
    }

    const verifySession = async () => {
      const validSession = await auth.validSession()
      if (validSession) {
        router.push('/admin')
      }
    }

    verifySession()

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [auth.sessionToken, router.route])

  if (auth.isLoading || (!auth.isLoading && auth.sessionToken !== '')) {
    return fallback
  }

  return <>{children}</>
}

export default GuestGuard
