<a name="readme-top"></a>

<div align="center">

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![ELv2 License][license-shield]][license-url]
[![CodeFactor][code-factor-shield]][code-factor-url]
[![WakaTime][wakatime-shield]][wakatime-url]

<!-- PROJECT LOGO -->
<br />
  <img src="images/logo.png" alt="Logo" width="160" height="160">

  <h3>Forge4Flow</h3>

  <p>
    <a href="http://forge4flow.gitbook.io/docs"><strong>Explore the docs »</strong></a>
    <br />
    <a href="https://github.com/Forge4Flow/Forge4Flow-Core/issues">Report Bug</a>
    ·
    <a href="https://github.com/Forge4Flow/Forge4Flow-Core/issues">Request Feature</a>
  </p>

  <h3>Forge4Flow provides ecosystem and developer tools for the Flow Blockchain, including Identity and Access Management, Blockchain Event Monitoring, and SDKs to better integrate dApps with the ecosystem.</h3>

</div>

## The problem it solves

Developers face challenges in creating dApps due to issues like user authentication, access control, system monitoring, and third-party integration. Forge4Flow aims to address these problems with a developer infrastructure and tooling platform. We're excited to introduce three tool sets to advance the ecosystem:

### Auth4Flow:

> Blockchain-based authentication lacks comprehensive user verification, requiring custom solutions for advanced functionalities and role-based access control. Transitioning to a Web3 environment increases the complexity of achieving secure user access control, both within DApps and when interacting with Web2 technologies. Auth4Flow offers a simple, open-source Identity and Access Management platform that simplifies Web3 authentication. It supports various authorization schemes, including RBAC, FGAC, ReBAC, and NFT/FT/Event gated access.

### Alerts4Flow:

> One of the biggest advantages of the Flow Blockchain is its ability to emmit events from within contracts, thus allowing developers to react to changes as they occur. Unfortunate tooling in this area has not been widely developed. With Alerts4Flow developers can easily set up Event Monitors to receive alerts in realtime using Websockets or Webhooks.

### Ecosystem SDKs:

> Lack of mobile resources is a huge factor for their being very little Web3 Mobile apps. By releases ecosystem SDKs for multiple platforms we can lower the barrier to entry for new developers. We have scoped several SDKs to target for Swift (iOS).

By providing these tool sets, we aim to empower developers to focus on delivering exceptional user experiences without worrying about complex authentication, access control, and other Web3 challenges.

<!-- GETTING STARTED -->

## Getting Started

To get started using Forge4Flow, follow the deployment guide to self-host your own instance of Forge4Flow-Core. Once you have an instance started, follow one of our SDK quick start guides or check out or documentation for more information.

<!-- ROADMAP -->

## Roadmap To v0.1.0

### Forge Cloud

Once v0.1.0 is available the plan is to launch a service for hosting multiple Forge4Flow-Core environments with central management dashboard similar to Forge Manager

### Forge Manager

- [ ] Admin Dashboard
- [x] Verify/Install Docker Installation
- [ ] Launch Docker Images For Manager
- [ ] Launch Docker Images For Forge4Flow-Core Environments
- [ ] Manage SSL Certs
- [ ] Proxy Traffic Through Manager To Environments

### Forge4Flow-Core

#### Auth4Flow

- [x] Blockchain Native Login w/ Client & Server Sessions
- [ ] Walletless Onboarding w/ Client & Server Sessions
  - [ ] Transaction Signing APIs
  - [ ] Parent/Child Account Linking
  - [ ] Forced Hybrid Authentication (Creating Flow Child Accounts for Blockchain Native Accounts)
- [x] Blockchain FT/NFT/Event Gated Access Control
- [ ] .find Name and Profile Integration
- [ ] GO Server SDK
- [x] [JS SDK](https://github.com/Forge4Flow/Forge4Flow-JS)
- [x] [Node SDK](https://github.com/Forge4Flow/Forge4Flow-Node)
- [x] [React SDK](https://github.com/Forge4Flow/Forge4Flow-React)
- [x] [Next.js](https://github.com/Forge4Flow/Forge4Flow-NextJS)
- [x] [Swift SDK](https://github.com/Forge4Flow/Forge4Flow-Swift)
- [ ] Kotlin SDK
- [x] Multi-Tenant Support

#### Alerts4Flow

- [x] Custom Event Monitors
- [x] Websocket Support
- [ ] API Route To Configure Webhook Subscriptions To Events

### Ecosystem SDKs

- Flow Ecosystem
  - FLOAT
    - [x] [Swift (iOS)](https://github.com/Forge4Flow/FLOAT-Swift-SDK)
    - [ ] JS, Node
    - [ ] Go
  - .find
    - [x] [Swift (iOS)](https://github.com/Forge4Flow/FIND-Swift-SDK)
    - [ ] JS, Node
    - [ ] Go
  - Flow NFT Catalog
    - [ ] Swift (iOS)
    - [ ] Go
- Mobile Platforms
  - Swift (iOS)
    - [x] NFT.storage
    - [x] SwiftIPFS-Image
- General Purpose
  - [x] [GCP KMS authorizer (signer)](https://github.com/Forge4Flow/GCP-KMS-Flow-Authorizer)

See the [open issues](https://github.com/Forge4Flow/Forge4Flow-Core/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated** Please see `CONTRIBUTING.md` for more information.

<!-- LICENSE -->

## License

Forge4Flow is distributed under the ELv2 License. See `LICENSE` for more information. Our SDKs are distributed under MIT licenses, see each SDKs GitHub repo for more information.

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

- [warrant.dev](https://github.com/warrant-dev/warrant) Forge4Flow was built on top an amazing open source authorization project called Warrant. All credit for the Web2 authorization functionality within Forge4Flow goes to the amazing warrant.dev team.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/Forge4Flow/Forge4Flow-Core.svg?style=for-the-badge
[contributors-url]: https://github.com/Forge4Flow/Forge4Flow-Core/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Forge4Flow/Forge4Flow-Core.svg?style=for-the-badge
[forks-url]: https://github.com/Forge4Flow/Forge4Flow-Core/network/members
[stars-shield]: https://img.shields.io/github/stars/Forge4Flow/Forge4Flow-Core.svg?style=for-the-badge
[stars-url]: https://github.com/Forge4Flow/Forge4Flow-Core/stargazers
[issues-shield]: https://img.shields.io/github/issues/Forge4Flow/Forge4Flow-Core.svg?style=for-the-badge
[issues-url]: https://github.com/Forge4Flow/Forge4Flow-Core/issues
[license-shield]: https://img.shields.io/badge/license-elv2-blue?style=for-the-badge
[license-url]: https://github.com/Forge4Flow/Forge4Flow-Core/blob/master/LICENSE
[code-factor-shield]: https://img.shields.io/codefactor/grade/github/forge4flow/forge4flow-core/main?style=for-the-badge
[code-factor-url]: https://www.codefactor.io/repository/github/forge4flow/forge4flow-core
[wakatime-shield]: https://wakatime.com/badge/user/0a8af699-5f37-4933-8df8-b7282a2ab48c/project/ec5d6c75-4dc2-4d9c-a344-78e9fcdf151d.svg?style=for-the-badge
[wakatime-url]: https://wakatime.com/badge/user/0a8af699-5f37-4933-8df8-b7282a2ab48c/project/ec5d6c75-4dc2-4d9c-a344-78e9fcdf151d
