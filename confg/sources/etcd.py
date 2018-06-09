import etcd

from confg.sources.base import Source


class EtcdSourceConfigError(Exception):
    pass


def read_or_none(client: etcd.client, key, **args):
    try:
        return client.read(key, **args).value
    except etcd.EtcdKeyNotFound:
        return None


class EtcdSource(Source):

    def __init__(self, **client_config):
        super().__init__(**client_config)
        config = client_config or {}
        if not self.verify_config(client_config):
            raise EtcdSourceConfigError()

        self.client = etcd.Client(**config)

    def retrieve(self, lookup, **args):
        return read_or_none(self.client, lookup, **args)

    @staticmethod
    def verify_config(config):
        is_type = lambda obj, type_: isinstance(obj, type_)
        # empty config
        if is_type(config, dict) and not config:
            return True

        return True
