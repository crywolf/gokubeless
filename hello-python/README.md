# Simple Hello World function in Python

Make sure you have a Kubernetes endpoint running, [Kubeless](https://github.com/kubeless/kubeless) and [Serverless Framework](https://github.com/serverless) installed.

This function returns some data.

```console
$ npm install
$ serverless deploy

$ serverless invoke -f hello-python  -l
Serverless: Calling function: hello-python...
--------------------------------------------------------------------
HELLO world
$ serverless remove
```

or using Kubeles CI

```console
$ kubeless function call hello-python
HELLO world

$ kubeless function delete hello-python
```