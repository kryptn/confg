import os

from confg._abc import AbstractSource


class EmptySource(AbstractSource):

    def __init__(self, *args, **kwargs):
        pass

    def render(self, *args, **kwargs):
        return {}
