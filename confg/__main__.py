import logging

import click

logger = logging.getLogger()
logger.setLevel(logging.DEBUG)


@click.command()
@click.option('--file', default='confg.toml', help='config file location')
@click.option('--template', default='config.tmpl', help='template file location')
@click.option('--output', default=None, help='rendered file location')
def cli(file, template, output):
    import toml

    from confg.engine import ingest_config, resolve_config

    with open(file) as fd:
        config = toml.load(fd)

    config = ingest_config(config)

    data = resolve_config(config)

    print(toml.dumps(data))


if __name__ == '__main__':
    cli()
