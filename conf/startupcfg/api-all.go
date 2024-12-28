package startupcfg

func (c *ConfigAPI) All() *StartupConfig {
	if c.runConf == nil {
		return nil
	}
	return c.runConf
}

// ApiAll Database interface
func (c *ConfigAPI) ApiAll() map[string]*ServiceApiConfig {
	if c.runConf == nil {
		return nil
	}
	return c.runConf.ApiConfig
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
func (c *ConfigAPI) CustomSensitiveAll() map[string]Encrypted {
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
	newSensitive := make(map[string]Encrypted)
	for s, encrypted := range cCfg.Sensitive {
		newSensitive[s] = encrypted
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
