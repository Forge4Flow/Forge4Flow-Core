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
      icon: HomeOutline,
      path: '/admin'
    },
    {
      sectionTitle: 'Users & Tenants'
    },
    {
      title: 'Users',
      icon: AccountCogOutline,
      path: '/admin/users'
    },
    {
      title: 'Tenants',
      icon: AccountCogOutline,
      path: '/account-settings'
    },
    {
      sectionTitle: 'Role Based Access Control'
    },
    {
      title: 'Roles',
      icon: Login,
      path: '/admin/rbac/roles'
    },
    {
      title: 'Permissions',
      icon: AccountPlusOutline,
      path: '/admin/rbac/permissions'
    },
    {
      title: 'Check',
      icon: CreditCardOutline,
      path: '/admin/rbac/check'
    },
    {
      sectionTitle: 'Fine Grained Access Control'
    },
    {
      title: 'Object Types',
      icon: FormatLetterCase,
      path: '/typography'
    },
    {
      title: 'Objects',
      path: '/icons',
      icon: GoogleCirclesExtended
    },
    {
      title: 'Query',
      icon: CreditCardOutline,
      path: '/cards'
    },
    {
      title: 'Check',
      icon: CreditCardOutline,
      path: '/cards'
    }
  ]
}

export default navigation
