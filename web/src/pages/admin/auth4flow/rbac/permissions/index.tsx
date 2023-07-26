// ** MUI Imports
import Grid from '@mui/material/Grid'
import Card from '@mui/material/Card'
import Typography from '@mui/material/Typography'
import Modal from '@mui/material/Modal'
import Button from '@mui/material/Button'
import Box from '@mui/material/Box'

import { useEffect, useState } from 'react'

// ** Components Imports
import UsersTable from 'src/views/pages/users/UsersTable'
import CreateUserForm from 'src/views/pages/users/CreateUserForm'

export type UserType = { userId: string; email?: string }

const style = {
  position: 'absolute' as 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 600,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24
}

const UsersPage = () => {
  const [createUserOpen, setCreateUserOpen] = useState(false)
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

  const handleCloseCreateUser = () => {
    setCreateUserOpen(false)
  }

  return (
    <>
      <Grid container spacing={6}>
        <Grid item xs={10}>
          <Typography variant='h5'>Users</Typography>
          <Typography variant='body2'>Create and manage users and what they can access in your application.</Typography>
        </Grid>
        <Grid item xs={2}>
          <Button
            variant='contained'
            sx={{ px: 5.5 }}
            onClick={() => {
              setCreateUserOpen(true)
            }}
          >
            Create User
          </Button>
        </Grid>
        <Grid item xs={12}>
          <Card>
            {users.length > 0 ? (
              <UsersTable users={users} />
            ) : (
              <Typography variant='body2'>Loading users...</Typography>
            )}
          </Card>
        </Grid>
      </Grid>

      <Modal open={createUserOpen} aria-labelledby='modal-modal-title' aria-describedby='modal-modal-description'>
        <Box sx={style}>
          <CreateUserForm closeHandler={handleCloseCreateUser} />
        </Box>
      </Modal>
    </>
  )
}

export default UsersPage
