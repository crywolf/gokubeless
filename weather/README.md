# Weather report

This is a simple function that uses query.yahooapis.com to retrieve the weather for a specific location.

Make sure you have a Kubernetes endpoint running, [Kubeless](https://github.com/kubeless/kubeless) and [Serverless Framework](https://github.com/serverless) installed.

## Deploy the function

1/ manually with Kubeless

```
kubeless function deploy weather --runtime go1.10 \
  --from-file weather.go \
  --handler weather.Handler
```

2/ or with Kubeless Serverless Plugin

```console
$ npm install
$ serverless deploy
```

## Execute the function

```
$ kubeless function call weather --data 'Prague'
It is 24 °C in Prague and mostly cloudy
```

## Delete the function

```
$ kubeless function delete weather
or
$ serverless remove
```
