# Linux installation

Source:
https://golang.org/doc/install
https://golang.org/dl/


```bash
# INSTRUCTIONS FROM GOLANG.ORG
# this will error if Go is not already installed
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz
# instead only execute the tar command
tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz

# Extract to /usr/local
sudo tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz

# add Go to path
nano $HOME/.profile
vim $HOME/.profile # alternative
# add Golang to path
export PATH=$PATH:/usr/local/go/bin

# reload the configution
source $HOME/.profile

# test go command
go version
```
