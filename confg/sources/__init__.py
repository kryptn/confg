from typing import Generic

from confg._abc import AbstractSource

DEFAULT_SOURCE = 'empty'


def empty():
    from .empty import EmptySource
    return EmptySource


def env():
    from .env import EnvSource
    return EnvSource


def etcd():
    from .etcd import EtcdSource
    return EtcdSource


def vault():
    from .vault import VaultSource
    return VaultSource


_available_sources = {
    'env': env,
    'etcd': etcd,
    'vault': vault,
    'empty': empty,
}


class InvalidSourceException(Exception):
    pass


class UninstalledSourceException(Exception):
    pass


def valid_source(source_name):
    return source_name in _available_sources


def get_block_source(source_name):
    if not valid_source(source_name):
        msg = f"{source_name} not a valid source"
        raise InvalidSourceException(msg)

    import_source = _available_sources.get(source_name)

    try:
        return import_source()
    except ImportError:
        raise UninstalledSourceException()


