import etcd

from confg._abc import AbstractSource


class EtcdSourceConfigError(Exception):
    pass


def read_or_none(client: etcd.client, key, **args):
    try:
        return client.read(key, **args).value
    except etcd.EtcdKeyNotFound:
        return None


class EtcdSource(AbstractSource):

    def __init__(self, **client_config):
        config = client_config or {}
        if not self.verify_config(client_config):
            raise EtcdSourceConfigError()

        self.client = etcd.Client(**config)

    def render_slug(self, slug, **args):
        return read_or_none(self.client, slug, **args)

    def render(self, key_slugs):
        items = {}
        for key, slug in key_slugs.items():
            items[key] = self.render_slug(slug)
        return items

    @staticmethod
    def verify_config(config):
        is_type = lambda obj, type_: isinstance(obj, type_)
        # empty config
        if is_type(config, dict) and not config:
            return True

        return True
