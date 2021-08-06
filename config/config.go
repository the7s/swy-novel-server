package config

type Configuration struct{}

var Config = new(Configuration)

const webUrl = "http://www.fqxsw.org"

func (config Configuration) GetWebUrl() string {

	return webUrl
}
