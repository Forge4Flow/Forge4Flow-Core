import { ReactNode, useEffect } from 'react'
import { useRouter } from 'next/router'
import Spinner from 'src/@core/components/spinner'
import BlankLayout from 'src/@core/layouts/BlankLayout'

const Home = () => {
  const router = useRouter()

  useEffect(() => {
    if (router.route === '/') {
      router.push('/login')
    }
  }, [router])

  return <Spinner sx={{ height: '100%' }} />
}

Home.getLayout = (page: ReactNode) => <BlankLayout>{page}</BlankLayout>
Home.guestGuard = true
export default Home
