# nDPI test pcaps

The pcap/pcapng files and their expected results (`*.pcap.out`) are copied unchanged from the
[nDPI 6.0](https://github.com/ntop/nDPI/tree/6.0) test suite (`tests/cfgs/default/pcap` and
`tests/cfgs/default/result`). nDPI is licensed under LGPL-3.0.

`TestPcapClassification` checks that the wrapper classifies every TCP/UDP flow of these pcaps as
`ndpiReader` did. They were selected among those which give the same results with the default nDPI
configuration: `ndpiReader` in nDPI tests also loads custom protocols, categories and lists from
`example/`, and has its own handling of tunnels and IP fragments.

To check the whole nDPI test suite, point `NDPI_TEST_PCAP_DIR` to `tests/cfgs/default` of a nDPI
source tree:

```
NDPI_TEST_PCAP_DIR=/path/to/nDPI/tests/cfgs/default go test ./gondpi/ -run TestNdpiTestSuite -v
```

With nDPI 6.0, 360 of the 505 pcaps with TCP/UDP flows give the same results: the others depend on
the extra configuration of `ndpiReader`.
