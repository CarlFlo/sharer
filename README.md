# Sharer

A very simple program that allows for the sharing of data over the internet using HTTP or HTTPS

## Usaged

Running the program will create a folder named **root** in the current directory.

All files in said directory will be accessible over the web.

The port **25565**, or any desired port, will have to be port forwarded for access over the web.

## Security

This program supports HTTPS and you're able to generate the needed keeps with [OpenSSL](https://www.openssl.org/) with the following two commands:

```
openssl req -x509 -newkey rsa:4096 -keyout encPrivateKey.pem -out cert.pem -days 1000
openssl rsa -in encPrivateKey.pem -out key.pem
```

Using HTTPS can be toggled within the programs code when compiling