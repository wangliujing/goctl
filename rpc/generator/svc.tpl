package svc

import {{.imports}}

type ServiceContext struct {
	Config *config.Config
}

func NewServiceContext(c *config.Config, callBack func(context *ServiceContext)) *ServiceContext {
    sc := &ServiceContext{
        Config:c,
    }
    go callBack(sc)
	return &ServiceContext{
		Config:c,
	}
}
