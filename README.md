# passphrase-web

Simple password generator with web interface. [Vue.js](https://vuejs.org/) in front-end, [golang](https://golang.org/) in back-end.

![passphrase-web](https://github.com/sgrzywna/passphrase-web/blob/master/screenshot.jpg)

[Check it out!](https://passphrase.sgnet.xyz)

From time to time I need to generate password, but almost any web generator out there is over complicated or gives password maybe hard to crack but also hard to remember.

I wanted small, fast password generator and, what is most important - generated passwords must be easy to remember, but hard to crack. This little project tries to fulfill these requirements.

To build docker image:

```bash
build/build.sh
```

Docker image is versioned with the latest tag from repository.

To actually run application inside docker container listening at port 8080:

```bash
docker run -it --rm -p 8080:8080 passphrase-web:1.0.3
```

There is also passphrase version for those fallen in love with [CLI](https://github.com/sgrzywna/passphrase).
