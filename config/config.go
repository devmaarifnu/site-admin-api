package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	App        AppConfig        `mapstructure:"app"`
	Database   DatabaseConfig   `mapstructure:"database"`
	JWT        JWTConfig        `mapstructure:"jwt"`
	CDN        CDNConfig        `mapstructure:"cdn"`
	CORS       CORSConfig       `mapstructure:"cors"`
	RateLimit  RateLimitConfig  `mapstructure:"rate_limit"`
	Upload     UploadConfig     `mapstructure:"upload"`
	Logging    LoggingConfig    `mapstructure:"logging"`
	Email      EmailConfig      `mapstructure:"email"`
	Frontend   FrontendConfig   `mapstructure:"frontend"`
	Security   SecurityConfig   `mapstructure:"security"`
	Pagination PaginationConfig `mapstructure:"pagination"`
}

// AppConfig contains application settings
type AppConfig struct {
	Name       string `mapstructure:"name"`
	Env        string `mapstructure:"env"`
	Port       int    `mapstructure:"port"`
	APIVersion string `mapstructure:"api_version"`
}

// DatabaseConfig contains database connection settings
type DatabaseConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Name            string `mapstructure:"name"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

// JWTConfig contains JWT settings
type JWTConfig struct {
	Secret              string `mapstructure:"secret"`
	ExpiresHours        int    `mapstructure:"expires_hours"`
	RefreshSecret       string `mapstructure:"refresh_secret"`
	RefreshExpiresHours int    `mapstructure:"refresh_expires_hours"`
}

// CDNConfig contains CDN file server settings
type CDNConfig struct {
	BaseURL     string `mapstructure:"base_url"`
	BaseURLProd string `mapstructure:"base_url_prod"`
	Token       string `mapstructure:"token"`
	Timeout     int    `mapstructure:"timeout"`
}

// CORSConfig contains CORS settings
type CORSConfig struct {
	AllowedOrigins   []string `mapstructure:"allowed_origins"`
	AllowedMethods   []string `mapstructure:"allowed_methods"`
	AllowedHeaders   []string `mapstructure:"allowed_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	MaxAge           int      `mapstructure:"max_age"`
}

// RateLimitConfig contains rate limiting settings
type RateLimitConfig struct {
	Enabled              bool `mapstructure:"enabled"`
	RequestsPerMinute    int  `mapstructure:"requests_per_minute"`
	LoginMaxAttempts     int  `mapstructure:"login_max_attempts"`
	LoginWindowMinutes   int  `mapstructure:"login_window_minutes"`
}

// UploadConfig contains file upload settings
type UploadConfig struct {
	MaxSizeImageMB      int      `mapstructure:"max_size_image_mb"`
	MaxSizeDocumentMB   int      `mapstructure:"max_size_document_mb"`
	AllowedImageTypes   []string `mapstructure:"allowed_image_types"`
	AllowedDocumentTypes []string `mapstructure:"allowed_document_types"`
}

// LoggingConfig contains logging settings
type LoggingConfig struct {
	Level  string            `mapstructure:"level"`
	Format string            `mapstructure:"format"`
	Output string            `mapstructure:"output"`
	File   LogFileConfig     `mapstructure:"file"`
}

// LogFileConfig contains log file settings
type LogFileConfig struct {
	Path        string `mapstructure:"path"`
	MaxSizeMB   int    `mapstructure:"max_size_mb"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAgeDays  int    `mapstructure:"max_age_days"`
	Compress    bool   `mapstructure:"compress"`
}

// EmailConfig contains email settings
type EmailConfig struct {
	SMTPHost     string `mapstructure:"smtp_host"`
	SMTPPort     int    `mapstructure:"smtp_port"`
	SMTPUser     string `mapstructure:"smtp_user"`
	SMTPPassword string `mapstructure:"smtp_password"`
	FromAddress  string `mapstructure:"from_address"`
	FromName     string `mapstructure:"from_name"`
}

// FrontendConfig contains frontend settings
type FrontendConfig struct {
	URL               string `mapstructure:"url"`
	PasswordResetPath string `mapstructure:"password_reset_path"`
}

// SecurityConfig contains security settings
type SecurityConfig struct {
	PasswordMinLength      int  `mapstructure:"password_min_length"`
	PasswordRequireUpper   bool `mapstructure:"password_require_upper"`
	PasswordRequireLower   bool `mapstructure:"password_require_lower"`
	PasswordRequireNumber  bool `mapstructure:"password_require_number"`
	PasswordRequireSpecial bool `mapstructure:"password_require_special"`
	SessionTimeoutMinutes  int  `mapstructure:"session_timeout_minutes"`
}

// PaginationConfig contains pagination defaults
type PaginationConfig struct {
	DefaultPage  int `mapstructure:"default_page"`
	DefaultLimit int `mapstructure:"default_limit"`
	MaxLimit     int `mapstructure:"max_limit"`
}

// LoadConfig loads configuration from config.yaml file
func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/site-admin-api/")

	// Enable environment variable override
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}

	// Validate required fields
	if err := validateConfig(&config); err != nil {
		log.Fatalf("Config validation failed: %v", err)
	}

	return &config
}

// validateConfig validates required configuration fields
func validateConfig(cfg *Config) error {
	if cfg.App.Name == "" {
		return fmt.Errorf("app.name is required")
	}
	if cfg.App.Port == 0 {
		return fmt.Errorf("app.port is required")
	}
	if cfg.Database.Host == "" {
		return fmt.Errorf("database.host is required")
	}
	if cfg.Database.Name == "" {
		return fmt.Errorf("database.name is required")
	}
	if cfg.JWT.Secret == "" {
		return fmt.Errorf("jwt.secret is required")
	}
	if cfg.CDN.Token == "" {
		return fmt.Errorf("cdn.token is required")
	}
	return nil
}

// GetCDNBaseURL returns the appropriate CDN base URL based on environment
func (c *Config) GetCDNBaseURL() string {
	if c.App.Env == "production" {
		return c.CDN.BaseURLProd
	}
	return c.CDN.BaseURL
}

// IsProduction returns true if running in production environment
func (c *Config) IsProduction() bool {
	return c.App.Env == "production"
}

// IsDevelopment returns true if running in development environment
func (c *Config) IsDevelopment() bool {
	return c.App.Env == "development"
}

// GetDSN returns database connection string
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
}
