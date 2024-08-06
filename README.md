```
  ________        _____  _____ 
 /  _____/  _____/ ____\/ ____\
/   \  ___ /  _ \   __\\   __\ 
\    \_\  (  <_> )  |   |  |   
 \______  /\____/|__|   |__|   
        \/                     
```

# Goff - Compare Folder and Files

goff is a command-line tool for compare folder and files change, written in Go.

[Video when i create goff](https://youtu.be/AY8Azk_KWNM)

## Features
- Compare Folder B vs A
    + What file new in folder B vs folder A
    + What file have deleted in folder B vs folder A
    + What file change content 
        + What content change
- Compare file B vs A

## Usage
```sh
goff [path_old] [path_new]
```

Examples: 
```sh
$ goff old-hello.txt new-hello.txt
$ goff old/hello.txt new/hello.txt
$ goff old new
$ goff old/abc new/abc
```

## Install
Prerequisites: [Install Golang](https://go.dev/doc/install) 

### Option 1: Install Binary
```sh
go install github.com/codegram01/goff@latest
```

### Option 2: Install and build from source
```sh
git clone https://github.com/codegram01/goff.git

cd goff

# Run code: 
go run . 

# Build 
go build .

# Install command 
go install 
```


