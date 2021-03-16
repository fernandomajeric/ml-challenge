# ML - Challenge

## Analisis de Riesgo

Contenido:

- [Enunciado](#enunciado)
- [Ejecucion Local](#ejeucion-local)
- [Documentacion Api](#documentacion-api)

## Enunciado

Ejercicio de programación Para coordinar acciones de respuesta ante fraudes, es útil tener disponible información
contextual del lugar de origen detectado en el momento de comprar, buscar y pagar. Para ello, entre otras fuentes, se
decide crear una herramienta que dado un IP obtenga información asociada:
Construir una aplicación que dada una dirección IP, encuentre el país al que pertenece, y muestre:

- El nombre y código ISO del país
- Los idiomas oficiales del país
- Hora(s) actual(es) en el país (si el país cubre más de una zona horaria, mostrar todas)
- Distancia estimada entre Buenos Aires y el país, en km.
- Moneda local, y su cotización actual en dólares (si está disponible)
  Basado en la información anterior, es necesario contar con un mecanismo para poder consultar las siguientes
  estadísticas de utilización del servicio con los siguientes agregados
- Distancia más lejana a Buenos Aires desde la cual se haya consultado el servicio
- Distancia más cercana a Buenos Aires desde la cual se haya consultado el servicio
- Distancia promedio de todas las ejecuciones que se hayan hecho del servicio. Tener en cuenta que esta estadística
  puede recibir fluctuaciones agresivas de tráfico (Entre 100 y 1 millón de peticiones por segundo). Para resolver la
  información, pueden utilizarse las siguientes APIs públicas:
  Geolocalización de IPs: https://ip2country.info/
  Información de paises: http://restcountries.eu/
  Información sobre monedas: http://fixer.io/
- La aplicación puede ser en línea de comandos o web. En el primer caso se espera que el IP sea un parámetro, y en el
  segundo que exista un form donde escribir la dirección.
- La aplicación deberá hacer un uso racional de las APIs, evitando hacer llamadas innecesarias.
- La aplicación puede tener estado persistente entre invocaciones.
- Además de funcionamiento, prestar atención al estilo y calidad del código fuente.
- La aplicación deberá poder correr ser construida y ejecutada dentro de un contenedor Docker (incluir un Dockerfile e
  instrucciones para ejecutarlo). Ejemplo (el formato es tentativo y no tiene porque ser exactamente así):

> traceip 83.44.196.93 IP: 83.44.196.93, fecha actual: 21/11/2016 16:01:23 País: España (spain)
ISO Code: es Idiomas: Español (es)
Moneda: EUR (1 EUR = 1.0631 U$S)
Hora: 20:01:23 (UTC) o 21:01:23 (UTC+01:00)
Distancia estimada: 10270 kms (-34, -64) a (40, -4)

## 1) Ejecucion Local

*Requisitos:* Se debe tener instalado docker (docker-compose)

1. Ejecutar el siguiente comando para clonar el proyecto:
   `git clone https://github.com/fernandomajeric/mutant-recruiter.git`
2. Ejecutar `docker-compose build`
3. Luego `docker-compose up`

## Documentacion Api

[Documentacion Api](http://localhost:8080/ml-challenge/swagger-ui.html#/)

- POST /api/country/info Ingresando una IP valida retorna informacion del pais de procedencia, moneda, idioma, etc

- GET /api/statistics Retorna estadisticas basada en las IP ingresadas.


