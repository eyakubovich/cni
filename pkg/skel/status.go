package skel

import (
	"fmt"
	"os"
	"syscall"

	"github.com/appc/cni/pkg/ns"
	"github.com/appc/cni/pkg/types"

	"github.com/vishvananda/netlink"
)

func CmdStatus(args *CmdArgs) error {
	return ns.WithNetNSPath(args.Netns, true, func(_ *os.File) error {
		link, err := netlink.LinkByName(args.IfName)
		if err != nil {
			return fmt.Errorf("error looking up %q: %v", args.IfName, err)
		}

		addrs, err := netlink.AddrList(link, syscall.AF_INET)
		if err != nil {
			return fmt.Errorf("error retrieving IPv4 addresses for %q: %v", args.IfName, err)
		}

		if len(addrs) == 0 {
			return fmt.Errorf("%q has no IPv4 addresses", args.IfName)
		}

		r := types.StatusResult{
			IP4: *addrs[0].IPNet,
		}

		return r.Print()
	})

	return nil
}
