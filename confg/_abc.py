from abc import ABC, abstractmethod


class AbstractBlock(ABC):

    def render(self):
        ...


class AbstractSource(ABC):

    @abstractmethod
    def __init__(self, **config):
        self._config = config

    @abstractmethod
    def retrieve(self, lookup: str):
        ...

    @abstractmethod
    def retrieve_all(self):
        ...

    @abstractmethod
    def register(self, key):
        ...
