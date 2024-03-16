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

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]
<img src="https://github.com/alexandrelam/blame-detective/assets/25727549/79fa934c-66a9-4f62-a277-01a02fc2bd44" alt="Logo" height="35">

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/alexandrelam/blame-detective">
    <img src="public/detective.png" alt="Logo" width="120" height="120">
  </a>

<h3 align="center">B<i>lame</i> Detective CLI ✅</h3>

  <p align="center">
    Empowering Developers to Track and Expose Code Alterations! 🕵️
    <br />
    <a href="https://github.com/alexandrelam/blame-detective"><strong>Explore the docs </strong></a>
    <br />
    <br />
    <a href="https://github.com/alexandrelam/blame-detective/issues">Report Bug</a>
    ·
    <a href="https://github.com/alexandrelam/blame-detective/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

Blame Detective is a highly efficient and user-friendly application designed to streamline the process of bug tracking and debugging in software development projects. By leveraging its powerful features, Blame Detective empowers developers to identify and resolve bugs quickly and effectively, saving valuable time and resources.

In the cli!


<!-- GETTING STARTED -->

## Getting Started

Install

```sh
curl -s https://raw.githubusercontent.com/alexandrelam/blame-detective/main/install.sh | bash
```

## Usage

Blame Detective offers a range of powerful features to streamline the bug tracking and debugging process. Here are some examples of how you can effectively use the application:

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->

## Other

Open git diff from authors in vscode

```sh
git log --author="Alex\|Vincent" --pretty=format:%H | while read commit_hash; do git show "$commit_hash"; done | code -c "set ft=diff" -
```

## Roadmap


See the [open issues](https://github.com/alexandrelam/blame-detective/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Project Link: [https://github.com/alexandrelam/blame-detective](https://github.com/alexandrelam/blame-detective)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

- [me](https://github.com/alexandrelam/blame-detective)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/alexandrelam/blame-detective.svg?style=for-the-badge
[contributors-url]: https://github.com/alexandrelam/blame-detective/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/alexandrelam/blame-detective.svg?style=for-the-badge
[forks-url]: https://github.com/alexandrelam/blame-detective/network/members
[stars-shield]: https://img.shields.io/github/stars/alexandrelam/blame-detective.svg?style=for-the-badge
[stars-url]: https://github.com/alexandrelam/blame-detective/stargazers
[issues-shield]: https://img.shields.io/github/issues/alexandrelam/blame-detective.svg?style=for-the-badge
[issues-url]: https://github.com/alexandrelam/blame-detective/issues
[license-shield]: https://img.shields.io/github/license/alexandrelam/blame-detective.svg?style=for-the-badge
[license-url]: https://github.com/alexandrelam/blame-detective/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/alexandre-lam-74787b191/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
