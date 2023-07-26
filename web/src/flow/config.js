import { config } from '@onflow/fcl'

config({
  'app.detail.title': 'Forge4Flow Dashboard', // this adds a custom name to our wallet
  'app.detail.icon':
    'https://github.com/Forge4Flow/.github/blob/ceeaf8b56d8aec6a939b5d9fc17b2bf2db4e7c23/images/logo.png', // this adds a custom image to our wallet
  'accessNode.api': process.env.NEXT_PUBLIC_ACCESS_NODE, // this is for the local emulator
  'discovery.wallet': process.env.NEXT_PUBLIC_WALLET // this is for the local dev wallet
})
