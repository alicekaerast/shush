# Shush

Use templates to silence Alert Manager. Currently behaves very similarly to `amtool import`, but with yaml. In the future this will allow for actual templating

[![Release](https://img.shields.io/github/release-pre/alicekaerast/shush.svg?logo=github&style=flat&v=1)](https://github.com/alicekaerast/shush/releases)
[![Build Status](https://img.shields.io/github/workflow/status/alicekaerast/shush/run-go-tests?logo=github&v=1)](https://github.com/alicekaerast/shush/actions)
[![Mergify Status](https://img.shields.io/endpoint.svg?url=https://gh.mergify.io/badges/alicekaerast/shush&style=flat&v=1)](https://mergify.io)
[![Go](https://img.shields.io/github/go-mod/go-version/alicekaerast/shush?v=1)](https://golang.org/)
[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/alicekaerast/shush)

## Usage

```shell
shush silence -l # to list silences
shush silence -y exampleSilence.yaml # to create a silence
shush unsilence -i 1f5247d0-3f95-452a-9da2-3ab48c00c39d # to expire a silence
```

It uses http://localhost as the alertmanager URL by default. You can override this by:

* adding `--url http://example.org`
* setting the environment variable `SHUSH_URL=http://example.net`
* writing `url: http://example.com` to one of ./.shush.yml, ~/.shush.yml, ~/.config/.shush.yml

## License

[![License](https://img.shields.io/github/license/alicekaerast/shush.svg?style=flat&v=1)](LICENSE)
