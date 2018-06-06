from logging import getLogger

from confg.sources import valid_source, DEFAULT_SOURCE
from confg.block import blocks_from_config, reduce_blocks

logger = getLogger(__name__)


def do(config):
    _clean_config(config)

    valid_config = validate_config(config)
    if not valid_config:
        logger.error('Cannot validate config -- exiting')
        return None, None

    blocks = blocks_from_config(config)
    for block in blocks:
        block.render()

    result = {block.name: block.rendered for block in blocks}
    reduced = reduce_blocks(blocks)

    return result, reduced


def _clean_config(config):
    for name, block in config.items():
        source = block.get('source', DEFAULT_SOURCE)
        block['source'] = source.lower()


def _one_or_none_empty_source(config):
    empty_block_names = []

    for name, block in config.items():
        if block.get('source') == DEFAULT_SOURCE:
            empty_block_names.append(name)

    return len(empty_block_names) < 2, empty_block_names


def _unique_priorities(config):
    priorities = []
    for name, block in config.items():
        priority = block.get('priority', 0)
        if priority in priorities:
            return False
        priorities.append(priority)
    return True


def validate_config(config):
    valid, empty_blocks = _one_or_none_empty_source(config)
    if len(empty_blocks) > 1:
        logger.error('More than one empty source block -- %s', empty_blocks)

    if not _unique_priorities(config):
        logger.warning('ambiguous priorities -- cannot guarantee application order')

    for name, block in config.items():
        logger.debug('validating block %s', name)

        source = block.get('source')
        if not valid_source(source):
            logger.error('%s -- invalid source %s', name, source)
            valid = False

        priority = block.get('priority', 0)
        if source != DEFAULT_SOURCE and not priority:
            logger.warning('%s -- missing or zero priority on non-empty block', name)

        if 'keys' not in block and source != DEFAULT_SOURCE:
            logger.warning('%s -- missing keys map -- will only apply defaults', name)

    return valid
