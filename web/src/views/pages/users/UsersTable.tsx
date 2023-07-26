// ** MUI Imports
import Paper from '@mui/material/Paper'
import Table from '@mui/material/Table'
import TableRow from '@mui/material/TableRow'
import TableHead from '@mui/material/TableHead'
import TableBody from '@mui/material/TableBody'
import TableCell from '@mui/material/TableCell'
import TableContainer from '@mui/material/TableContainer'
import IconButton from '@mui/material/IconButton'
import EditOutlinedIcon from '@mui/icons-material/EditOutlined'

// ** React Imports
import React from 'react'

// ** Next IMports
import Link from 'next/link'

// ** Date Util Import
import { convertDate } from 'src/utils/date-tools'

// ** Type Import
import { UserType } from 'src/utils/types/user'

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
            <TableCell align='right'>Created At</TableCell>
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
              <TableCell>{user.email || 'N/A'}</TableCell>
              <TableCell align='right'>{user.createdAt ? convertDate(user.createdAt) : 'N/A'}</TableCell>
              <TableCell align='right'>
                <IconButton aria-aria-label='edit'>
                  <Link className='customLink' href={`/admin/auth4flow/users/edit/${user.userId}`}>
                    <EditOutlinedIcon />
                  </Link>
                </IconButton>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  )
}

export default UsersTable
