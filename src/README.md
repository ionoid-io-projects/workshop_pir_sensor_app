# Introduction
Detect motion
follow this link to configure pin
https://projects.raspberrypi.org/en/projects/physical-computing/13

# How to
Compile pir.go like this
```
go get periph.io/x/periph/cmd/...
env GOOS=linux GOARCH=arm GOARM=6 go build pir-periph.go
```
Copy the generated file to your raspberry pi device and execute it with this command

```
./pir-periph
```

Happy blinking 