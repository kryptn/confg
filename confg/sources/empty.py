from confg.sources.base import Source


class EmptySource(Source):

    def __init__(self, **config):
        super().__init__(**config)

    def retrieve(self, lookup: str):
        return None
