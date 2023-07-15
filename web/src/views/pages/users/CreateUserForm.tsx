// ** React Imports
import { SyntheticEvent, useState } from 'react'

// ** MUI Imports
import Card from '@mui/material/Card'
import Grid from '@mui/material/Grid'
import Button from '@mui/material/Button'
import TextField from '@mui/material/TextField'
import CardContent from '@mui/material/CardContent'
import InputAdornment from '@mui/material/InputAdornment'
import Tab from '@mui/material/Tab'
import TabContext from '@mui/lab/TabContext'
import TabList from '@mui/lab/TabList'
import TabPanel from '@mui/lab/TabPanel'

// ** Icons Imports
import EmailOutline from 'mdi-material-ui/EmailOutline'
import AccountOutline from 'mdi-material-ui/AccountOutline'

const CreateUserForm = () => {
  const [currentTab, setCurrentTab] = useState('native')

  const handleTabChange = (_: SyntheticEvent, newValue: string) => {
    setCurrentTab(newValue)
  }
  return (
    <>
      <Card>
        <CardContent>
          <TabContext value={currentTab}>
            <TabList variant='fullWidth' onChange={handleTabChange} aria-label='new user tabs'>
              <Tab value='native' label='Flow Native' />
              <Tab value='walletless' label='Walletless' />
            </TabList>
            <TabPanel value='native'>
              <form onSubmit={e => e.preventDefault()}>
                <Grid container spacing={5}>
                  <Grid item xs={12}>
                    <TextField
                      fullWidth
                      label='Wallet Address'
                      placeholder='0xf8d6e0586b0a20c7'
                      required={true}
                      InputProps={{
                        startAdornment: (
                          <InputAdornment position='start'>
                            <AccountOutline />
                          </InputAdornment>
                        )
                      }}
                    />
                  </Grid>
                  <Grid item xs={12}>
                    <TextField
                      fullWidth
                      type='email'
                      label='Email'
                      placeholder='AwesomeUser@Forge4Flow.com'
                      helperText='optional'
                      InputProps={{
                        startAdornment: (
                          <InputAdornment position='start'>
                            <EmailOutline />
                          </InputAdornment>
                        )
                      }}
                    />
                  </Grid>
                </Grid>
              </form>
            </TabPanel>
            <TabPanel value='walletless'>
              <form onSubmit={e => e.preventDefault()}>
                <Grid container spacing={5}>
                  <Grid item xs={12}>
                    <TextField
                      fullWidth
                      type='email'
                      label='Email'
                      placeholder='AwesomeUser@Forge4Flow.com'
                      helperText='optional'
                      InputProps={{
                        startAdornment: (
                          <InputAdornment position='start'>
                            <EmailOutline />
                          </InputAdornment>
                        )
                      }}
                    />
                  </Grid>
                </Grid>
              </form>
            </TabPanel>
          </TabContext>
          <Grid item xs={12}>
            <Button type='submit' variant='contained' size='large'>
              Create User
            </Button>
          </Grid>
        </CardContent>
      </Card>
    </>
  )
}

export default CreateUserForm
