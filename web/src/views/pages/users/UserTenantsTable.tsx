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

// ** Type Import
import { TenantType } from 'src/utils/types/tenants'

type UserTenantsTableProps = {
  tenants: TenantType[]
}

const UserTenantsTable = ({ tenants }: UserTenantsTableProps) => {
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label='simple table'>
        <TableHead>
          <TableRow>
            <TableCell>Tenant ID</TableCell>
            <TableCell>Name</TableCell>
            <TableCell align='right'>Created At</TableCell>
            <TableCell align='right'>Action</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {tenants.length > 0 ? (
            tenants.map(tenant => (
              <TableRow
                key={tenant.tenantId}
                sx={{
                  '&:last-of-type td, &:last-of-type th': {
                    border: 0
                  }
                }}
              >
                <TableCell component='th' scope='row'>
                  {tenant.tenantId}
                </TableCell>
                <TableCell>{tenant.name || 'N/A'}</TableCell>
                <TableCell align='right'>{tenant.createdAt}</TableCell>
                <TableCell align='right'>
                  <IconButton aria-aria-label='edit'>
                    <Link className='customLink' href={`/admin/auth4flow/rbac/tenants/edit/${tenant.tenantId}`}>
                      <EditOutlinedIcon />
                    </Link>
                  </IconButton>
                </TableCell>
              </TableRow>
            ))
          ) : (
            <Typography variant='body2'>This user hasn't been assigned to any tenants yet.</Typography>
          )}
        </TableBody>
      </Table>
    </TableContainer>
  )
}

export default UserTenantsTable
