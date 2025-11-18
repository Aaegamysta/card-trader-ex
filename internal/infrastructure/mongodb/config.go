package mongodb

type Config struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Project  string `toml:"project"`
	Cluster  string `toml:"cluster"`
}
