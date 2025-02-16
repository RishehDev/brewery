# Brewery

## Create a regular Project

you should execute `brewery create project` or `brewery create project -n {name}`


## Create a web Project

you should execute `brewery create project -r regular-app` or `brewery create project -n {name} -t regular-app`

The listen port is 8080


## Create a CLI Application

You should execute `brewery create project -t cli-app` or `brewery create project -n {name} -t cli-app`

The parameter added is `first`

You can run the application executing `go run . first`
