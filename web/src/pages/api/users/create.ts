import { NextApiRequest, NextApiResponse } from 'next/types'
import { withSessionPermission, Forge4FlowServer, CreateUserParams } from '@forge4flow/forge4flow-nextjs'
import { use } from 'next-api-route-middleware'

const createUser = async (req: NextApiRequest, res: NextApiResponse) => {
  if (req.method !== 'POST') {
    res.status(405).json({ message: 'Method Not Allowed' })
  }

  const forge4flow = new Forge4FlowServer({
    endpoint: process.env.AUTH4FLOW_BASE_URL,
    apiKey: process.env.AUTH4FLOW_API_KEY || ''
  })

  const user = JSON.parse(req.body)

  if (user) {
    const userOptions: CreateUserParams = {}
    forge4flow.User.create()
  }

  res.status(200).json({ message: 'Hello, world!' })
}

export default use(withSessionPermission('forge4flow-admin'), createUser)
