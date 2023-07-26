// ** MUI Imports
import Divider from '@mui/material/Divider'
import { styled } from '@mui/material/styles'
import Typography, { TypographyProps } from '@mui/material/Typography'
import MuiListSubheader, { ListSubheaderProps } from '@mui/material/ListSubheader'

// ** Types
import { NavSectionTitle } from 'src/@core/layouts/types'
import { Settings } from 'src/@core/context/settingsContext'

// ** Custom Components Imports

interface Props {
  navHover: boolean
  settings: Settings
  item: NavSectionTitle
  collapsedNavWidth: number
  navigationBorderWidth: number
}

// ** Styled Components
const ListSubheader = styled((props: ListSubheaderProps) => <MuiListSubheader component='li' {...props} />)(
  ({ theme }) => ({
    lineHeight: 1,
    display: 'flex',
    position: 'static',
    marginTop: theme.spacing(7),
    marginBottom: theme.spacing(2),
    backgroundColor: 'transparent'
  })
)

const TypographyHeaderText = styled(Typography)<TypographyProps>(({ theme }) => ({
  fontSize: '0.75rem',
  lineHeight: 'normal',
  letterSpacing: '0.21px',
  fontWeight: theme.typography.fontWeightMedium
}))

const VerticalNavSectionTitle = (props: Props) => {
  // ** Props
  const { item } = props

  return (
    <ListSubheader
      className='nav-section-title'
      sx={{
        ...{ px: 0, py: 1.75 }
      }}
    >
      <Divider
        textAlign='left'
        sx={{
          m: '0 !important',
          lineHeight: 'normal',
          ...{
            width: '100%',
            '&:before, &:after': { top: 7, transform: 'none' },
            '& .MuiDivider-wrapper': { px: 2.5, fontSize: '0.75rem', letterSpacing: '0.21px' }
          }
        }}
      >
        <TypographyHeaderText noWrap sx={{ color: 'text.disabled' }}>
          {item.sectionTitle}
        </TypographyHeaderText>
      </Divider>
    </ListSubheader>
  )
}

export default VerticalNavSectionTitle
