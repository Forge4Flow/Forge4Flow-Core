// ** React Imports
import { useEffect, Fragment } from 'react'

// ** Next Import
import { useRouter } from 'next/router'

// ** MUI Imports
import Chip from '@mui/material/Chip'
import Collapse from '@mui/material/Collapse'
import ListItem from '@mui/material/ListItem'
import { styled } from '@mui/material/styles'
import Typography from '@mui/material/Typography'
import Box, { BoxProps } from '@mui/material/Box'
import ListItemIcon from '@mui/material/ListItemIcon'
import ListItemButton from '@mui/material/ListItemButton'

// ** Third Party Imports
import clsx from 'clsx'

// ** Icon Imports
import Icon from 'src/@core/components/icon'

// ** Configs Import
import themeConfig from 'src/configs/themeConfig'

// ** Utils
import { hasActiveChild, removeChildren } from 'src/@core/layouts/utils'

// ** Type Import
import { NavGroup, LayoutProps } from 'src/@core/layouts/types'

// ** Custom Components Imports
import VerticalNavItems from './VerticalNavItems'

interface Props {
  item: NavGroup
  navHover: boolean
  parent?: NavGroup
  navVisible?: boolean
  groupActive: string[]
  collapsedNavWidth: number
  currentActiveGroup: string[]
  navigationBorderWidth: number
  settings: LayoutProps['settings']
  isSubToSub?: NavGroup | undefined
  saveSettings: LayoutProps['saveSettings']
  setGroupActive: (values: string[]) => void
  setCurrentActiveGroup: (items: string[]) => void
}

const MenuItemTextWrapper = styled(Box)<BoxProps>(({ theme }) => ({
  width: '100%',
  display: 'flex',
  alignItems: 'center',
  gap: theme.spacing(2),
  justifyContent: 'space-between',
  transition: 'opacity .25s ease-in-out',
  ...(themeConfig.menuTextTruncate && { overflow: 'hidden' })
}))

const VerticalNavGroup = (props: Props) => {
  // ** Props
  const {
    item,
    parent,
    navVisible,
    isSubToSub,
    groupActive,
    setGroupActive,
    currentActiveGroup,
    setCurrentActiveGroup
  } = props

  // ** Hooks & Vars
  const router = useRouter()
  const currentURL = router.asPath

  // ** Accordion menu group open toggle
  const toggleActiveGroup = (item: NavGroup, parent: NavGroup | undefined) => {
    let openGroup = groupActive

    // ** If Group is already open and clicked, close the group
    if (openGroup.includes(item.title)) {
      openGroup.splice(openGroup.indexOf(item.title), 1)

      // If clicked Group has open group children, Also remove those children to close those groups
      if (item.children) {
        removeChildren(item.children, openGroup, currentActiveGroup)
      }
    } else if (parent) {
      // ** If Group clicked is the child of an open group, first remove all the open groups under that parent
      if (parent.children) {
        removeChildren(parent.children, openGroup, currentActiveGroup)
      }

      // ** After removing all the open groups under that parent, add the clicked group to open group array
      if (!openGroup.includes(item.title)) {
        openGroup.push(item.title)
      }
    } else {
      // ** If clicked on another group that is not active or open, create openGroup array from scratch

      // ** Empty Open Group array
      openGroup = []

      // ** push Current Active Group To Open Group array
      if (currentActiveGroup.every(elem => groupActive.includes(elem))) {
        openGroup.push(...currentActiveGroup)
      }

      // ** Push current clicked group item to Open Group array
      if (!openGroup.includes(item.title)) {
        openGroup.push(item.title)
      }
    }
    setGroupActive([...openGroup])
  }

  // ** Menu Group Click
  const handleGroupClick = () => {
    toggleActiveGroup(item, parent)
  }

  useEffect(() => {
    if (hasActiveChild(item, currentURL)) {
      if (!groupActive.includes(item.title)) groupActive.push(item.title)
    } else {
      const index = groupActive.indexOf(item.title)
      if (index > -1) groupActive.splice(index, 1)
    }
    setGroupActive([...groupActive])
    setCurrentActiveGroup([...groupActive])

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [router.asPath])

  const icon = parent && !item.icon ? themeConfig.navSubItemIcon : item.icon

  const menuGroupCollapsedStyles = { opacity: 1 }

  return (
    <Fragment>
      <ListItem
        disablePadding
        className='nav-group'
        onClick={handleGroupClick}
        sx={{ mt: 1.5, px: '0 !important', flexDirection: 'column' }}
      >
        <ListItemButton
          className={clsx({
            'Mui-selected': groupActive.includes(item.title) || currentActiveGroup.includes(item.title)
          })}
          sx={{
            py: 2.25,
            width: '100%',
            borderTopRightRadius: 100,
            borderBottomRightRadius: 100,
            transition: 'padding-left .25s ease-in-out',
            '&.Mui-selected': {
              backgroundColor: 'action.hover',
              '&:hover': {
                backgroundColor: 'action.hover'
              }
            },
            '&.Mui-selected.Mui-focusVisible': {
              backgroundColor: 'action.focus',
              '&:hover': {
                backgroundColor: 'action.focus'
              }
            }
          }}
        >
          {isSubToSub ? null : (
            <ListItemIcon
              sx={{
                color: 'text.primary',
                transition: 'margin .25s ease-in-out',
                ...(parent && item.children ? { ml: 1.25, mr: 3.75 } : {})
              }}
            >
              <Icon icon={icon as string} {...(parent && { fontSize: '0.875rem' })} />
            </ListItemIcon>
          )}
          <MenuItemTextWrapper sx={{ ...menuGroupCollapsedStyles, ...(isSubToSub ? { ml: 9 } : {}) }}>
            <Typography
              {...(themeConfig.menuTextTruncate && {
                noWrap: true
              })}
            >
              {item.title}
            </Typography>
            <Box
              className='menu-item-meta'
              sx={{
                display: 'flex',
                alignItems: 'center',
                '& svg': {
                  color: 'text.primary',
                  transition: 'transform .25s ease-in-out',
                  ...(groupActive.includes(item.title) && {
                    transform: 'rotate(-90deg)'
                  })
                }
              }}
            >
              {item.badgeContent ? (
                <Chip
                  label={item.badgeContent}
                  color={item.badgeColor || 'primary'}
                  sx={{
                    mr: 1.5,
                    height: 20,
                    fontWeight: 500,
                    '& .MuiChip-label': { px: 1.5, textTransform: 'capitalize' }
                  }}
                />
              ) : null}
              <Icon icon='mdi:chevron-left' />
            </Box>
          </MenuItemTextWrapper>
        </ListItemButton>
        <Collapse
          component='ul'
          onClick={e => e.stopPropagation()}
          in={groupActive.includes(item.title)}
          sx={{
            pl: 0,
            width: '100%',
            ...menuGroupCollapsedStyles,
            transition: 'all 0.25s ease-in-out'
          }}
        >
          <VerticalNavItems
            {...props}
            parent={item}
            navVisible={navVisible}
            verticalNavItems={item.children}
            isSubToSub={parent && item.children ? item : undefined}
          />
        </Collapse>
      </ListItem>
    </Fragment>
  )
}

export default VerticalNavGroup
