package config

// Configuration for app
type Configuration struct {
	ListenOn         string `mapstructure:"LISTEN_ON"`
	UFOAddress       string `mapstructure:"UFO_ADDRESS"`
	UFOReset         string `mapstructure:"UFO_RESET"`
	CameraPollerDown string `mapstructure:"CAMERA_POLLER_DOWN"`
	CameraPollerUp   string `mapstructure:"CAMERA_POLLER_UP"`
}
