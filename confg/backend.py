from confg.sources import get_source

_protected_backend_names = ('env',)


def default_backends():
    env_backend = Backend('env', source='env')
    return {'env': env_backend}


class Backend:
    name: str
    source_kind: str
    source_config: dict

    _source = None

    def __init__(self, name, source=None, **source_config):
        self.name = name
        self.source_kind = source
        self.source_config = source_config

    def update(self, **source_config):
        self.source_config.update(source_config)

    @property
    def source(self):
        if not self._source:
            self._source = get_source(
                kind=self.source_kind,
                **self.source_config
            )

        return self._source

    @classmethod
    def validate(cls, config):
        valid = True

        return valid
