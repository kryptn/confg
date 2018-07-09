import logging
from typing import List, Mapping

import toml

from confg.backend import default_backends
from confg.group import Key

logging.basicConfig()
logger = logging.getLogger(__name__)

logger.setLevel(logging.DEBUG)


def keys_From_group(name: str, data) -> List[Key]:
    dest = name
    backend = data.get('backend', None)
    priority = data.get('priority', 1)


class Confg:

    _protected_keys = ('config', 'confg', 'backend', 'default')

    default_priority = 1

    orig_config = None

    keys = []
    backends = []
    groups = []

    def __init__(self, config):
        self.backends = default_backends()

        if not isinstance(config, dict):
            config = self.from_file(config)

        self.orig_config = config

        self.validate_config(config)

        self.ingest(config)

    def keys_from_group(self, name, data):

        group_default = {
            'dest': name,
            'backend': data.get('backend'),
            'priority': data.get('priority', self.default_priority),
        }

        def Key_from_data(name, data):
            if not isinstance(data, Mapping):
                data = {'name': name, 'lookup': data}

            defaults = group_default.copy()
            defaults.update(data)

            return Key(**defaults)

        keys = [Key_from_data(*items) for items in data['keys'].items()]

        return keys

    def keys_from_groups(self, config):
        groups = ((k,v) for k,v in config.items() if k not in self._protected_keys)

        for group in groups:
            self.keys.extend(self.keys_from_group(*group))


    def validate_config(self, config):
        return True

    def from_file(self, filename):
        with open(filename) as fd:
            raw_config = toml.load(fd)

        config = raw_config
        return config

    def ingest(self, config):
        pass

    def from_dict(self, data):
        pass


from confg import Confg

settings = Confg
