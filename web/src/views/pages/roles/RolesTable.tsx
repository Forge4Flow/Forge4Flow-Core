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
import Typography from '@mui/material/Typography'

// ** React Imports
import React from 'react'

// ** Next IMports
import Link from 'next/link'

// ** Date Util Import
import { convertDate } from 'src/utils/date-tools'

// ** Type Import
import { RoleType } from 'src/utils/types/roles'

type RolesTableProps = {
  roles: RoleType[]
}

const UsersTable = ({ roles }: RolesTableProps) => {
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label='simple table'>
        <TableHead>
          <TableRow>
            <TableCell>Role ID</TableCell>
            <TableCell>Name</TableCell>
            <TableCell align='right'>Created At</TableCell>
            <TableCell align='right'>Action</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {roles.length > 0 ? (
            roles.map(role => (
              <TableRow
                key={role.roleId}
                sx={{
                  '&:last-of-type td, &:last-of-type th': {
                    border: 0
                  }
                }}
              >
                <TableCell component='th' scope='row'>
                  {role.roleId}
                </TableCell>
                <TableCell>{role.name || 'N/A'}</TableCell>
                <TableCell align='right'>{role.createdAt ? convertDate(role.createdAt) : 'N/A'}</TableCell>
                <TableCell align='right'>
                  <IconButton aria-aria-label='edit'>
                    <Link className='customLink' href={`/admin/auth4flow/rbac/roles/edit/${role.roleId}`}>
                      <EditOutlinedIcon />
                    </Link>
                  </IconButton>
                </TableCell>
              </TableRow>
            ))
          ) : (
            <Typography variant='body2'>You haven't created any roles.</Typography>
          )}
        </TableBody>
      </Table>
    </TableContainer>
  )
}

export default UsersTable
