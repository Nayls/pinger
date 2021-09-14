<!--
Title: Pinger
Description: Web application for monitoring responses from services (http/https, ping, etc)
Author: Svyatoslav Gagarin (Nayls)
-->

# Pinger


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them

* Go v1.16 ([link install page](https://golang.org/doc/install))
* Make

### Develop

To build executable files, run the command 

```bash
make all
```

If you want to run for development use the following command 

```bash
go run cmd/pinger/pinger.go
```

## Cli usage

Use "**pinger [command] --help**" for more information about a command.

See all command in docs [**pinger commands**](docs/cli/pinger.md)

## Built With

* [go](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [viper](https://github.com/spf13/viper) - Go configuration with fangs 
* [cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions 
* [fsnotify](https://github.com/fsnotify/fsnotify) - Cross-platform file system notifications for Go. 
* [validate](https://github.com/go-playground/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving 
* [chi](https://github.com/go-chi/chi) - lightweight, idiomatic and composable router for building Go HTTP services 
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library. 

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/nayls-cloud/notify/tags).

## Authors

* **Svyatoslav Gagarin** - *Initial work* - [Nayls](https://github.com/Nayls)

See also the list of [contributors](https://github.com/nayls-cloud/notify/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
