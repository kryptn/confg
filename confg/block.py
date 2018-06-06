from logging import getLogger
from operator import attrgetter
from typing import Dict, Any

from confg.sources import get_block_source, DEFAULT_SOURCE

logger = getLogger(__name__)


class Block:
    name: str = 'config'
    source: str = None
    priority: int = 1

    keys: Dict[str, str] = None
    defaults: Dict[str, Any] = None

    def __init__(self, block_name, block):
        self.name = block_name,
        self.source = block.get('source') or DEFAULT_SOURCE
        self.priority = block.get('priority') or 0

        self.keys = block.get('keys') or {}
        self.defaults = block.get('defaults') or {}
        self.rendered = self.defaults.copy()

    def render(self):
        source = get_block_source(self.source)()
        rendered = source.render(self.keys)
        self.rendered.update(rendered)


def blocks_from_config(config):
    empty_block = []
    blocks = []
    for block_name, block_data in config.items():
        block = Block(block_name, block_data)
        if block.source == DEFAULT_SOURCE:
            empty_block.append(block)
        else:
            blocks.append(block)

    sorted_blocks = sorted(blocks, key=attrgetter('priority'))
    empty_block.extend(sorted_blocks)
    return empty_block


def reduce_blocks(blocks):
    reduced = {}
    for block in blocks:
        truthy = {k: v for k, v in block.rendered.items() if v}
        reduced.update(truthy)
    return reduced
