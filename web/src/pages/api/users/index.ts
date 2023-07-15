import { NextApiRequest, NextApiResponse } from 'next/types'
import { withSessionPermission, Forge4FlowServer } from '@forge4flow/forge4flow-nextjs'
import { use } from 'next-api-route-middleware'

const users = async (req: NextApiRequest, res: NextApiResponse) => {
  const forge4flow = new Forge4FlowServer({
    endpoint: process.env.AUTH4FLOW_BASE_URL,
    apiKey: process.env.AUTH4FLOW_API_KEY || ''
  })

  const users = await forge4flow.User.listUsers()

  if (users) {
    res.status(200).json(users)
    return
  }

  res.status(200).json({ message: 'Hello, world!' })
}

export default use(withSessionPermission('forge4flow-admin'), users)
