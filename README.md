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

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/alexandrelam/blame-detective-cli">
    <img src="public/detective.png" alt="Logo" width="120" height="120">
  </a>

<h3 align="center">B<i>lame</i> Detective CLI ✅</h3>

  <p align="center">
    Empowering Developers to Track and Expose Code Alterations! 🕵️
    <br />
    <a href="https://github.com/alexandrelam/blame-detective-cli"><strong>Explore the docs </strong></a>
    <br />
    <br />
    <a href="https://github.com/alexandrelam/blame-detective-cli/issues">Report Bug</a>
    ·
    <a href="https://github.com/alexandrelam/blame-detective-cli/issues">Request Feature</a>
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

## Prerequisites

- VSCode (or edit the file editor in `/home/alex/Documents/blame-detective-cli/cmd/blamed.go`)

<!-- GETTING STARTED -->

## Install

Linux

```sh
curl -L https://github.com/alexandrelam/blame-detective-cli/raw/main/linux-blamed -o "$HOME/.local/bin/blamed" && chmod +x "$HOME/.local/bin/blamed"
```

Mac ARM

```sh
curl -L https://github.com/alexandrelam/blame-detective-cli/raw/main/macos-blamed -o "$HOME/.local/bin/blamed" && chmod +x "$HOME/.local/bin/blamed"
```

Windows

```sh
curl no-windows-🤡
```

## Uninstall

```sh
rm "$HOME/.local/bin/blamed" && rm -rf /tmp/blamed
```

## Usage

Blame Detective offers a range of powerful features to streamline the bug tracking and debugging process. Here are some examples of how you can effectively use the application:

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Options

- -from <from_commit_hash>: Specifies the starting commit hash for viewing the log.
- -since <since_date>: Specifies the starting date for viewing the log.
- -author <author>: Specifies the author of the commits to filter.
- -to <to_commit_hash>: Specifies the ending commit hash for viewing the log (default: HEAD).
- -until <until_date>: Specifies the ending date for viewing the log (default: now).
- -editor <editor>: Specifies the text editor to use for displaying commit changes (default: code).

### Example

View commits since a specific date by a particular authors:

```bash
blamed -since "2023-01-01"
```

```bash
blamed -since "2023-01-01 -ignore \.json"
```

```bash
blamed -since "2023-01-01" -author "Alexandre Lam\|John Doe"
```

View commits between two specific commit hashes:

```bash
blamed -from abc123 -to def456
```

### Note

Either a starting commit hash (-f) or a starting date (-s) must be provided.
If an ending date (-u) is not provided, it defaults to the current date.
The script uses the git log command to generate the commit history.
The commits and their changes are displayed using the specified text editor.

## Roadmap

See the [open issues](https://github.com/alexandrelam/blame-detective-cli/issues) for a full list of proposed features (and known issues).

- [x] Generate batch of whole.diff files
- [x] Generate concurently diff folder from batched whole.diff files
- [ ] Cache query, generate folder name with a hash from the command
- [ ] Add .blamedignore config file
- [ ] Ignore big useless files
- [ ] Readd loading bar

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

Project Link: [https://github.com/alexandrelam/blame-detective-cli](https://github.com/alexandrelam/blame-detective-cli)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

- [me](https://github.com/alexandrelam/blame-detective-cli)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/alexandrelam/blame-detective-cli.svg?style=for-the-badge
[contributors-url]: https://github.com/alexandrelam/blame-detective-cli/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/alexandrelam/blame-detective-cli.svg?style=for-the-badge
[forks-url]: https://github.com/alexandrelam/blame-detective-cli/network/members
[stars-shield]: https://img.shields.io/github/stars/alexandrelam/blame-detective-cli.svg?style=for-the-badge
[stars-url]: https://github.com/alexandrelam/blame-detective-cli/stargazers
[issues-shield]: https://img.shields.io/github/issues/alexandrelam/blame-detective-cli.svg?style=for-the-badge
[issues-url]: https://github.com/alexandrelam/blame-detective-cli/issues
[license-shield]: https://img.shields.io/github/license/alexandrelam/blame-detective-cli.svg?style=for-the-badge
[license-url]: https://github.com/alexandrelam/blame-detective-cli/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/alexandre-lam-74787b191/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
