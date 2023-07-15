// ** MUI Imports
import Grid from '@mui/material/Grid'

// ** Custom Components Imports
import CardWelcomeBack from 'src/views/pages/admin/CardWelcomeBack'

const Dashboard = () => {
  return (
    <Grid container spacing={6}>
      <Grid item xs={12}>
        <CardWelcomeBack />
      </Grid>
    </Grid>
  )
}

export default Dashboard
