from confg._abc import AbstractSource


class Source(AbstractSource):
    known_keys: list = None

    def __init__(self, **config):
        super().__init__(**config)
        self.known_keys = []

    def retrieve(self, lookup):
        pass

    def register(self, key):
        self.known_keys.append(key)

    def retrieve_all(self):
        for key in self.known_keys:
            key.retrieve(self)
