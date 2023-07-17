// ** MUI Imports
import Grid from '@mui/material/Grid'
import Card from '@mui/material/Card'
import Typography from '@mui/material/Typography'
import Modal from '@mui/material/Modal'
import Button from '@mui/material/Button'
import Box from '@mui/material/Box'

import { useEffect, useState } from 'react'

// ** Components Imports
import RolesTable from 'src/views/pages/roles/RolesTable'
import CreateRoleForm from 'src/views/pages/roles/CreateRoleForm'

// ** Type Imports
import { RoleType } from 'src/utils/types/roles'

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

const RolesPage = () => {
  const [createRoleOpen, setCreateRoleOpen] = useState(false)
  const [roles, setRoles] = useState<RoleType[]>([])
  const [fetching, setFetching] = useState(true)

  useEffect(() => {
    const fetchUsers = async () => {
      setFetching(true)
      const res = await fetch('/api/roles', {
        credentials: 'same-origin'
      })

      const json = await res.json()

      setRoles(json)
      setFetching(false)
    }

    fetchUsers()
  }, [])

  const handleCloseCreateRole = () => {
    setCreateRoleOpen(false)
  }

  return (
    <>
      <Grid container spacing={6}>
        <Grid item xs={10}>
          <Typography variant='h5'>Roles</Typography>
          <Typography variant='body2'>
            Create and manage roles that allow you to group related permissions by user persona or function.
          </Typography>
        </Grid>
        <Grid item xs={2}>
          <Button
            variant='contained'
            sx={{ px: 5.5 }}
            onClick={() => {
              setCreateRoleOpen(true)
            }}
          >
            Create Role
          </Button>
        </Grid>
        <Grid item xs={12}>
          <Card>
            {fetching ? <Typography variant='body2'>Loading roles...</Typography> : <RolesTable roles={roles} />}
          </Card>
        </Grid>
      </Grid>

      <Modal
        open={createRoleOpen}
        onClose={handleCloseCreateRole}
        aria-labelledby='modal-modal-title'
        aria-describedby='modal-modal-description'
      >
        <Box sx={style}>
          <CreateRoleForm closeHandler={handleCloseCreateRole} />
        </Box>
      </Modal>
    </>
  )
}

export default RolesPage
