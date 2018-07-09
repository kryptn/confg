test-files:
	confg --file examples/second/confg.toml
etcd-load:
	etcdctl set /confg/services/amqp/url amqp://u:p@rmq:5672/
	etcdctl set /confg/services/redis/url redis://localhost
etcd-unload:
	etcdctl rmdir /confg
