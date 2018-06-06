from abc import ABC, abstractmethod


class AbstractBlock(ABC):



    def render(self):
        ...


class AbstractSource(ABC):

    def render_slug(self, slug):
        ...

    @abstractmethod
    def render(self, block: AbstractBlock):
        ...

    def clean(self, config):
        return config
