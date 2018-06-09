# confg


## Concepts:

Group:
  - toml table that logically groups like configs
  - named
  
Backend:
  - named reference to a source with optional configs
  - allows multiple named backends to the same source
  
Source:
  - the mechanism to retrieve a key for a block

Engine:
  - Ingests the config from a file or dictionary and produces the config
  
## process

1. read in config
  
## todo:

- make every key within the Block.keys dict be an object that stores retrieval metadata
  - also allows setting up keys to be live-queryable or watchable
  - progress on this, each key on a group is its own object -- should be able to add extras here
- add validation to the objects

## maybes:

- determine which packages need to be installed from just the toml config

## eventually:

- jinja templating -- can send entire bocks as the contexts

## reference

- [toml](https://github.com/toml-lang/toml)

- [click](http://click.pocoo.org/6/)
- [jinja](http://jinja.pocoo.org/)



## sources

- env via os.environ
- [vault via hvac](https://github.com/ianunruh/hvac)
- [etcd via python-etcd](http://python-etcd.readthedocs.io/en/latest/)
- [aws parameter store (SSM) via boto](https://boto3.readthedocs.io/en/latest/reference/services/ssm.html)
