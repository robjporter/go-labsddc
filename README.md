# go-labsddc

## Objective
This is a proof of concept to show how easy it is to create a simple Software Defined Data Centre, using Cisco infrastructure components and common tools and API's.

## Outcome


## System requirements
For this POC, the system running the management platform can be one of the following;
  + A standalone server with a web browser and access to the IP addresses for the infrastructure. (This requires the source code has been compiled and transferred to this standalone server.
  + A server with a web browser, Go version 1.10 or newer and the go dependencies for this application installed.
  

## SDDC components
For this POC we will be working with the following;
  + ACI - Application Centric Infrastructure - Any version should work fine, however we will in the future use features from version 3 and newer.
  + UCS - Unified Compute Platform - Any UCS platform version will work and any version 3 or newer UCSM version.
  + Golang version 1.10 or newer
  + VMware ESXi 6.5 or newer with compatible vCenter
  + 6 x access ports on an ACI leaf
  + An L2 link, PO or VPC to an IP gateway with remote access - L3 will be added in the future.
  
  
## Setting up your GO environment
Depending on your particular environment, there are a number of ways to setup and install GO.  This repo was developed on a MAC and was installed using Brew.  For instructions on installing HomeBrew, please check [here](https://brew.sh/); and then entering;
```fish
> brew install go
```

If you do not want to use HomeBrew or you are running on a different platform, you can install the GO language using a binary from here;
https://golang.org/dl/

Once this has completed, open a cmd or terminal window and check GO has been installed and configured correctly;

Enter <b>echo $GOPATH</b>, hopefully you will be presented with a path and should be ready to go.

```fish
> echo $GOPATH
/path/to/go/bin/src/pkg folders
```

## Testing your GO environment
Once you have completed the above, its time to create a very simple test script to ensure everything is ready.

Go to a path where you are happy to store the source code for your application, this could be anywhere, including your desktop, documents, root folder, etc.

Create a folder and enter the directory.  Create a new file called "main.go" and enter the following code into it;

```go
package main

import "fmt"

func main() {
    fmt.Println("GO is working!")
}
```

At the command line, change directory using cd to the directory where your main.go file is and execute the following;
```fish
> go run main.go
```

You should see as output, something similar to;

"GO is working!"

If you reached this point, everything is working and you are ready to run the included code!

## Getting the code
There are a couple of ways you can get the code, depending on how comfortable you are with the command line and development envrionments;

You could download the zip file, [here](https://github.com/robjporter/go-labsddc/archive/master.zip).

You could use the command line git command to clone the repository to your local machine;
1. At the command line, change directory using cd to the directory where the repository will be stored.
2. Enter, git clone https://github.com/robjporter/go-labsddc.git
3. You will see output similar to the following while it is copied.
```fish
Cloning into `go-labsddc`...
remote: Counting objects: 10, done.
remote: Compressing objects 100% (8/8), done.
remove: Total 10 (delta 1), reused 10 (delta 1)
unpacking objects: 100% (10/10), done.
```
4. Change into the new directory, cd go-labsddc.
5. Move onto setting up the application.

## Application dependencies
For the application to work correctly, we need to get one dependency and we can achieve that with the following, via the cmd line.
```fish
> go get -u github.com/robjporter/go-xtools/...
> go get -u github.com/sirupsen/logrus
> go get -u github.com/kataras/iris
> go get -u github.com/jteeuwen/go-bindata
```

## Building to a Binary
One of the great advantages of GO is the ability to compile the code and all dependencies into a single binary file.  This is enhanced by building for multiple platforms.  I have included a short script to compile to most of the common formats and place them in the ./bin folder.  To run this;
```fish
> ./buildall.sh
```

## Plans for the future


**This lab has been written on and for a MAC. Other *nix platforms should be able to follow the commands exactly, however Windows users will need to amend the paths used.**
