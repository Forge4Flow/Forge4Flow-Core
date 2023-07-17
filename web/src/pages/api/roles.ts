import { NextApiRequest, NextApiResponse } from 'next/types'
import { withSessionPermission, Forge4FlowServer, CreateUserParams } from '@forge4flow/forge4flow-nextjs'

const roles = async (req: NextApiRequest, res: NextApiResponse) => {
  const forge4flow = new Forge4FlowServer({
    endpoint: process.env.AUTH4FLOW_BASE_URL,
    apiKey: process.env.AUTH4FLOW_API_KEY || ''
  })

  if (req.method === 'GET') {
    const roles = await forge4flow.Role.listRoles()

    if (roles) {
      res.status(200).json(roles)
      return
    }
  }

  if (req.method === 'POST') {
    const user = JSON.parse(req.body)

    if (user) {
      const userOptions: CreateUserParams = {
        userId: user.id,
        email: user.email
      }
      const createdUser = await forge4flow.User.create(userOptions)

      if (createdUser) {
        res.status(200).json({ message: 'Success' })
        return
      }

      return res.status(500).json({ message: 'error' })
    }
  }

  if (req.method === 'PUT') {
    res.status(405).json({ message: 'Method Not Allowed' })
    return
  }

  res.status(500).json({ message: 'Unknown Error' })
}

export default withSessionPermission('forge4flow-admin', roles)
