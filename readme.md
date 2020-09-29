#TCP PORT SCANNER
Software created as an introduction to Golang.

##What does it do ?
You can use it to scan a network as NMap would do.
It lets you many parameters using flags, that will be described in a following section.

##Why should I use it over NMap
You should not, trust me.

##How to use it ?

You can either compile the source code and run it on your desired platform, or if you are on MacOS you can run the executable "main".

```shell script
./main --startPort=1 --endPort=65535 --workersCount=1000 --url=scanme.nmap.org
```

##Compatibility
This software should run on MacOS, Linux and Windows.