# workshop_pir_sensor_app
Material for raspberry pi workshop

![circuit](doc/img/pir_wiring.png)

![circuit](doc/img/pir-sensor-pinout_1_orig.png)

# Code

## 1- Power your raspberry

You can achive it with connecting it to your pc trought the Micro USB Port of the raspberry pi

![power](doc/img/1-min.jpg)

## 2- Connect to your raspberry pi
Using putty if you're on windows, Ssh if you're on a linux based os
Follow the following instruction if you dont know how to connect to raspberry pi
[Connect to raspberry pi using Putty](https://github.com/ionoid-io-projects/workshop/blob/master/doc/od-iot-raspbian-rpi-zero-windows.md#5-first-boot)

## 3- Download Pir binary file

Assuming you're connected with... copy and past this command
If you're using Raspberry zero
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_pir_sensor_app/master/bin/arm6/pir
```

If you're using Raspberry 3 b
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_pir_sensor_app/master/bin/arm7/pir
```
## make it executable
```
chmod +x pir
```

## 4- execute binary to make led blink
```
./pir
```

## How to stop the program
To quit or stop the program click on **Ctrl+C**


https://projects.raspberrypi.org/en/projects/physical-computing/13