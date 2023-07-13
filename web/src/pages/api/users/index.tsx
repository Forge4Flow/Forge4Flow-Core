import { NextApiRequest, NextApiResponse } from 'next/types'
import { withSessionPermission, Auth4FlowServer } from '@auth4flow/auth4flow-nextjs'
import { use } from 'next-api-route-middleware'

const users = async (req: NextApiRequest, res: NextApiResponse) => {
  const auth4flow = new Auth4FlowServer({
    endpoint: process.env.AUTH4FLOW_BASE_URL,
    apiKey: process.env.AUTH4FLOW_API_KEY || ''
  })

  const users = await auth4flow.User.listUsers()

  if (users) {
    res.status(200).json(users)
    return
  }

  res.status(200).json({ message: 'Hello, world!' })
}

export default use(withSessionPermission('auth4flow-admin'), users)
