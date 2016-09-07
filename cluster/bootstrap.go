package cluster

import (
	"github.com/Sirupsen/logrus"
	provider_registry "github.com/cheyang/fog/cloudprovider/registry"
	"github.com/cheyang/fog/cluster/ansible"
	"github.com/cheyang/fog/cluster/deploy"
	"github.com/cheyang/fog/types"
	"github.com/cheyang/fog/util/dump"
)

func Bootstrap(spec types.Spec) error {

	err := types.Validate(spec)
	if err != nil {
		return err
	}

	logrus.Infof("spec: %+v", spec)

	//register core dump tool
	dump.InstallCoreDumpGenerator()

	hosts, err := provisionVMs(spec)

	if err != nil {
		return err
	}

	cp := provider_registry.GetProvider(spec.CloudDriverName, spec.ClusterType)
	if cp != nil {
		cp.SetHosts(hosts)
		cp.Configure() // configure IaaS
	}

	var deployer deploy.Deployer
	deployer, err = ansible.NewDeployer(spec.Name)
	if err != nil {
		return err
	}
	deployer.SetHosts(hosts)
	if len(spec.Run) > 0 {
		deployer.SetCommander(spec.Run)
	} else {
		deployer.SetCommander(spec.DockerRun)
	}

	return deployer.Run()
}
