package commands

import "errors"

import "github.com/docker/machine/libmachine"
import "github.com/docker/machine/libmachine/persist"

var (
	errInvalidHost = errors.New("Invalid hostname")
)

func cmdUpgrade(c CommandLine, api libmachine.API) error {
	if len(c.Args()) > 2 {
		return ErrTooManyArguments
	}
	hostName, err := targetHost(c, api)
	if err != nil {
		return err
	}
	//func LoadHosts(s Store, hostNames []string) ([]*host.Host, map[string]error)
	hostSlice, errs := persist.LoadHosts(api, []string{hostName})
	if errs[hostName] != nil {
		return errs[hostName]
	}
	if len(hostSlice) == 0 {
		return errInvalidHost
	}
	return hostSlice[0].Upgrade(c.String("package"))
}
