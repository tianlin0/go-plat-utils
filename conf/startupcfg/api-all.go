package startupcfg

func (c *ConfigAPI) All() *StartupConfig {
	if c.runConf == nil {
		return nil
	}
	return c.runConf
}

// MysqlAll Database interface
func (c *ConfigAPI) MysqlAll() map[string]*MysqlConfig {
	if c.runConf == nil {
		return nil
	}
	return c.runConf.MySqlMap
}

// RedisAll Database interface
func (c *ConfigAPI) RedisAll() map[string]*RedisConfig {
	if c.runConf == nil {
		return nil
	}
	return c.runConf.RedisMap
}

// CustomSensitiveAll get all custom sensitive configs(kv)
func (c *ConfigAPI) CustomSensitiveAll() map[string]string {
	if c.runConf == nil {
		return nil
	}
	cCfg := c.runConf.CustomConfig
	if cCfg == nil {
		return nil
	}
	if cCfg.Sensitive == nil {
		return nil
	}
	newSensitive := make(map[string]string)
	for s, encrypted := range cCfg.Sensitive {
		m, err := encrypted.Get()
		if err == nil {
			newSensitive[s] = m
		}
	}
	return newSensitive
}

// CustomNormalAll get all custom normal configs(kv)
func (c *ConfigAPI) CustomNormalAll() map[string]interface{} {
	if c.runConf == nil {
		return nil
	}
	cCfg := c.runConf.CustomConfig
	if cCfg == nil {
		return nil
	}
	if cCfg.Normal == nil {
		return nil
	}
	return cCfg.Normal
}
