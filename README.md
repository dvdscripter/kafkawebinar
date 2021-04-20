# Requerimentos

docker e docker-compose

# Executando

Abra um terminal e execute o comando abaixo.

```
docker-compose up
```

Três containers agora estão em execução: kafkawebinar (exemplos em go), kafka e o zookeeper.

Abra um novo terminal e execute:

```
docker exec -it kafkawebinar ./consumer
```

Em outro terminal execute:

```
docker exec -it kafkawebinar ./producer
```

# Experimentos

Neste momento você tem os dois exemplos (produtor e consumidor) em execução. Alguns experimentos possíveis:
* O que acontece quando paramos a execução do consumidor?
* O que acontece quando reiniciamos o consumidor?
* O que acontece com consumidor se o produtor é parado?
* O que acontece se temos dois consumidores ao mesmo tempo?