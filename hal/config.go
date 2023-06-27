package hal

import (
	"context"

	commonpb "github.com/Intrising/intri-type/common"
	configpb "github.com/Intrising/intri-type/config"

	utilsMisc "github.com/Intrising/intri-utils/misc"
	utilsRpc "github.com/Intrising/intri-utils/rpc"
	"google.golang.org/protobuf/proto"
)

type ConfigClient struct {
	SavedPath   string
	DefaultPath string

	Ctx    context.Context
	Client configpb.RunServiceClient
}

func ConfigClientInit(ctx context.Context, service commonpb.ServicesEnumTypeOptions) *ConfigClient {
	client := utilsRpc.NewClientConn(ctx, service, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CONFIG)
	return &ConfigClient{
		Ctx:    ctx,
		Client: configpb.NewRunServiceClient(client.GetGrpcClient()),
	}
}

func (c *ConfigClient) RestoreDefault(in configpb.FactoryDefaultModeTypeOptions) error {
	_, err := c.Client.RunRestoreDefaultConfig(c.Ctx, &configpb.RestoreDefaultRequest{Type: in})
	return err
}

func (c *ConfigClient) SetPath(savedPath, defaultPath string) {
	c.SavedPath = savedPath
	c.DefaultPath = defaultPath
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
