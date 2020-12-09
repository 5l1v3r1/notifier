# notifier

![made with go](https://img.shields.io/badge/made%20with-Go-0040ff.svg) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0040ff.svg) [![open issues](https://img.shields.io/github/issues-raw/drsigned/notifier.svg?style=flat&color=0040ff)](https://github.com/drsigned/notifier/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/drsigned/notifier.svg?style=flat&color=0040ff)](https://github.com/drsigned/notifier/issues?q=is:issue+is:closed) [![license](https://img.shields.io/badge/License-MIT-gray.svg?colorB=0040FF)](https://github.com/drsigned/notifier/blob/master/LICENSE) [![twitter](https://img.shields.io/badge/twitter-@drsigned-0040ff.svg)](https://twitter.com/drsigned)

notifier is a helper utility written for the purpose of sending notifications via webhooks to Slack.

## Resources

* [Installation](#installation)
    * [From Binary](#from-binary)
    * [From Source](#from-source)
    * [From Github](#from-github)
* [Post Install Setup](#post-install-setup)
    * [Slack](#slack)
    * [Config File](#config-file)
* [Usage](#usage)

## Installation

#### From Binary

You can download the pre-built binary for your platform from this repository's [releases](https://github.com/drsigned/notifier/releases/) page, extract, then move it to your `$PATH`and you're ready to go.

#### From Source

notifier requires **go1.14+** to install successfully. Run the following command to get the repo

```bash
$ GO111MODULE=on go get github.com/drsigned/notifier/cmd/notifier
```

#### From Github

```bash
$ git clone https://github.com/drsigned/notifier.git; cd notifier/cmd/notifier/; go build; mv notifier /usr/local/bin/; notifier -h
```

## Post Insall Setup 

#### Slack

* __step 1:__ [Get yours Slack incoming webhook URL](https://slack.com/intl/en-id/help/articles/115005265063-Incoming-webhooks-for-Slack)
* __step 2:__ Add the webhook URL to the config file.

#### Config File

The default config file should be located in `$HOME/.config/notifier/conf.yaml` and has the following contents:

```yaml
version: 1.0.0
platforms:
    slack:
        use: true
        webhook_url: "https://hooks.slack.com/services/xxxxxxxxxxx/xxxxxxxxxxx/xxxxxxxxxxx"
```

## Usage

To display help message for sigurls use the `-h` flag:

```bash
$ notifier -h
```

| Flag | Description               | Example                                 |
| :--- | :------------------------- | :-------------------------------------- |
| -l   | send message line-by-line | `cat urls.txt \| notifier -l`           |

## Contribution

[Issues](https://github.com/drsigned/notifier/issues) and [Pull Requests](https://github.com/drsigned/notifier/pulls) are welcome!