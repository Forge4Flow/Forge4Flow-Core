// ** MUI Imports
import Paper from '@mui/material/Paper'
import Table from '@mui/material/Table'
import TableRow from '@mui/material/TableRow'
import TableHead from '@mui/material/TableHead'
import TableBody from '@mui/material/TableBody'
import TableCell from '@mui/material/TableCell'
import TableContainer from '@mui/material/TableContainer'
import React, { useEffect } from 'react'

import { UserType } from 'src/pages/admin/users'

type UsersTableProps = {
  users: UserType[]
}

const UsersTable = ({ users }: UsersTableProps) => {
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label='simple table'>
        <TableHead>
          <TableRow>
            <TableCell>User ID</TableCell>
            <TableCell>Email</TableCell>
            <TableCell align='right'>Action</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {users.map(user => (
            <TableRow
              key={user.userId}
              sx={{
                '&:last-of-type td, &:last-of-type th': {
                  border: 0
                }
              }}
            >
              <TableCell component='th' scope='row'>
                {user.userId}
              </TableCell>
              <TableCell>{user.email}</TableCell>
              <TableCell align='right'></TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  )
}

export default UsersTable
