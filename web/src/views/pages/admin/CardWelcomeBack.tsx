// ** MUI Imports
import Card from '@mui/material/Card'
import { styled } from '@mui/material/styles'
import Typography from '@mui/material/Typography'
import CardContent from '@mui/material/CardContent'
import Grid, { GridProps } from '@mui/material/Grid'
import Button from '@mui/material/Button'
import Link from 'next/link'

// Styled Grid component
const StyledGrid = styled(Grid)<GridProps>(({ theme }) => ({
  [theme.breakpoints.down('sm')]: {
    display: 'flex',
    justifyContent: 'center'
  }
}))

// Styled component for the image
const Img = styled('img')(({ theme }) => ({
  right: 13,
  bottom: 0,
  height: 200,
  position: 'absolute',
  [theme.breakpoints.down('sm')]: {
    height: 165,
    position: 'static'
  }
}))

const CardWelcomeBack = () => {
  return (
    <Card sx={{ position: 'relative', overflow: 'visible', mt: { xs: 0, sm: 14.4, md: 0 } }}>
      <CardContent sx={{ p: theme => theme.spacing(7.25, 7.5, 7.75, 7.5) }}>
        <Grid container spacing={6}>
          <Grid item xs={12} sm={4}>
            <Typography variant='h5' sx={{ mb: 6.5 }}>
              Welcome Back! 🥳
            </Typography>
            <Typography variant='body2'>What would you like to see on this dashboard page?</Typography>
            <Typography variant='body2'>Join our Discord server to voice your opinion</Typography>
          </Grid>
          <Grid item xs={12} sm={2}>
            <Link href='https://discord.gg/S85mDy2qxE' className='customLink' target='_blank'>
              <Button variant='contained' sx={{ px: 5.5, position: 'absolute', bottom: 25 }}>
                Join Discord
              </Button>
            </Link>
          </Grid>
          <StyledGrid item xs={12} sm={6}>
            <Img alt='Welcome back John' src='/images/cards/illustration-john.png' />
          </StyledGrid>
        </Grid>
      </CardContent>
    </Card>
  )
}

export default CardWelcomeBack
