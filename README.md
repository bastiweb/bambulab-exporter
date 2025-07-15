# Bambulab Exporter

A prometheus exporter for Bambu Lab 3D printers, pulling metrics directly from the printer using MQTT. This works for both "cloud" and "LAN Only" modes. Tested with the P1S and one AMS, but this likely works for others configurations as well.

A sample of the provided metrics can be found in [metrics-sample.txt](./metrics-sample.txt). (note: this sample is not guaranteed to be 100% up to date at all times)

### Setup

Things you need:

- The printer IP address/host. It's recommended to configure a static IP for your printer
- The printer [serial number](https://wiki.bambulab.com/en/general/find-sn)
- The printer Access Code. On P1S this can be found on the printer in the WLAN menu, below the current IP address

For setup with docker, grab the provided [docker-compose.sample.yml](./docker-compose.sample.yml) sample and name it `docker-compose.yml` and configure the environment variables there as described, then run `docker compose up -d` to start it

Advanced users who do not wish to use docker may build the binaries from sources and configure either with environment variables or a .env file
