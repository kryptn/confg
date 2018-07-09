from collections import namedtuple
from logging import getLogger

from confg.backend import Backend, default_backends
from confg.group import Group

logger = getLogger(__name__)

protected_top_level_keys = (
    'config',
    'confg',
    'backend',
    'default',
    'coalesce',
)


def unprotected_items(data, protected_keys):
    return ((k, v) for k, v in data.items() if k not in protected_keys)


def check_config(config) -> (bool, list):
    return True, []


Config = namedtuple('Confg', ('backends', 'groups'))


def ingest_v1_config(config):
    valid, results = check_config(config)
    if not valid:
        logger.error(results)
        return
    if results:
        logger.warning(results)

    backends = default_backends()

    raw_backends = config.get('backend', {})
    for name, data in raw_backends.items():
        logger.info('parsing backend %s', name)
        if name not in backends:
            source = data.pop('source')
            backends[name] = Backend(name=name, source=source)

        backends[name].update(**data)

    groups = {}
    for name, data in unprotected_items(config, protected_top_level_keys):
        logger.info('parsing group %s', name)
        group_backend_name = data.pop('backend')
        group_backend = backends.get(group_backend_name)
        groups[name] = Group(name=name, backend=group_backend, **data)

    for group_name, group in groups.items():
        logger.info('registering keys for group %s', group_name)
        for key_name, key in group.keys.items():
            logger.info('registering %s -- %s', group_name, key_name)
            if isinstance(key.backend, str):
                logger.info('matching key %s backend %s', key_name, key.backend)
                key.backend = backends.get(key.backend)

            key.backend.source.register(key)

    return Config(backends=backends, groups=groups)


def resolve_config(config: Config):
    for backend_name, backend in config.backends.items():
        backend.source.retrieve_all()

    return {name: group.asdict() for name, group in config.groups.items()}
