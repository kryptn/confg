from typing import Dict, Mapping

from confg.backend import Backend
from confg.sources.base import Source

DEFAULT_GROUP_NAME = 'default'


class Key:
    """
    represents a named config value with optional
    specifically defined source
    """
    key: str
    lookup: str

    value: str
    backend: Backend

    def __init__(self, key, value=None, backend=None):
        """

        :param key: String name for the Key
        :param value: Either a string lookup
        :param backend:
        """
        if not isinstance(value, Mapping):
            value = {
                'lookup': value
            }

        self.key = value.get('key', key)
        self.lookup = value.get('lookup')
        self.backend = value.get('backend', backend)
        self.value = None

    def set_default(self, value):
        self.value = value

    def retrieve(self, source: Source):
        obj = source.retrieve(self.lookup)
        if obj:
            self.value = obj


class Group:
    name: str
    default_source: str
    priority: int

    keys: Dict[str, Key]

    _required_group_keys = ('name', 'backend')

    def __init__(self,
                 name: str,
                 backend: Backend,
                 priority: int = 0,
                 keys: Mapping = None,
                 defaults: Mapping = None):
        """

        :param name: The group name
        :param backend: The user defined backend name
        :param priority: An integer used to determine priority
        :param keys:
        :param defaults:
        """
        self.name = name
        self.backend = backend
        self.priority = priority

        self.keys = {}

        for key, value in (keys or {}).items():
            self.keys[key] = Key(key, value, self.backend)

        for key, value in (defaults or {}).items():
            if key not in self.keys:
                self.keys[key] = Key(key)
            self.keys[key].set_default(value)

    @classmethod
    def validate_group(cls, name, group):
        valid = True
        if 'source' not in group:
            valid = False

        if name != DEFAULT_GROUP_NAME:
            if any(key not in group for key in cls._required_group_keys):
                valid = False

        return valid

    def asdict(self):
        return {name: k.value for name, k in self.keys.items()}
