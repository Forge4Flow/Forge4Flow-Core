// ** MUI Imports
import Grid from '@mui/material/Grid'
import Card from '@mui/material/Card'
import Typography from '@mui/material/Typography'
import Modal from '@mui/material/Modal'
import Button from '@mui/material/Button'
import Box from '@mui/material/Box'
import Link from 'next/link'
import KeyboardReturnIcon from '@mui/icons-material/KeyboardReturn'

// ** React Imports
import { useEffect, useState } from 'react'

// ** Next Imports
import { GetServerSideProps } from 'next/types'
import { useRouter } from 'next/router'

// ** Forge4Flow Imports
import { Forge4FlowServer } from '@forge4flow/forge4flow-nextjs'

// ** Date Util Import
import { convertDate } from 'src/utils/date-tools'

// ** Type Imports
import { UserType } from 'src/utils/types/user'

type EditUserPageProps = {
  user: UserType
}

// ** Components Imports

const EditUserPage = (props: EditUserPageProps) => {
  const { user } = props
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

  return {
    props: {
      user
    }
  }
}

export default EditUserPage
