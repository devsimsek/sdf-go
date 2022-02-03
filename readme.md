<div id="top"></div>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/devsimsek/sdf-go">
  </a>

<h3 align="center">SDF GO</h3>

  <p align="center">
    A new way to create web applications using go and sdf framework
    <br />
    <a href="https://github.com/devsimsek/sdf-go"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/devsimsek/sdf-go">View Demo</a>
    ·
    <a href="https://github.com/devsimsek/sdf-go/issues">Report Bug</a>
    ·
    <a href="https://github.com/devsimsek/sdf-go/issues">Request Feature</a>
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
  </ol>
</details>



<!-- ABOUT THE PROJECT -->

## About The Project

Sdf is firstly created for my favorite web language, php! But in time I've looked go, c#, python (flask) and more
languages for web just in case I need them. Then I noticed the power of go. I started this project like 2 days ago
30/01/2022 and it is now published under development version v1.0.

It is not finished or not available for production!

I will work on this project. This project is now my main focus.

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

* [Gorilla Session](http://gorillatoolkit.org/pkg/sessions)
* [Built-in Go Libraries](https://go.dev/)
* [Go v1.17](https://go.dev/)
* [FDB (flat json database created by devsimsek)](https://github.com/devsimsek/project-sdf-examples/blob/main/blog/app/libraries/Fdb.php)
* devsimsek's goUtils

Most of the parts are created from scratch. Soon this framework will be independent

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->

## Getting Started

Please follow steps bellow.

### Prerequisites

This is an example of how to prepare your application to use the sdf and how to install requirements of sdf.

* Git
* Go v1.17

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/devsimsek/sdf-go.git
   ```
2. Install Required Go Packages
   ```sh
   go get
   ```
3. Enter your session secret in `.env`
   ```
   SESSION_SECRET='ENTER YOUR API';
   ```
4. Create Your Example Handler

    ```go
      package handlers
      import (
          "SDF/core"
          "fmt"
          "net/http"
      )
      
      func init() {
           core.RegisterHandle("/", homeHandler, "GET")
	     }
   
      func homeHandler(w http.ResponseWriter, r *http.Request) { 
           // Load View
		         _, err := fmt.Fprintf(w, core.LoadView("views/home.html", core.PageData{
		             PageTitle: "Home",
		             PageBody: map[string]interface{}{"version": "v1.0"},
          }))
          core.CheckError(err)
       }
    ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->

## Usage

Just open handlers directory and create a new handler :) You should be good to go :)

_For more examples, please refer to the [Documentation](https://gihtub.com/devsimsek/sdf-go/wiki)_

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->

## Roadmap

- [ ] Create Model Support (Maybe)
- [ ] Create Session Library
- [ ] Create Email Library
- [ ] Create Templating Engine (barebones)
- [ ] CSRF and DDOS protection
- [ ] Integration With sdf-php to sdf-go
- [ ] Release v1.0 Public
- [ ] Support MVC (Maybe)

See the [open issues](https://github.com/devsimsek/sdf-go/issues) for a full list of proposed features (and known
issues).

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any
contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also
simply open an issue with the tag "enhancement". Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->

## Contact

Devsimsek - [@devsimsek](https://linkedin.com/in/devsimsek) - devsimsek@outlook.com

Project Link: [https://github.com/devsimsek/sdf-go](https://github.com/devsimsek/sdf-go)

<p align="right">(<a href="#top">back to top</a>)</p>




<!-- MARKDOWN LINKS & IMAGES -->

[contributors-shield]: https://img.shields.io/github/contributors/devsimsek/sdf-go.svg?style=for-the-badge

[contributors-url]: https://github.com/devsimsek/sdf-go/graphs/contributors

[forks-shield]: https://img.shields.io/github/forks/devsimsek/sdf-go.svg?style=for-the-badge

[forks-url]: https://github.com/devsimsek/sdf-go/network/members

[stars-shield]: https://img.shields.io/github/stars/devsimsek/sdf-go.svg?style=for-the-badge

[stars-url]: https://github.com/devsimsek/sdf-go/stargazers

[issues-shield]: https://img.shields.io/github/issues/devsimsek/sdf-go.svg?style=for-the-badge

[issues-url]: https://github.com/devsimsek/sdf-go/issues

[license-shield]: https://img.shields.io/github/license/devsimsek/sdf-go.svg?style=for-the-badge

[license-url]: https://github.com/devsimsek/sdf-go/blob/master/LICENSE.txt

[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555

[linkedin-url]: https://linkedin.com/in/devsimsek
