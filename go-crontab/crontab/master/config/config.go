package config

import "github.com/spf13/viper"

var (
	G_config *Config
)

type Config struct {
	ApiPort               int      `json:"apiPort"`
	ApiReadTimeout        int      `json:"apiReadTimeout"`
	ApiWriteTimeout       int      `json:"apiWriteTimeout"`
	EtcdEndpoints         []string `json:"etcdEndpoints"`
	EtcdDialTimeout       int      `json:"etcdDialTimeout"`
	MongodbUri            string   `json:"mongodbUri"`
	MongodbConnectTimeout int      `json:"mongodbConnectTimeout"`
	Mode                  string   `json:"mode"`
}

/**
初始化配置：
问题：
1）如何使用viper读取并封装成一个map类型的值
2）是否可以使用viper读取配置并封装成一个
 */
func InitConfig(configFile string) (err error) {
	var (
		v                     *viper.Viper
		apiPort               int
		apiReadTimeout        int
		apiWriteTimeout       int
		etcdEndpoints         []string
		etcdDialTimeout       int
		mongodbUri            string
		mongodbConnectTimeout int
		mode                  string
	)
	// 生成一个viper去读取配置文件的内容:
	v = viper.New()
	// 设置读取文件路径
	v.SetConfigFile(configFile)
	// 读取配置
	if err = v.ReadInConfig(); err != nil{
		return err
	}
	apiPort = v.GetInt("api.apiPort")
	apiReadTimeout = v.GetInt("api.apiReadTimeout")
	apiWriteTimeout = v.GetInt("api.apiWriteTimeout")
	etcdEndpoints = v.GetStringSlice("etcd.etcdEndPoints")
	etcdDialTimeout = v.GetInt("etcd.etcdDialTimeout")
	mongodbUri = v.GetString("mongodb.mongodbUri")
	mongodbConnectTimeout = v.GetInt("mongodb.mongodbConnectTimeout")
	mode = v.GetString("mode")

	G_config = &Config{}

	G_config.ApiPort = apiPort
	G_config.ApiReadTimeout = apiReadTimeout
	G_config.ApiWriteTimeout = apiWriteTimeout
	G_config.EtcdEndpoints = etcdEndpoints
	G_config.EtcdDialTimeout = etcdDialTimeout
	G_config.MongodbUri = mongodbUri
	G_config.MongodbConnectTimeout = mongodbConnectTimeout
	G_config.Mode = mode

	return nil
}
