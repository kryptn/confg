import os
from typing import Optional

from confg.sources.base import Source

def valid_env_line(line):
    return len(line) > 3 and '=' in line and not line.startswith('#')

def parse_line(line):
    # expecting lines like this:
    # SOME_VAL=Abcd
    key, val = line.split('=', 1)
    return {key.strip(): val.strip()}


def read_envs_from_file(filename):
    with open(filename) as fd:
        # ignore commented lines
        lines = [l for l in fd.readlines() if valid_env_line(l)]
    env_map = {}
    for line in lines:
        env_map.update(parse_line(line))

    return env_map


def first(gen):
    for item in gen:
        if item:
            return item


class EnvSource(Source):

    def __init__(self, **config):
        super().__init__(**config)
        self.env_map = {}

        if 'file' in config:
            filename = config.get('file')
            self.env_map = read_envs_from_file(filename)

    def env_attempts(self, lookup: str) -> Optional[str]:
        """

        :param lookup: a string key to check a loaded environment
                       file first then the environment
        :return: the resulting key or None
        """
        yield self.env_map.get(lookup)
        yield os.environ.get(lookup)

    def retrieve(self, lookup):
        return first(self.env_attempts(lookup))
