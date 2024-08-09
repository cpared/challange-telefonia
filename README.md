# Ejercicio Técnico

## Tabla de contenidos

## Introduccion

El siguiente proyecto es un challange tecnico el cual representa el desarrollo de una app
para una empresa de generacion de facturas telefonicas.

*El input del programa es el siguiente: *

- Número de teléfono del usuario a generar la factura
- Periodo de facturación AAAA-MM-DD/AAAA-MM-DD
- CSV con una lista de llamadas, compuesto de la siguiente forma:
número destino, número origen, duración (en segundos), fecha (ISO 8601 en UTC)

*Output*

- Una factura con el siguiente formato e información que se imprime por STDOUT en formato JSON:

```json
    {
    "user":{
        "address":"Avenida siempre viva", // Dirección de usuario
        "name":"Juan Sanchez", // Nombre del usuario
        "phone_number":"+54911111111" // Número del usuario
    },
    "calls":[
        {
            "phone_number":"+19911111113", // Número destino
            "duration":60, // Duración de la llamada
            "timestamp":"2020-08-09T04:45:25Z", // Fecha y hora de la llamada
            "amount":60.00 // Costo de llamada
        },
        {
            "phone_number":"+54911111112",
            "duration":120,
            "timestamp":"2020-08-09T03:45:25Z",
            "amount": 2.5
        },
        {
            "phone_number":"+54911111113",
            "duration":122,
            "timestamp":"2020-01-10T15:45:25Z",
            "amount":0.0
        }
    ],
    "total_international_seconds":60, // Total segundos internacionales
    "total_national_seconds":242, // Total segundos nacionales
    "total_friends_seconds":122, // Total segundos a amigos
    "total":62.50 // Total de la factura
    }
```

*Usuario*: Es el usuario de la línea telefónica, consta de un nombre, una dirección, un
número de teléfono y una lista de amigos (lista de números de teléfonos)

*Nros de teléfono*: Formato +549XXXXXXXXXX, los primeros 2 dígitos luego del +
representan el país: en este caso +54 es Argentina

*Llamadas*:
Están representadas por:
- *Orígen*: Número de teléfono de quien llama
- *Destino*: Número al cual se llamó
- *Fecha*: En que se realizó la llamada, formato AAAA-MM-DD HH:MM:SS -
Duración: Cuánto tiempo duró esa llamada, expresada en segundos
Pueden ser de 3 tipos: Nacionales, Internacionales y Amigos
Nacionales: Origen y Destino deben tener del mismo código de país
Internacionales: Origen y Destino deben tener distinto código de país
Amigos: Las primeras 10 llamadas hacia un amigo son sin costo

*Tarifas*:
- Nacionales ($2.5 por llamada)
- Internacionales ($0.75 x segundo)
- Amigos (Gratis hasta 10 llamadas)

## Instalacion

```bash
go mod tidy
```

```bash
go build
```

## Como usarlo

Para correr el programa se debe ejecutar el siguiente comando por terminal

```bash
go run main.go <phoneNumber> <2015-01-01>/<2023-08-01> <nombre_del_archivo.csv>
```

A continuacion se deja un ejemplo con un caso de prueba

```bash
go run main.go +5491167920930 2015-01-01/2023-08-01 calls.csv
```

## Correr los Tests

```bash
go test ./...
```
