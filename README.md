# Udid Getting Server
![GitHub last commit](https://img.shields.io/github/last-commit/totoval/udid-getting-server.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/totoval/udid-getting-server)](https://goreportcard.com/report/github.com/totoval/totoval)
![Travis (.org)](https://img.shields.io/travis/totoval/udid-getting-server.svg)
![GitHub top language](https://img.shields.io/github/languages/top/totoval/udid-getting-server.svg)
![GitHub](https://img.shields.io/github/license/totoval/udid-getting-server.svg)

## About Udid Getting Server
Udid Getting Server is a server that could be used for getting Iphone's [UDID](https://en.wikipedia.org/wiki/UDID).

## Demo
Use your iPhone open this link below:  
[https://udid.herokuapp.com/v1/udid/mobileconfig](https://udid.herokuapp.com/v1/udid/mobileconfig)

## Requirement
* https server

## How to use
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)  
1. Move `.example.env.json` to `.env.json`

```shell
mv ./.example.env.json ./env.json
```

2. Change the env config `APP_DOMAIN` in `.env.json` to your own server domain
3. Start application

```shell
go run main.go
```
  
**or**
  
```shell
go build main.go
./main
```

4. Use your Iphone open this url [https://{YOUR-SERVER-DOMAIN}/v1/udid/mobileconfig](https://{YOUR-SERVER-DOMAIN}/v1/udid/mobileconfig) using Safari
5. Follow the instructions

## Thanks
* howett.net/plist
