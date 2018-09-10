# Confg

[![Build Status](https://travis-ci.org/kryptn/confg.svg?branch=master)](https://travis-ci.org/kryptn/confg)

### Config Gatherer

Inspiration from confd but able to use multiple backend sources. Wrote
my own because [confd will not support multiple backends.][1]

## Concepts

Backends are the containers for sources, where the values come from. all
config values for the backend goes here.

```toml
[backend.env_backend]
source = "env"
```

Groups are logical combinations of keys and provide default values for
each key.

Groups must provide a backend that matches one defined within the
`backend` map.

```toml
[web_config]
backend = "env_backend"
```


Keys are defined within each group as a submap. Simple keys are just a
key/value pair where the key will be the output config key, and the
value is the lookup used in the source to get the value.

Because we're using the `env` source within `[backend.env_backend]` this
will just look in the environment for `HTTP_HOST` and `HTTP_PORT`.


```toml
[web_config.keys]
host = "HTTP_HOST"
port = "HTTP_PORT"
```

Complex keys define any value within the key they want, while still
inheriting the group's defaults if not defined.

Right now complex keys are the only way to set a default value for a
key.

```toml
[web_config.keys.db_url]
lookup = "DATABASE_URL"
default = "localhost"
```


These keys are functionally equivalent:

```toml
[web_config.keys]
host = "HTTP_HOST"

[web_config.keys.host]
lookup = "HTTP_HOST"

[web_config.keys.web_host]
lookup = "HTTP_HOST"
key = "host"

[other_web_config.keys.host]
lookup = "HTTP_HOST"
dest = "web_config"
```

example:

```toml
[backend.env_backend]
source = "env"

[web_config]
backend = "env_backend"

[web_config.keys]
host = "HTTP_HOST"
port = "HTTP_PORT"

[web_config.keys.db_url]
lookup = "DATABASE_URL"
default = "localhost"
```

run with: `HTTP_HOST=localhost HTTP_PORT=8080 ./confg -f example_confg.toml -o example_settings.toml`

outputs:

```toml
[web_config]
  db_url = "localhost"
  host = "localhost"
  port = "8080"
```


### To do

- add prefixing within groups
  - would be useful for ssm and etcd
  - not sure how much i want to do this
- output to more than toml
  - json and yaml?
  - another package, new variable or read extension?
- ci/cd
- backend inheritance?
- should the group function as a default key instead of its own type?
  - would this open the door to recursive keys?
- add a meta output


[1]: https://github.com/kelseyhightower/confd/issues/414#issuecomment-232388171