# Hunter Douglas Platinum CLI

Simple CLI client that talks to the Hunter Douglas Platinum Gateway
to control shades.

See usage/help for features. Also see [libhdplatinum](https://github.com/vincer/libhdplatinum), the
underlying library for more information.

## Installing

You must have `go` on your system.

```
go get github.com/vincer/hdp
```

## Configuration
You will need the IP address of your HD Platinum Gateway to use this (support for autodiscovery
may come in the future). 

One option for finding it is looking at the devices connected to your router. Mine has a 3COM NIC with
a MAC starting with 00-0B-3C.

You can save your IP address in a configuration file so you don't need to enter it everytime. See usage/help.

## Usage/Help
```
hdp -h
```
