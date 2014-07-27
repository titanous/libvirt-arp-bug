# libvirt-lxc ARP issue

**Update:** The issue was that Spanning Tree Protocol is enabled by default on
`virbr0` which spends 2s in each of the listening and learning states before
enabling the interface. The solution is to disable STP with `brctl stp virbr0
off`.

This repository contains a reproduction case for ARP issue that I've been
running into with libvirt-lxc.

To run this test case, clone the repo and update the path in the `init` element
of `domain.xml` to point at the `repro` binary. Then run:

    virsh -c lxc:/// create domain.xml && virsh -c lxc:/// console arp-test

In my tests, ARP fails for ~4s. The packets appear on the veth (see
`veth0.pcap`) but it takes four seconds for an ARP to show up on the bridge (see
`virbr0.pcap`).

The test program is `containerinit.go`, and all it does is use netlink to
configure and bring up eth0 and then starts ARPing the bridge. The source code
for the netlink and arp libraries are in `Godeps/`.

I have tested and encountered this issue 100% of the time on Ubuntu 14.04 with
libvirt 1.2.2 and Linux 3.13.0-32 as well as Fedora 20 with libvirt 1.1.3.5 and
Linux 3.15.6-200.
