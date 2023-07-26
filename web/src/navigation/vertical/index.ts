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
      title: 'NFT Gated Access Control',
      children: [
        {
          title: 'Verifiers',
          icon: 'mdi:archive-outline',
          path: '/admin/auth4flow/nfts/verifiers'
        },
        {
          title: 'Tracked NFTs',
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
          path: '/admin/auth4flow/rbac/check',
          badgeContent: 'API Only Currently',
          disabled: true
        }
      ]
    },
    {
      title: 'Fine Grained Access Control',
      children: [
        {
          title: 'Object Types',
          icon: 'mdi:archive-outline',
          path: '/admin/auth4flow/fgac/object-types',
          badgeContent: 'API Only Currently',
          disabled: true
        },
        {
          title: 'Objects',
          path: '/admin/auth4flow/fgac/objects',
          icon: 'mdi:archive-outline',
          badgeContent: 'API Only Currently',
          disabled: true
        },
        {
          title: 'Check',
          icon: 'mdi:archive-outline',
          path: '/admin/auth4flow/fgac/check',
          badgeContent: 'API Only Currently',
          disabled: true
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
    // {
    //   sectionTitle: 'Ecosystem SDKs'
    // },
    // {
    //   title: 'Swift (iOS)',
    //   children: [
    //     {
    //       title: 'FLOAT',
    //       icon: 'mdi:archive-outline',
    //       path: '/admin/sdks/float'
    //     },
    //     {
    //       title: '.find',
    //       icon: 'mdi:archive-outline',
    //       path: '/admin/sdks/find'
    //     },
    //     {
    //       title: 'NFT Catalog',
    //       icon: 'mdi:archive-outline',
    //       path: '/admin/sdks/nft-storage',
    //       badgeContent: 'Coming Soon',
    //       disabled: true
    //     }
    //   ]
    // }
  ]
}

export default navigation
