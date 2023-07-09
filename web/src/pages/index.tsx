import { useEffect } from 'react'
import { useRouter } from 'next/router'
import Spinner from 'src/@core/components/spinner'

const Home = () => {
  const router = useRouter()

  useEffect(() => {
    if (router.route === '/') {
      router.replace('/login')
    }
  }, [router])

  return <Spinner sx={{ height: '100%' }} />
}

export default Home
