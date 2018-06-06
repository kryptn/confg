from functools import partial
from itertools import chain
from logging import getLogger
from operator import attrgetter
from typing import Dict, Any, List

from confg.sources import get_block_source, DEFAULT_SOURCE

logger = getLogger(__name__)


class Block:
    name: str = 'config'
    source: str = None
    priority: int = 1

    keys: Dict[str, str] = None
    defaults: Dict[str, Any] = None
    config: Dict[str, Any] = None

    def __init__(self, block_name, block):
        self.name = block_name,
        self.source = block.get('source') or DEFAULT_SOURCE
        self.priority = block.get('priority') or 0

        self.keys = block.get('keys') or {}
        self.defaults = block.get('defaults') or {}
        self.rendered = self.defaults.copy()

        self.config = block.get('config') or {}

    def render(self):
        Source = get_block_source(self.source)
        if self.config:
            Source = partial(Source, **self.config)
        source = Source()
        logger.debug('attempting to render from source: %s', self.source)
        rendered = source.render(self.keys)
        self.rendered.update(rendered)


def blocks_from_config(config):
    overlay = []
    blocks = []
    for block_name, block_data in config.items():
        block = Block(block_name, block_data)
        if block.source == DEFAULT_SOURCE:
            overlay.append(block)
        else:
            blocks.append(block)

    sorted_blocks = sorted(blocks, key=attrgetter('priority'))
    overlay.extend(sorted_blocks)
    return overlay


def reduce_blocks(blocks: List[Block]) -> dict:
    "take all truthy values and apply, keeping the last value provided"
    block_chain = chain(*(b.rendered.items() for b in blocks))
    base = {k: v for k, v in block_chain if v}
    return base
