
import os

from confg._abc import AbstractSource


class EtcdSource(AbstractSource):

    def _make_attempts(self, slug):
        return [slug, slug.upper(), slug.lower()]

    def render_slug(self, slug):
        for attempt in self._make_attempts(slug):
            value = os.environ.get(attempt)
            if value:
                break
        else:
            value = None

        return value

    def render(self, key_slugs):
        items = {}
        for key, slug in key_slugs.items():
            items[key] = self.render_slug(slug)
        return items
