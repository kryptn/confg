cp ../../confg .s
HTTP_HOST=localhost HTTP_PORT=9090 ./confg -f defaults.toml -o settings_default.toml
HTTP_HOST=localhost HTTP_PORT=9090 ./confg -f defaults.toml -f local.toml -o settings_local.toml
HTTP_HOST=localhost HTTP_PORT=9090 ./confg -f defaults.toml -f local.toml -f prod.toml -o settings_prod.toml