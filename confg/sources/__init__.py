from .empty import EmptySource
from .env import EnvSource
from .etcd import EtcdSource
from .vault import VaultSource

DEFAULT_SOURCE = 'empty'


_available_sources = {
    'env': EnvSource,
    'etcd': EtcdSource,
    'vault': VaultSource,
    'empty': EmptySource,
}


class InvalidSourceException(Exception):
    pass


def valid_source(source_name):
    return source_name in _available_sources


def get_block_source(source_name):
    if not valid_source(source_name):
        msg = f"{source_name} not a valid source"
        raise InvalidSourceException(msg)

    return _available_sources.get(source_name)
