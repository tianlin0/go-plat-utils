package startupcfg

import (
	"fmt"
)

// Mysql Database interface
func (c *ConfigAPI) Mysql(name string) Database {
	if c.runConf == nil {
		return nil
	}
	return c.runConf.MySQL(name)
}

// Redis Database interface
func (c *ConfigAPI) Redis(name string) Database {
	if c.runConf == nil {
		return nil
	}
	return c.runConf.Redis(name)
}

// ServiceAPI interface
func (c *ConfigAPI) ServiceAPI(service string) ServiceAPI {
	if c.runConf == nil {
		return nil
	}
	return c.runConf.ServiceAPI(service)
}

// CustomSensitive get value of sensitive custom config key (value encrypted)
func (c *ConfigAPI) CustomSensitive(key string) (string, error) {
	if c.runConf == nil {
		return "", fmt.Errorf("runConf is nil")
	}
	custom := c.runConf.Custom()
	if custom == nil {
		return "", fmt.Errorf("runConf.custom is nil")
	}

	return custom.GetSensitive(key)
}

// CustomNormal get value of insensitive custom config key
func (c *ConfigAPI) CustomNormal(key string) interface{} {
	if c.runConf == nil {
		return nil
	}
	custom := c.runConf.Custom()
	if custom == nil {
		return nil
	}
	return custom.GetNormal(key)
}
