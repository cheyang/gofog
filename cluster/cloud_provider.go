package cluster

import (
	"github.com/Sirupsen/logrus"
	"github.com/cheyang/fog/cloudprovider"
	aliyun_k8s "github.com/cheyang/fog/cloudprovider/providers/aliyun/k8s"
)

func initProivder(provider, clusterType string) cloudprovider.CloudInterface {

	providerFunc := providerFuncMap[provider][clusterType]

	if providerFunc == nil {
		logrus.Infof("Not able to find provider %s for %s, ignore it...", provider, clusterType)
		return nil
	}

	return providerFunc()
}

var providerFuncMap = map[string](map[string]func() cloudprovider.CloudInterface){
	"aliyun": map[string]func() cloudprovider.CloudInterface{
		"k8s": aliyun_k8s.New,
	},
}
