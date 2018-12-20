# bootchecker

![bootchecker](./bootchecker.png)

_Because you need to check things after a server booted!_

Bootchecker is a simple app that runs the commands from the `config.yml` file and sends the output to a predefined email address.

## Build and Run Bootchecker

### Getting the sources

The username and password for the app are set with flags during compile time, so the best way to get the sources is by cloning the repo

```bash
git clone https://github.com/retgits/bootchecker
```

### Building the app

To build the app run `make deps` to get all the dependencies, update the variables `EMAIL_ADDRESS` and `EMAIL_PASSWORD`, and run

```bash
make build-linux
```

or

```bash
make build-macos
```

### Configuring the app

The `config.yml` has a few cofiguration values you can use

```yml
config:
  smtpserver: smtp.gmail.com ## The FQDN of the SMTP server
  smtpport: 587              ## The port of the SMTP server
commands:                    ## commands is an array of commands that are executed and the result is sent by email
  - "ls -alh"
  - "ifconfig"
```

## License

See the [LICENSE](./LICENSE) file in the repository

## Icon

The amazing icon is made by [Eucalyp](https://www.flaticon.com/authors/eucalyp) from [www.flaticon.com](https://www.flaticon.com/) and is licensed by [CC 3.0 BY](http://creativecommons.org/licenses/by/3.0/)