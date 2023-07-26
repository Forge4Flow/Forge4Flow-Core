// ** MUI Imports
import Grid from '@mui/material/Grid'
import Typography from '@mui/material/Typography'
import Button from '@mui/material/Button'
import Divider from '@mui/material/Divider'
import Link from 'next/link'
import KeyboardReturnIcon from '@mui/icons-material/KeyboardReturn'
import Tab from '@mui/material/Tab'
import TabContext from '@mui/lab/TabContext'
import TabList from '@mui/lab/TabList'
import TabPanel from '@mui/lab/TabPanel'

// ** React Imports
import { SyntheticEvent, useState } from 'react'

// ** Next Imports
import { GetServerSideProps } from 'next/types'

// ** Forge4Flow Imports
import { Forge4FlowServer } from '@forge4flow/forge4flow-nextjs'

// ** Date Util Import
import { convertDate } from 'src/utils/date-tools'

// ** Type Imports
import { UserType } from 'src/utils/types/user'
import { TenantType } from 'src/utils/types/tenants'
import { RoleType } from 'src/utils/types/roles'

// ** Component Imports
import UserTenantsTable from 'src/views/pages/users/UserTenantsTable'
import UserRolesTable from 'src/views/pages/users/UserRolesTable'

type EditUserPageProps = {
  user: UserType
  tenants: TenantType[]
  roles: RoleType[]
}

// ** Components Imports

const EditUserPage = (props: EditUserPageProps) => {
  const [currentTab, setCurrentTab] = useState('tenant')
  const { user, tenants, roles } = props

  const handleTabChange = (_: SyntheticEvent, newValue: string) => {
    setCurrentTab(newValue)
  }

  return (
    <>
      <Grid container>
        <Grid item>
          <Link className='customLink' href='/admin/auth4flow/users'>
            <KeyboardReturnIcon fontSize='large' />
          </Link>
        </Grid>
        <Grid item xs={10}>
          <Typography variant='h5'>User: {user.userId}</Typography>
          <Typography>Email: {user.email || 'N/A'}</Typography>
          <Typography>Created: {user.createdAt}</Typography>
        </Grid>
        <Grid item xs={1}>
          <Button variant='contained' sx={{ px: 5.5 }} color='error'>
            Delete
          </Button>
        </Grid>
      </Grid>
      <Divider />
      <TabContext value={currentTab}>
        <TabList variant='fullWidth' onChange={handleTabChange} aria-label='new user tabs'>
          <Tab value='tenant' label='Tenant' />
          <Tab value='roles-perms' label='Roles & Permissions' />
          <Tab value='pricing-tiers-feat' label='Pricing Tiers & Features' />
          <Tab value='fine-grained' label='Fine Grained Access' />
        </TabList>
        <TabPanel value='tenant'>
          <UserTenantsTable tenants={tenants} />
        </TabPanel>
        <TabPanel value='roles-perms'>
          <UserRolesTable roles={roles} />
        </TabPanel>
        <TabPanel value='pricing-tiers-feat'></TabPanel>
        <TabPanel value='fine-grained'></TabPanel>
      </TabContext>
    </>
  )
}

export const getServerSideProps: GetServerSideProps<EditUserPageProps> = async context => {
  const { params } = context
  const userId = params?.userId as string

  const forge4flow = new Forge4FlowServer({
    endpoint: process.env.AUTH4FLOW_BASE_URL,
    apiKey: process.env.AUTH4FLOW_API_KEY || ''
  })

  const userObject = await forge4flow.User.get(userId)
  const user: UserType = {
    userId: userObject.userId,
    email: userObject.email,
    createdAt: userObject.createdAt ? convertDate(userObject.createdAt?.toString()) : 'N/A'
  }

  const tenantObjexts = await forge4flow.Tenant.listTenantsForUser(userId)

  const tenants: TenantType[] = tenantObjexts.map(tenant => {
    return {
      tenantId: tenant.tenantId,
      name: tenant.name,
      createdAt: tenant.createdAt ? convertDate(tenant.createdAt?.toString()) : 'N/A'
    }
  })

  const roleObjects = await forge4flow.Role.listRolesForUser(userId)

  const roles: RoleType[] = roleObjects.map(role => {
    return {
      roleId: role.roleId,
      name: role.name,
      description: role.description,
      createdAt: role.createdAt ? convertDate(role.createdAt?.toString()) : 'N/A'
    }
  })

  return {
    props: {
      user,
      tenants,
      roles
    }
  }
}

export default EditUserPage
