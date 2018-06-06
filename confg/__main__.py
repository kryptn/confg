import logging

import click

logger = logging.getLogger(__name__, )
logger.setLevel(logging.DEBUG)


@click.command()
@click.option('--file', default='confg.toml', help='config file location')
@click.option('--template', default='config.tmpl', help='template file location')
@click.option('--output', default=None, help='rendered file location')
def cli(file, template, output):
    import toml, pprint

    from confg.engine import validate_config, do

    with open(file) as fd:
        config = toml.load(fd)

    validate_config(config)

    blocks, reduced = do(config)

    print('\n\n -- debug --')
    pprint.pprint(config)
    pprint.pprint(blocks)
    pprint.pprint(reduced)

    print(toml.dumps(reduced))


if __name__ == '__main__':
    cli()
