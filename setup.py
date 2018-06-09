from setuptools import find_packages, setup

setup_args = {
    'name': 'confg',
    'version': '0.0.1',
    'description': 'config utility',
    'packages': find_packages(exclude=['contrib', 'docs', 'tests']),
    'entry_points': {
        'console_scripts': [
            'confg = confg.__main__:cli'
        ],
    },
    'install_requires': [
        'toml==0.9.4',
        'click==6.7',

    ],
    'extras_require': {
        'dev': [
            'ipdb',
        ],
        'test': [
            'atomicwrites==1.1.5',
            'attrs==18.1.0',
            'more-itertools==4.2.0',
            'pluggy==0.6.0',
            'py==1.5.3',
            'pytest==3.6.0',
            'six==1.11.0'
        ],
        'etcd': [
            'python-etcd==0.4.5',
            'urllib3==1.23',
            'dnspython==1.15.0'
        ],
        'vault': [
            'certifi==2018.4.16',
            'chardet==3.0.4',
            'hvac==0.5.0',
            'idna==2.6',
            'requests==2.18.4',
            'urllib3==1.23',  # iffy, technically requires 1.22
        ],
    },
}

setup(**setup_args)
