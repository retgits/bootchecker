# bootchecker

![bootchecker](./bootchecker.png)

_Because you need to do things after a server booted!_

Bootchecker is a simple app that runs commands from the `config.yml` file and sends the output to a predefined email address.

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
loglevel: info               ## loglevel for the app

config:
  smtpserver: smtp.gmail.com ## The FQDN of the SMTP server
  smtpport: 587              ## The port of the SMTP server

bootmail:                    ## The array of commands that is run by `bootchecker bootmail`
  commands:
    - "ls -alh"
    - "ifconfig"

someOtherCommand:            ## The array of commands that is run by `bootchecker generic --cmd someOtherCmd`
  email: false               ## If set to true, the program will try to send an email
  commands:
    - "echo HelloWorld"
```

## Commands

### usage

```bash
Because you need to do things after a server booted

Usage:
  bootchecker [command]

Available Commands:
  bootmail    Sends an email on boot
  generic     Runs a set of commands from the configuration file
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.bootchecker.yml)
  -h, --help            help for bootchecker
      --version         version for bootchecker

Use "bootchecker [command] --help" for more information about a command.
```

### generic

```bash
$ bootchecker generic

Runs a set of commands from the configuration file

Usage:
  bootchecker generic [flags]

Flags:
      --cmd string   The name of the command group you want to execute (required)
  -h, --help         help for generic

Global Flags:
      --config string   config file (default is $HOME/.bootchecker.yml)
```

## License

See the [LICENSE](./LICENSE) file in the repository

## Icon

The amazing icon is made by [Eucalyp](https://www.flaticon.com/authors/eucalyp) from [www.flaticon.com](https://www.flaticon.com/) and is licensed by [CC 3.0 BY](http://creativecommons.org/licenses/by/3.0/)