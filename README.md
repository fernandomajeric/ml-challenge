# ML - Challenge

## Analisis de Acciones Fraudulentas

Contenido:

- [Enunciado](#enunciado)
- [Ejecucion Local](#ejeucion-local)
- [Documentacion Api](#documentacion-api)

## Enunciado

Para coordinar acciones de respuesta ante fraudes, es útil tener disponible información
contextual del lugar de origen detectado en el momento de comprar, buscar y pagar. Para ello, entre otras fuentes, se decide crear una herramienta que dado un IP obtenga información asociada:

Construir una aplicación que dada una dirección IP, encuentre el país al que pertenece, y muestre:

- El nombre y código ISO del país
- Los idiomas oficiales del país
- Hora(s) actual(es) en el país (si el país cubre más de una zona horaria, mostrar todas)
- Distancia estimada entre Buenos Aires y el país, en km.
- Moneda local, y su cotización actual en dólares (si está disponible)

Basado en la información anterior, es necesario contar con un mecanismo para poder consultar las siguientes estadísticas de utilización del servicio con los siguientes agregados
- Distancia más lejana a Buenos Aires desde la cual se haya consultado el servicio
- Distancia más cercana a Buenos Aires desde la cual se haya consultado el servicio
- Distancia promedio de todas las ejecuciones que se hayan hecho del servicio.

Ejemplo:

> traceip 83.44.196.93
IP: 83.44.196.93, fecha actual: 21/11/2016 16:01:23
País: España (spain)
ISO Code: es
Idiomas: Español (es)
Moneda: EUR (1 EUR = 1.0631 U$S)
Hora: 20:01:23 (UTC) o 21:01:23 (UTC+01:00)
Distancia estimada: 10270 kms (-34, -64) a (40, -4)


## 1) Ejecucion Local

*Requisitos:* Se debe tener instalado docker (docker-compose)

1. Ejecutar el siguiente comando para clonar el proyecto: `https://github.com/fernandomajeric/ml-challenge.git`
2. Ejecutar `docker-compose build`
3. Luego `docker-compose up`

## 2) Endpoints

#### Estadisticas de los servicios

Obtiene las estadisticas de utilizacion de los servicios (pais, distancia mas proxima a bs as , invocaciones y distancia promedio) 

`$ curl -X GET \ 'http://localhost:8080/statistics' \ -H 'accept: application/json'`

#### Traceo de Ip
Obtiene pais de procedencia, moneda, idioma a traves de una ip valida.

`$ curl -X GET \ 'http://localhost:8080/trace-ip/139.82.0.0' \ -H 'accept: application/json'`


