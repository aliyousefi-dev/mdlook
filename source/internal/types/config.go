package types

type MdLookConfig struct {
	DocName                string `json:"docname"`
	DisableThemes          bool   `json:"disableThemes"`
	DisablePrintOptions    bool   `json:"disablePrintOptions"`
	DisableMarkdownOptions bool   `json:"disableMarkdownOptions"`
	GitUrl                 string `json:"GitUrl"`
	SyncNav                bool   `json:"syncNav"`
}

func GetDefaultConfigData() MdLookConfig {
	return MdLookConfig{
		DocName:                "Documentation",
		DisableThemes:          false,
		DisablePrintOptions:    false,
		DisableMarkdownOptions: false,
		GitUrl:                 "https://github.com/user/repo",
		SyncNav:                true,
	}
}
