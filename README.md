# go-ndpi

## 1. Introduction
This project is a [nDPI](https://github.com/ntop/nDPI) wrapper by golang based on CGO. You can use nDPI by golang to do packet inspection.

## 2. How to use this lib
### 2.1 Install nDPI Lib
Clang lib nDPI should be installed firstly. Ubuntu 22.04 is used in follow guide.

- Install Package Repository for Ubuntu 22.04 LTS
```bash
apt-get install software-properties-common wget
add-apt-repository universe
wget https://packages.ntop.org/apt-stable/22.04/all/apt-ntop-stable.deb
apt install ./apt-ntop-stable.deb
```

- Install nDPI
```
apt-get install ndpi-dev
ldconfig -p | grep -i ndpi
find /usr/lib/ | grep -i ndpi
find /usr/include/ | grep -i ndpi
```

- Install required libs
  - Libpcap is used for packet capture
  - Maxminddb is enabled in 'libndpi' from apt-get. If you are using 'libndpi' which is installed by apt, 'libmaxminddb' is also required
```
apt-get install libpcap-dev libmaxminddb-dev
ldconfig -p | grep -i maxmind
find /usr/lib/ | grep -i maxmind
find /usr/include/ | grep -i maxmind
```

### 2.2 Use this lib
There is a simple usage in in folder `example` which could be a good start.

### 2.3 More information of compile and link in CGO
You can see follow codes in `ndpi.go`
- Linux lib path is `/usr/include/`, you can adapt it according to your ENV
- Lib ndpi is static link while others are dynamic link

```
#cgo CFLAGS: -I/usr/include/
#cgo LDFLAGS: -Wl,-Bstatic -lndpi -lmaxminddb -Wl,-Bdynamic -lpcap -lm -pthread
```
