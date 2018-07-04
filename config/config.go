package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// ProviderConfig represents a provider's configuration.
type ProviderConfig struct {
	ClientID     string
	ClientSecret string
	Subdomain    string
	Type         string
	Username     string
}

// GetProvider returns a ProviderConfig struct containing the configuration for provider "p".
func GetProvider(p string) (*ProviderConfig, error) {
	clientSecret := viper.GetString(fmt.Sprintf("providers.%s.clientSecret", p))
	clientID := viper.GetString(fmt.Sprintf("providers.%s.clientID", p))
	subdomain := viper.GetString(fmt.Sprintf("providers.%s.subdomain", p))
	user := viper.GetString(fmt.Sprintf("providers.%s.username", p))

	if clientSecret == "" {
		return nil, errors.New("authClientSecret config value must bet set")
	}
	if clientID == "" {
		return nil, errors.New("authClientID config value must bet set")
	}
	if subdomain == "" {
		return nil, errors.New("subdomain config value must bet set")
	}

	c := ProviderConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Subdomain:    subdomain,
		Username:     user,
	}

	return &c, nil
}

// AppConfig represents an app's configuration.
type AppConfig struct {
	ID           string
	PrincipalARN string
	Provider     string
	RoleARN      string
}

// GetApp returns an AppConfig struct containing the configuration for app "app".
func GetApp(app string) (*AppConfig, error) {
	appID := viper.GetString(fmt.Sprintf("apps.%s.appId", app))
	principal := viper.GetString(fmt.Sprintf("apps.%s.principalArn", app))
	provider := viper.GetString(fmt.Sprintf("apps.%s.provider", app))
	role := viper.GetString(fmt.Sprintf("apps.%s.roleArn", app))

	if appID == "" {
		return nil, errors.New("appId config value must be set")
	}
	if principal == "" {
		return nil, errors.New("principalARN config value must be set")
	}
	if role == "" {
		return nil, errors.New("roleARN config value must be set")
	}

	c := AppConfig{
		ID:           appID,
		PrincipalARN: principal,
		Provider:     provider,
		RoleARN:      role,
	}

	return &c, nil
}