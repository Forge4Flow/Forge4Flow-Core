// ** Icon imports
import Login from 'mdi-material-ui/Login'
import Table from 'mdi-material-ui/Table'
import CubeOutline from 'mdi-material-ui/CubeOutline'
import HomeOutline from 'mdi-material-ui/HomeOutline'
import FormatLetterCase from 'mdi-material-ui/FormatLetterCase'
import AccountCogOutline from 'mdi-material-ui/AccountCogOutline'
import CreditCardOutline from 'mdi-material-ui/CreditCardOutline'
import AccountPlusOutline from 'mdi-material-ui/AccountPlusOutline'
import AlertCircleOutline from 'mdi-material-ui/AlertCircleOutline'
import GoogleCirclesExtended from 'mdi-material-ui/GoogleCirclesExtended'

// ** Type import
import { VerticalNavItemsType } from 'src/@core/layouts/types'

const navigation = (): VerticalNavItemsType => {
  return [
    {
      title: 'Dashboard',
      icon: 'mdi:archive-outline',
      path: '/admin'
    },
    {
      sectionTitle: 'Auth4Flow'
    },
    {
      title: 'Users & Tenants',
      children: [
        {
          title: 'Users',
          icon: 'mdi:archive-outline',
          path: '/admin/auth4flow/users'
        },
        {
          title: 'Tenants',
          icon: 'mdi:archive-outline',
          path: '#',
          badgeContent: 'API Only Currently',
          disabled: true
        }
      ]
    },
    {
      title: 'FT/NFT Gated Access Control',
      children: [
        {
          title: 'Fungible Tokens',
          icon: 'mdi:archive-outline',
          path: '/admin/auth4flow/fts'
        },
        {
          title: 'NonFungible Tokens',
          icon: 'mdi:archive-outline',
          path: '/admin/auth4flow/nfts'
        }
      ]
    },
    {
      title: 'Role Based Access Control',
      children: [
        {
          title: 'Roles',
          icon: 'mdi:archive-outline',
          path: '/admin/auth4flow/rbac/roles'
        },
        {
          title: 'Permissions',
          icon: 'mdi:archive-outline',
          path: '/admin/auth4flow/rbac/permissions'
        },
        {
          title: 'Check',
          icon: 'mdi:archive-outline',
          path: '/admin/auth4flow/rbac/check'
        }
      ]
    },
    {
      title: 'Fine Grained Access Control',
      children: [
        {
          title: 'Object Types',
          icon: 'mdi:archive-outline',
          path: '/admin/auth4flow/fgac/object-types'
        },
        {
          title: 'Objects',
          path: '/admin/auth4flow/fgac/objects',
          icon: 'mdi:archive-outline'
        },
        {
          title: 'Check',
          icon: 'mdi:archive-outline',
          path: '/admin/auth4flow/fgac/check'
        }
      ]
    },
    {
      title: 'Pricing Tiers and Features',
      children: [
        {
          title: 'Pricing Tiers',
          icon: 'mdi:archive-outline',
          path: '#',
          badgeContent: 'API Only Currently',
          disabled: true
        },
        {
          title: 'Features',
          path: '#',
          icon: 'mdi:archive-outline',
          badgeContent: 'API Only Currently',
          disabled: true
        },
        {
          title: 'Check',
          icon: 'mdi:archive-outline',
          path: '#',
          badgeContent: 'API Only Currently',
          disabled: true
        }
      ]
    },
    {
      sectionTitle: 'Alerts4Flow'
    },
    {
      title: 'Event Monitors',
      icon: 'mdi:archive-outline',
      path: '/admin/alerts4flow/monitors'
    },
    {
      title: 'Webhooks',
      icon: 'mdi:archive-outline',
      path: '#',
      badgeContent: 'Coming Soon',
      disabled: true
    }
  ]
}

export default navigation
