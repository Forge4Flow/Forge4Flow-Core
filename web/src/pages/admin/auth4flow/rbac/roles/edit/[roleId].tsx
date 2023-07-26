// ** MUI Imports
import Grid from '@mui/material/Grid'
import Typography from '@mui/material/Typography'
import Button from '@mui/material/Button'
import Link from 'next/link'
import KeyboardReturnIcon from '@mui/icons-material/KeyboardReturn'

// ** Next Imports
import { GetServerSideProps } from 'next/types'

// ** Forge4Flow Imports
import { Forge4FlowServer } from '@forge4flow/forge4flow-nextjs'

// ** Date Util Import
import { convertDate } from 'src/utils/date-tools'

// ** Type Import
import { RoleType } from 'src/utils/types/roles'

type EditRolePageProps = {
  role: RoleType
}

// ** Components Imports

const EditUserPage = (props: EditRolePageProps) => {
  const { role } = props

  return (
    <>
      <Grid container>
        <Grid item>
          <Link className='customLink' href='/admin/auth4flow/users'>
            <KeyboardReturnIcon fontSize='large' />
          </Link>
        </Grid>
        <Grid item xs={10}>
          <Typography variant='h5'>Role: {role.roleId}</Typography>
          <Typography>Name: {role.name || 'N/A'}</Typography>
          <Typography>Created: {role.createdAt}</Typography>
        </Grid>
        <Grid item xs={1}>
          <Button variant='contained' sx={{ px: 5.5 }} color='error'>
            Delete
          </Button>
        </Grid>
      </Grid>
    </>
  )
}

export const getServerSideProps: GetServerSideProps<EditRolePageProps> = async context => {
  const { params } = context
  const roleId = params?.roleId as string

  const forge4flow = new Forge4FlowServer({
    endpoint: process.env.AUTH4FLOW_BASE_URL,
    apiKey: process.env.AUTH4FLOW_API_KEY || ''
  })

  const roleObject = await forge4flow.Role.get(roleId)
  const role: RoleType = {
    roleId: roleObject.roleId,
    name: roleObject.name,
    createdAt: roleObject.createdAt ? convertDate(roleObject.createdAt?.toString()) : 'N/A'
  }

  return {
    props: {
      role: role
    }
  }
}

export default EditUserPage
