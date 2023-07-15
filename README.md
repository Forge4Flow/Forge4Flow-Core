<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->

<a name="readme-top"></a>

<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<div align="center">

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![ELv2 License][license-shield]][license-url]

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

> Blockchain-based authentication lacks comprehensive user verification, requiring custom solutions for advanced functionalities and role-based access control. Transitioning to a Web3 environment increases the complexity of achieving secure user access control, both within DApps and when interacting with Web2 technologies. Auth4Flow offers a simple, open-source Identity and Access Management platform that simplifies Web3 authentication. It supports various authorization schemes, including RBAC, FGAC, ReBAC, and NFT/FT gated access.

### Alerts4Flow:

> One of the biggest advanges of the Flow Blockchain is it's ability to emmit events from within contracts, thus allowing developers to react to changes as they occur. Unfortuantley tooling in this area has not been widely developed. With Alerts4Flow developers can easily setup Event Monitors to receive alerts in realtime using Websockets or Webhooks.

### Ecosystem SDKs:

> [Information about Ecosystem SDKs and their purpose]

By providing these tool sets, we aim to empower developers to focus on delivering exceptional user experiences without worrying about complex authentication, access control, and other Web3 challenges.

<!-- GETTING STARTED -->

## Getting Started

To get started using Forge4Flow, follow the deployment guide to self host your own instance of Forge4Flow-Core. Once you have an instance started, follow one of our SDK quick start guides or check out or documentation for more information.

<!-- ROADMAP -->

## Roadmap

### Auth4Flow

- [x] Blockchain Native Login w/ Client & Server Sessions
- [ ] Walletless Onboarding w/ Client & Server Sessions
- [ ] NFT Gated Roles
- [ ] Token Gated Roles
- [x] GO Server SDK
- [x] JS, Node, and React SDKs
- [ ] Swift SDK
- [ ] Kotlin SDK
- [x] User Management & Admin Dashboard
- [ ] Multi-Tenant Support

### Alerts4Flow

- Predefined Event Monitors
  - FLOAT
    - [ ] FLOAT Minted
    - [ ] FLOAT Transfered
    - [ ] FLOAT Distroyed
  - Emerald ID
    - [ ] Emerald ID Created
    - [ ] Emerald ID Removed
  - Standard NFT Events
    - [ ] NFT Deposit
    - [ ] NFT Withdrawal
  - Standard FT Events
    - [ ] FT Deposit
    - [ ] FT Withdrawal
- [ ] Custom Event Monitors

### Ecosystem SDKs

- FLOAT
  - [x] Swift (iOS)
  - [ ] JS, Node
  - [ ] Go
- .find
  - [x] Swift (iOS)
  - [ ] JS, Node
  - [ ] Go
- NFT Storage
  - [x] Swift (iOS)

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
- []()
- []()

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
