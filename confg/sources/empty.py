import os

from confg._abc import AbstractSource


class EmptySource(AbstractSource):

    def render(self, key_slugs):
        return {}
