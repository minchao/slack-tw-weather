# slack-tw-weather

[![Build Status](https://travis-ci.com/minchao/slack-tw-weather.svg?branch=master)](https://travis-ci.com/minchao/slack-tw-weather)

This is a Taiwan weather app for Slack.

Available slash commands:

- /weather
  - forecast: 36 hour weather forecasts
  - radar: Weather radar (Composite reflectivity)
  - help

## Requirements

- Golang
- Node.js >= v6.5.0
- Serverless >= 1.28.0 (You can run `yarn global add serverless` to install it)

## Quick start

### Installation

Use `go get` to download app into $GOPATH.

```bash
go get github.com/minchao/slack-tw-weather
```

Then you can change the working directory to this app.

```bash
cd $GOPATH/src/github.com/minchao/slack-tw-weather
```

Set environment variables.

```bash
export AWS_ACCESS_KEY_ID=your-key-here
export AWS_SECRET_ACCESS_KEY=your-secret-key-here
export CWB_API_KEY=your-cwb-api-key-here
```

### Deploy the service

```bash
make deploy
```

### Cleanup the service

```bash
serverless remove
```
