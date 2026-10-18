# go-ndpi

## 1. Introduction
This project is a [nDPI](https://github.com/ntop/nDPI) wrapper by golang based on CGO. You can use nDPI by golang to do packet inspection.

It is based on ndpi 6.0 release currently.

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

Packages from apt may lag behind the latest nDPI release. In that case build nDPI 6.0 from source:
```
git clone --branch 6.0 https://github.com/ntop/nDPI.git
cd nDPI && ./autogen.sh && ./configure --with-maxminddb && make && make install
```

### 2.2 Use this lib
There is a simple usage in in folder `example` which could be a good start.
`example01` reads packets from a live interface or from a pcap file:
```
go run ./example01 -i eth0
go run ./example01 -f capture.pcap
```

Since nDPI 4.10 the detection module has a two-step initialization: configuration and custom files
must be applied before the module is finalized.
```go
bitmask := gondpi.NdpiProtocolBitmaskSetAll(gondpi.NewNdpiProtocolBitmask())
dm, err := gondpi.NewNdpiDetectionModule(types.NDPI_LICENSE_NOT_FOR_PROFIT_LGPL, bitmask)
// optional, before Finalize: see doc/configuration.md in nDPI
err = dm.SetConfig("", "dpi.guess_on_giveup", "0")
err = dm.LoadProtocolsFile("protos.txt")
err = dm.Finalize()
```
`NdpiDetectionModuleInitialize(license, bitmask)` creates and finalizes the module in one call.

### 2.3 Upgrade from nDPI 4.4.0
- `NdpiDetectionModuleInitialize` takes a `types.NdpiLicenseType` instead of `ndpi_init_prefs` (removed upstream, `types/ndpi_prefs.go` is gone)
- `SetLogLevel` and `SetDebugBitmask` must be called before `Finalize`
- `DetectionGiveup(flow)` drops the `enableGuess` argument: use config `dpi.guess_on_giveup`
- `NdpiProto` has new `Breed` and `State` fields: a flow is done when `State == types.NDPI_STATE_CLASSIFIED` and `IsExtraDissectionPossible` is false
- `LoadMaliciousJa3File` is replaced by `LoadMaliciousJa4File`; `Load*File` functions return an error too
- Flow getters follow the new `ndpi_flow_struct`: `GetSrc/DstAddr/Port` become `GetClient/ServerAddr/Port`, `GetGuessedHostProtocolId` becomes `GetGuessedProtocolIdByIp`, `GetStun`, `GetFlowExtraInfo`, `GetNextTcpSeqNR` and the removed bit flags are gone
- `NdpiProtoDefaults` ports are `[]NdpiPortRange` and `ProtoIdx` is now `DissectorIdx`
- Enums are regenerated from nDPI 6.0 headers: protocol breed and confidence values changed, `NDPI_PROTOCOL_CATEGORY_SOCIAL_NETWOR` is now `NDPI_PROTOCOL_CATEGORY_SOCIAL_NETWORK` and `CUSTOM_CATEGORY_*` are now `NDPI_PROTOCOL_CATEGORY_*`

### 2.4 More information of compile and link in CGO
You can see follow codes in `ndpi.go`
- Linux lib path is `/usr/include/`, you can adapt it according to your ENV
- Lib ndpi is static link while others are dynamic link
  - The wrapper accesses nDPI structs directly, so linking it statically guarantees that headers and library match
  - Maxminddb is required by `libndpi` built with `--with-maxminddb` (as the apt one). It is dynamic linked, since some
    distributions (e.g. Arch/Manjaro) do not ship `libmaxminddb.a`: `libmaxminddb.so` is needed at runtime

```
#cgo CFLAGS: -I/usr/include/
#cgo LDFLAGS: -Wl,-Bstatic -lndpi -Wl,-Bdynamic -lmaxminddb -lpcap -lm -ldl -pthread
```

### 2.5 Tests
Tests classify the pcaps of nDPI test suite in `gondpi/testdata/ndpi` and compare every TCP/UDP flow with
the results of `ndpiReader`:
```
go test ./...
```
See `gondpi/testdata/ndpi/README.md` to run them against the whole nDPI test suite.
