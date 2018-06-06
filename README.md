# confg


## Concepts:

Block:
  - toml table that is constrained to one source
  
Source:
  - the mechanism to retrieve a key for a block

Engine:
  - the part that reads and validates the toml config
  
## todo:

- make every key within the Block.keys dict be an object that stores retrieval metadata
  - also allows setting up keys to be live-queryable or watchable
- make one config block per source that can be referenced throughout
- ensure source imports are one-time

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


