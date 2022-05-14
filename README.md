# Sistemas Distribuidos - Ayudantia

El siguiente repositorio contiene ejemplos para las ayudantias del curso sistemas distribuidos, cabe destacar que los proyectos estan dockerizados, basta con ejecutarlos para probarlos.

## Caching con Redis

Contiene una simple api en nodejs con express, la cual se conecta con redis para una api de manejo de sesiones.

## gRPC

Contiene tres simples api en golang, python y node. Las cuales se comunican mediante gRPC para completar una query.

## Kafka

docker-compose exec kafka /opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=259200000 --topic ordenes

docker-compose exec kafka /opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=259200000 --topic resumenes

docker-compose exec kafka /opt/bitnami/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list