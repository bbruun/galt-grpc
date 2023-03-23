# Galt - a SaltStack alternative written in Go

A sortof rewrite of SaltStack in Go using gPRC and only one port (4505/tcp).

The main reason for writing this is to get to know Golang better but also to see if the SaltStack (though very powerfull) can be sped up as it is very slow on larger installations.

The idea behind it is to

* have one binary to run it all aka server, minion and salt master cli + config files
* try to implement/extend a templating system that kinda resembels jinja, though most likely only to a certain degree
* implement user controll so the `galt` utility can be run with limitations by non-root users on the minions (managed servers) without compromising security
* have an easy way to extend using spf13/cobalt to add eg salt-cloud, switch configuration functionality etc. etc.

Why am I writing this:
I've been using BASH, Python, C/C++ since '98 when I switched to GNU/Linux and work as a systems admninstrator and over the years with the different things that I've come across then SaltStack is one of the, if not thé, most powerfull and versatile utilities that a syadmin can have. SaltStack in and of it self is in my experience superior in many wasy to eg Ansible, Puppet, CFEngine etc. They all have their forces but rely on a single basic programing language and some are agent-less and others have agents - each to their own advantage and disadvantage. SaltStack offers both, but the speed and versatility of grains on the minons (managed servers) is by and far the fastets and easiest to use and utilize. Not everything has a GUI but the design of this Galt version will be able to have a native GUI for all OS'es if need be with only gRPC as the requirement.

# Rquirements

* Golang 1.19+
* Golang Cobra
* Golang gRPC

# Development environment requirements

```
git clone github.com/bbruun/galt
cd galt
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
go get github.com/golang/protobuf/protoc-gen

```
Depending on your OS and/or distribution then do the equivalant to install the `protoc` binary:
```
sudo apt-get install protobuf-compiler
```


# Galt usage

## Run the server

`galt server [options]`

## Run the client (not implemented yet)

`galt client [options]`

## Run a galt command

`galt <command...>`
eg
`galt -m "*" cmd.run 'ls -l '`

This will just return a serialized list of static content as of this writing.


# Example usage

// WORK IN PROGRESS - NOTHING IS WORKGING YET //

Example: Run "ls -l" on all servers named "*jms*"
```
galt -m "*jms*" cmd.run "ls -l"
```
