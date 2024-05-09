# linear-cli
CLI application to manage tasks in Linear.app written in GO. This CLI is not meant to be feature complete with Linear.app CLI, it is meant to have common short cuts in the CLI to help stay in focus when coding thus it is opiniated. 

# Install
There is no binary it must be compiled and installed manually

## 0: Prerequesits
Have go CLI installed

## 1: Clone Repo
``` 
git clone 
```

## 2: Build
```
go build
```

## 3: Install
```
go install
```

# Use

## Connect to graphQL
### Create API key
Settings -> My account -> API
Create a personal API key

Create an environment variable for `apiKey`
```
export apiKey=<enter API key here>
```
## Commands
### tasks
```
linear-cli tasks -p example-project -s Todo
```
`-p` Project name
`-s` Status name

### Short cut commands for status
#### backlog
Sets task as backlog with an identifier
```
linear-cli backlog ABC-01
```

#### todo
Sets task as done with an identifier
```
linear-cli todo ABC-01
```

#### wip
Sets task as Work in progress with an identifier
```
linear-cli wip ABC-01
```

#### done
Sets task as done with an identifier
```
linear-cli done ABC-01
```


