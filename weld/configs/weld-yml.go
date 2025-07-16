package configs;
// weld.yml

import(
	_ "gopkg.in/yaml.v2"
);



type WeldYml struct {
	Project  ProjectYml `yaml:"Project"`
	Projects []string   `yaml:"Projects"`
}

type ProjectYml struct {
	Name       string   `yaml:"Name"`
	Desc       string   `yaml:"Desc"`
	Aliases    []string `yaml:"Aliases"`
	Arch       string   `yaml:"Arch"`
	License    string   `yaml:"License"`
	URL        string   `yaml:"URL"`
	PublicRepo string   `yaml:"PublicRepo"`
	Requires   []string `yaml:"Requires"`
	Recommends []string `yaml:"Recommends"`
	Generate   []string `yaml:"Generate"`
	TagFiles   []string `yaml:"TagFiles"`
	Files      []string `yaml:"Files"`
}
