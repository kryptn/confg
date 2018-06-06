from abc import ABC

class AbstractBlock(ABC):



    def render(self):
        ...


class AbstractSource(ABC):

    def render_slug(self, slug):
        ...

    def render(self, block: AbstractBlock):
        ...
