# Confg

like confd but generalized

Backends are the containers for sources, where the values come from. all
config values for the backend goes here.

```
[backend.env_backend]
source = "env"
other_value = true
```

Groups are logical combinations of keys and provide default values for
each key.

Groups must provide a backend that matches one defined within the
`backend` map.

```
[web_config]
backend = "env_backend"

```


Keys are defined within each group as a submap. Simple keys are just a
key/value pair where the key will be the output config key, and the
value is the lookup used in the source to get the value.

Because we're using the `env` source within `[backend.env_backend]` this
will just look in the environment for `HTTP_HOST` and `HTTP_PORT`.


```
[web_config.keys]
host = "HTTP_HOST"
port = "HTTP_PORT"
```

Complex keys define any value within the key they want, while still
inheriting the group's defaults if not defined.

Right now complex keys are the only way to set a default value for a
key.

```
[web_config.keys.db_url]
lookup = "DATABASE_URL"
default = "localhost"
```


these keys are functionally equivalent:

```
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



todo:
- make env source read from file too
- i don't like the `confgFromParsed` like functions
- add more sources (etcd and aws ssm for sure)
- allow stacking input configs
- output to more than toml
- decide if i want to redo the flagreader
- determine what top-level keys i need to reserve
  - known: `version`, `backend`
  - everything else is considered a group

