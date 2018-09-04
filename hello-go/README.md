# Simple Hello World function in Golang

Make sure you have a Kubernetes endpoint running, [Kubeless](https://github.com/kubeless/kubeless) and [Serverless Framework](https://github.com/serverless) installed.

This function returns the given data.

```console
$ npm install
$ serverless deploy

$ serverless invoke -f hello-go -l --data 'hello!'
Serverless: Calling function: hello-go...
--------------------------------------------------------------------
Hello, I heard you said: hello!
$ serverless remove
```

or using Kubeles CI

```console
$ kubeless function call hello-go -d 'Hi there!'
Hello, I heard you said: Hi there!

$ kubeless function delete hello-go
```