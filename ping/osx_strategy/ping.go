package osx_strategy

import (
	"os/exec"
	"strconv"
)

func Ping(host string, timeout int) error {
	cmd, args := getOSXPingCommand(host, timeout)

	_, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return err
	}
	return nil
}

// Ping util on Mac OSX:
//   ping -o -n -c 1 -t 1 ip_or_host
//   -o = exit on first received packet
//   -n = numeric only (don't resolve hostnames of ip addrs)
//   -c 1 = only send one ping
//   -t S = only wait S seconds for reply before exiting
func getOSXPingCommand(host string, timeoutSeconds int) (string, []string) {
	cmd := "/sbin/ping"
	args := []string{"-o", "-n", "-c", "1", "-t", strconv.Itoa(timeoutSeconds), host}
	return cmd, args
}
