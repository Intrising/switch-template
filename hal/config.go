package hal

import (
	utilsMisc "github.com/Intrising/intri-utils/misc"
	"google.golang.org/protobuf/proto"
)

type ConfigClient struct {
	SavedPath   string
	DefaultPath string
}

func (c *ConfigClient) GetSavedConfigContent() (string, error) {
	yamlBs, err := utilsMisc.ReadFile(c.SavedPath)
	return string(yamlBs), err
}
func (c *ConfigClient) LoadSavedConfig(cfg proto.Message) error {
	yamlStr, err := c.GetSavedConfigContent()
	if err != nil {
		return err
	}
	return utilsMisc.UnmarshalYamlToProtoMessage(yamlStr, cfg)
}
func (c *ConfigClient) RestoreSavedConfig(cfg proto.Message) error {
	yamlStr, err := c.GetSavedConfigContent()
	if err != nil {
		return err
	}
	return utilsMisc.RestoreConfig(cfg, yamlStr)
}
func (c *ConfigClient) SaveSavedConfig(cfg proto.Message) error {
	yamlStr, err := utilsMisc.MarshalProtoMessageToYamlData(cfg)
	if err != nil {
		return err
	}
	return utilsMisc.WriteFile(c.SavedPath, yamlStr)
}
func (c *ConfigClient) SaveDefaultConfig(cfg proto.Message) error {
	yamlStr, err := utilsMisc.MarshalProtoMessageToYamlData(cfg)
	if err != nil {
		return err
	}
	return utilsMisc.WriteFile(c.DefaultPath, yamlStr)
}
