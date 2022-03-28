# Meli Proxy Challenge

## De que se compone?

### Meli Proxy
- Servicio desarrollado en Go para gestionar las peticiones

### Nginx
- Utilizado como balanceador entre las diferentes instancias de Meli Proxy

### Grafana y Prometheus
- Utilizamos Grafana y Prometheus para visualizar y almacenar metricas respectivamente

### Locust
- Servicio utilizado para realizar tests de carga

### Redis
- Base de datos para el conteo de peticiones y control de criterios

### RedisInsight
- Se utiliza como visualizador de Redis

## Como ejecutarlo?
    docker-compose up -d --build
