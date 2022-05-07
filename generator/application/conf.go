package application

const (
	dubbogoConfigFile = `dubbo:
  protocols:
    triple:
      name: tri
      port: 20000
      interface: test
`
)

func init() {
	fileMap["dubbogoConfigFile"] = &fileGenerator{
		path:    "./conf",
		file:    "dubbogo.yaml",
		context: dubbogoConfigFile,
	}
}
