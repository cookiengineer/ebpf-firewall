# Usage

```bash
# Rebuild the eBPF kernel module
bash make.sh ebpf;

# Rebuild the test/main userspace program
bash make.sh source;
```

```bash
cd ./source;
sudo go run ./main.go;
```


# Tasks

All DNS filtering functionality is hidden behind the compile flag `ENABLE_DNSFILTER`
which has to be added to the `make.sh` file inside the `build_ebpf()` function.

The [BPF kernel module](./kernel/ebpf/module.c) needs the additional featureset:

- [ ] Add the `ENABLE_DNSFILTER` flag to the `build_ebpf()` method in the `make.sh`
- [ ] A method to parse DNS headers, which is defined inside [module.h](./kernel/ebpf/module.h) as the `parse_dnshdr()` method.
- [ ] A method to parse DNS questions, answers and records as separate structs
- [x] Integration with the `xdp` program to filter out all requests to DNS servers which are a blocked IP (automatically blocked inside the UDP relevant code)
- [ ] Integration with the `xdp` program to filter out all responses that contain records that point to IPs which are blocked
- [ ] Integration with the `bpf_map` to allow to filter out specific `domains` (use `domain_bans` as the BPF map name)

The code must be failsafe and allow arbitrary future DNS packets to go through undetected.

The following DNS responses must be parse-able:

- [ ] The parsing method needs support for DNS message compression
- [ ] The parsing method must not be vulnerable to DNS NAMEWRECK vulnerability (not allow buffer or stack overflows by modifying the DNS pointer field (abusing `0xc0` into a loop)
- [ ] Support for `A` records (IPv4)
- [ ] Support for `AAAA` records (IPv6)
- [ ] Support for `MX` records (domains)
- [ ] Support for `PTR` records
- [ ] Support for `SRV` records
- [ ] Support for `TXT` records (which are used for DNS exfiltration)


# License

Proprietary

