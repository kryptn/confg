import pytest
from hypothesis import given, strategies as s, assume

from confg.sources.env import parse_line, read_envs_from_file

pytestmark = pytest.mark.env()


@s.composite
def env_key(draw, key=s.text(min_size=1), value=s.text(min_size=1)):
    ex_key = draw(key).strip().replace('\n', '')
    ex_value = draw(value).strip().replace('\n', '')
    test = f"{ex_key}={ex_value}"
    return test, {ex_key: ex_value}

@given(env_key())
def test_parse_line(line_expected):
    line, expected = line_expected
    assert parse_line(line) == expected

@s.composite
def env_file(draw, line=env_key()):
    expected = {}
    items = draw(s.lists(line, min_size=1))
    lines = []
    for item, line_expected in items:
        expected.update(line_expected)
        lines.append(item)

    env_file_contents = '\n'.join(lines)
    return env_file_contents, items, expected



@given(file_data_expected=env_file())
def test_read_envs_from_file(file_data_expected, tmpdir):
    file_data, orig, expected = file_data_expected

    p = tmpdir.join(".env")
    p.write(file_data)

    map = read_envs_from_file(str(p))

    assert map == expected


def test_EnvSource_happy_path():
    pass


def test_EnvSource_from_file():
    pass


def test_EnvSource_file_and_env():
    pass


