// ** React Import
import { ReactNode } from 'react'

// ** Next Import
import Link from 'next/link'

// ** MUI Imports
import Box, { BoxProps } from '@mui/material/Box'
import { styled, useTheme } from '@mui/material/styles'
import Typography, { TypographyProps } from '@mui/material/Typography'

// ** Type Import
import { Settings } from 'src/@core/context/settingsContext'

// ** Configs
import themeConfig from 'src/configs/themeConfig'

interface Props {
  hidden: boolean
  settings: Settings
  toggleNavVisibility: () => void
  saveSettings: (values: Settings) => void
  verticalNavMenuBranding?: (props?: any) => ReactNode
}

// ** Styled Components
const MenuHeaderWrapper = styled(Box)<BoxProps>(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',
  paddingRight: theme.spacing(4.5),
  transition: 'padding .25s ease-in-out',
  minHeight: theme.mixins.toolbar.minHeight
}))

const HeaderTitle = styled(Typography)<TypographyProps>(({ theme }) => ({
  fontWeight: 600,
  lineHeight: 'normal',
  color: theme.palette.text.primary,
  transition: 'opacity .25s ease-in-out, margin .25s ease-in-out'
}))

const StyledLink = styled('a')({
  display: 'flex',
  alignItems: 'center',
  textDecoration: 'none'
})

const VerticalNavHeader = (props: Props) => {
  // ** Props
  const { verticalNavMenuBranding: userVerticalNavMenuBranding } = props

  // ** Hooks
  const theme = useTheme()

  return (
    <MenuHeaderWrapper className='nav-header' sx={{ pl: 6 }}>
      {userVerticalNavMenuBranding ? (
        userVerticalNavMenuBranding(props)
      ) : (
        <Link href='/' passHref>
          <StyledLink>
            <svg
              version='1.1'
              id='Layer_1'
              xmlns='http://www.w3.org/2000/svg'
              xmlnsXlink='http://www.w3.org/1999/xlink'
              x='0px'
              y='0px'
              width={30}
              height={23}
              viewBox='0 0 122.88 78.5'
              xmlSpace='preserve'
            >
              <g>
                <path
                  className='st0'
                  fill={theme.palette.primary.main}
                  d='M48.17,0.36l73.7,13.39c0.54,0.1,1,0.45,1,1V63.6c0,0.55-0.46,0.91-1,1l-73.7,13.47c-0.54,0.1-1-0.45-1-1V1.36 C47.17,0.81,47.63,0.26,48.17,0.36L48.17,0.36z'
                />
                <polygon
                  className='st0'
                  fillOpacity='0.077704'
                  fill={theme.palette.common.black}
                  points='0 8.58870968 7.25806452 12.7505183 7.25806452 16.8305646'
                />
                <polygon
                  className='st0'
                  fillOpacity='0.077704'
                  fill={theme.palette.common.black}
                  points='0 8.58870968 7.25806452 12.6445567 7.25806452 15.1370162'
                />
                <polygon
                  className='st0'
                  fillOpacity='0.077704'
                  fill={theme.palette.common.black}
                  points='22.7419355 8.58870968 30 12.7417372 30 16.9537453'
                  transform='translate(26.370968, 12.771227) scale(-1, 1) translate(-26.370968, -12.771227)'
                />
                <polygon
                  className='st0'
                  fillOpacity='0.077704'
                  fill={theme.palette.common.black}
                  points='22.7419355 8.58870968 30 12.6409734 30 15.2601969'
                  transform='translate(26.370968, 11.924453) scale(-1, 1) translate(-26.370968, -11.924453)'
                />
                <path
                  className='st0'
                  fillOpacity='0.15'
                  fill={theme.palette.common.white}
                  d='M3.04512412,1.86636639 L15,9.19354839 L15,17.1774194 L0,8.58649679 L0,3.5715689 C3.0881846e-16,2.4669994 0.8954305,1.5715689 2,1.5715689 C2.36889529,1.5715689 2.73060353,1.67359571 3.04512412,1.86636639 Z'
                />
                <path
                  className='st0'
                  fillOpacity='0.35'
                  fill={theme.palette.common.white}
                  transform='translate(22.500000, 8.588710) scale(-1, 1) translate(-22.500000, -8.588710)'
                  d='M18.0451241,1.86636639 L30,9.19354839 L30,17.1774194 L15,8.58649679 L15,3.5715689 C15,2.4669994 15.8954305,1.5715689 17,1.5715689 C17.3688953,1.5715689 17.7306035,1.67359571 18.0451241,1.86636639 Z'
                />
              </g>
            </svg>

            <HeaderTitle variant='h6' sx={{ ml: 3 }}>
              {themeConfig.templateName}
            </HeaderTitle>
          </StyledLink>
        </Link>
      )}
    </MenuHeaderWrapper>
  )
}

export default VerticalNavHeader
