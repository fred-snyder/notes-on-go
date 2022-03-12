# Go workspace

**Deprecated GOPATH workflow**

> This workflow has been replaced by Go Modules. This is a reference to better understand older Go documentation/books.

## Go workspace folder

```
- some folder, located anywhere (contains 3 folders)
    - bin (binary folder, where compiled code writes to)
    - pkg (package folder)
    - src
        - github.com
            - <yourUsername>
                - <repoName>
```

```bash
# name the folder whatever you want
mkdir someFolder && cd someFolder
mkdir someFolder && cd $_ # same as above
# $_: same parameter as last argument of previous command

# create Go folders
mkdir bin pkg src

# add to .bashrc to change in every shell environment
export GOPATH=$HOME/someFolder
```


## Environment Variables

GOPATH: path to the workspace folder

GOROOT: path to binary installation of Go

```bash
go env # list all go envvars
```
