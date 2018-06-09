from functools import lru_cache

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


def valid_source(source_kind):
    return source_kind in _available_sources


@lru_cache()
def get_source(kind, **config):
    if not valid_source(source_kind=kind):
        msg = "%s not a valid source kind"
        raise InvalidSourceException(msg, kind)

    source_importer = _available_sources.get(kind)

    try:
        Source = source_importer()
    except ImportError:
        msg = 'Source %s not installed'
        raise UninstalledSourceException(msg, kind)

    return Source(**config)


def get_block_source(source_name):
    if not valid_source(source_name):
        msg = f"{source_name} not a valid source"
        raise InvalidSourceException(msg)

    import_source = _available_sources.get(source_name)

    try:
        return import_source()
    except ImportError:
        raise UninstalledSourceException()
