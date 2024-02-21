package config

type GaeConfigManager struct {
	cfg *GaeConfig
}

func NewGaeConfigManager() *GaeConfigManager {
	return &GaeConfigManager{
		cfg: ReadGaeConfig(),
	}
}

func (m *GaeConfigManager) GetDefaultProvider() *GaeConfigProvider {
	c := m.cfg.Providers[m.cfg.DefaultProvider]
	if c.Host == "" {
		return nil
	}
	return &c
}

func (m *GaeConfigManager) SetDefaultProviderToken(token string) {
	c := m.cfg.Providers[m.cfg.DefaultProvider]
	c.Token = token
	m.cfg.Providers[m.cfg.DefaultProvider] = c
	CreateOrUpdateGaeConfig(*m.cfg)
}
