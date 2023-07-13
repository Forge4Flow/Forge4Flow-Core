// ** MUI Imports
import Grid from '@mui/material/Grid'
import Link from '@mui/material/Link'
import Card from '@mui/material/Card'
import Typography from '@mui/material/Typography'
import CardHeader from '@mui/material/CardHeader'
import { Button } from '@mui/material'

import { useEffect, useState } from 'react'

// ** Demo Components Imports
import UsersTable from 'src/views/pages/users/UsersTable'

export type UserType = { userId: string; email?: string }

const UsersPage = () => {
  const [users, setUsers] = useState<UserType[]>([])

  useEffect(() => {
    const fetchUsers = async () => {
      const res = await fetch('/api/users', {
        credentials: 'same-origin'
      })

      const json = await res.json()

      setUsers(json)
    }

    fetchUsers()
  }, [])

  return (
    <Grid container spacing={6}>
      <Grid item xs={10}>
        <Typography variant='h5'>Users</Typography>
        <Typography variant='body2'>Create and manage users and what they can access in your application.</Typography>
      </Grid>
      <Grid item xs={2}>
        <Button variant='contained' sx={{ px: 5.5 }}>
          Create User
        </Button>
      </Grid>
      <Grid item xs={12}>
        <Card>
          {users.length > 0 ? <UsersTable users={users} /> : <Typography variant='body2'>Loading users...</Typography>}
        </Card>
      </Grid>
    </Grid>
  )
}

export default UsersPage
