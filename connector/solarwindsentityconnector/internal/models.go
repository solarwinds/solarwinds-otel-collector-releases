package internal

type Entity struct {
	Type       string   `mapstructure:"entity"`
	IDs        []string `mapstructure:"id"`
	Attributes []string `mapstructure:"attributes"`
}

type Events struct {
	Relationships []Relationship `mapstructure:"relationships"`
}

type Relationship struct {
	Type        string   `mapstructure:"type"`
	Source      string   `mapstructure:"source_entity"`
	Destination string   `mapstructure:"destination_entity"`
	Attributes  []string `mapstructure:"attributes"`
}
