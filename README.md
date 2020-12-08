# Operación Fuego de Quasar
<img src=https://media.moddb.com/images/members/1/118/117287/QUASAR_WIP02.png align="right" width="534px" heigth="268px">
Han Solo ha sido recientemente nombrado General de la Alianza Rebelde y busca dar un gran golpe contra el Imperio Galáctico para reavivar la llama de la resistencia.

El servicio de inteligencia rebelde ha detectado un llamado de auxilio de una nave portacarga imperial a la deriva en un campo de asteroides. El manifiesto de la nave es ultra clasificado, pero se rumorea que transporta raciones y armamento para una legión entera.

Esta aplicación se encargará de retornar la fuente y el contenido del mensaje de auxilio a partir de tres satélites configurados que permitirán triangular la posición.

Para poder localizar la posición de la nave portacarga se utilizó la ecuación de distancia entre dos puntos. Conociendo los puntos donde se encuentran los satélites configurados
y desconociendo el punto donde se encuentra la nave portacarga podemos obtener un sistema de ecuaciones lineales de dos variables y luego resolverlo para obtener el punto solución donde se encuentra la nave portacarga.
Para mayor detalle acerca de las fórmulas utilizadas, se puede consultar el siguiente [documento](https://docs.google.com/document/d/1PiRTU7AOkIDrYFgIyB2mdwVcFwVOUy7tlasMTPcu1gc/edit?usp=sharing).

## Acerca de la solución
API RESTful instanciada en [Heroku](https://www.heroku.com/) para obtener la fuente y contenido del mensaje de auxilio.

La solución consta de los siguientes endpoints:

[POST] `/api/topsecret` para obtener la ubicación de la nave y el mensaje que emite a partir de la información de los satélites en el payload.
[POST] `/api/topsecret_split/{satellite_name}` permite cargar la información de cada satélite por separado en el payload.                                                
[GET ]  `/api/topsecret_split` para obtener la ubicación de la nave y el mensaje que emite si es posible y si todos los satélites fueron cargados.


## Documentación

Disponible en [https://fire-operation-api.herokuapp.com/api/doc/index.html](https://fire-operation-api.herokuapp.com/api/doc/index.html) con todas las firmas 
y definiciones de los endpoints con sus correspondientes códigos de respuesta. Pueden realizarse pruebas a las API's desde el link.

## Requerimientos

Se requiere tener instalado:

*   Golang 1.13.5 


## Buildeo

Para realizar el build ejecutar en consola los siguientes comandos:

```
go mod tidy
go run main.go
```

Para acceder a los endpoints de forma local utilizar como url base: http://localhost:8080/api/

## Testing

Para correr los test unitarios correr el siguiente comando:

```
go test -v -coverpkg ./... ./...
```


## Swagger Docs

Para actualizar la documentación de swagger ejecutar:

```
go get -u github.com/swaggo/swag/cmd/swag
swag init
```

## Utilización de los endpoints

Para llamar al endpoint [POST] `/api/topsecret` enviar un payload con el siguiente formato:

```
{
    "satellites": [
		{
		"name": "kenobi",
		"distance": 485.41,
		"message": ["este", "", "", "mensaje", ""]
		},
		
		{
		"name": "skywalker",
		"distance": 265.75,
		"message": ["", "es", "", "", "secreto"]
		},
		{
		"name": "sato",
		"distance": 600.52,
		"message": ["este", "", "un", "", ""]
		}
	]

}
```
En caso de ser exitosa la petición se obtendrá una respuesta con el siguiente formato:
```
{
    "position": {
        "x": -100,
        "y": 75
    },
    "message": "este es un mensaje secreto"
}
```

Para llamar al endpoint [POST] `/api/topsecretsplit/{satellite_name}` enviar un payload con el siguiente formato:

```
{
	"distance": 266.1,
	"message": ["este", "", "un", "", ""]
}
```

Para llamar al endpoint [GET] `/api/topsecretsplit` deberán enviarse previamente los tres satélites por separado para obtener una respuesta con el siguiente formato:

```
{
    "position": {
        "x": -100,
        "y": 75
    },
    "message": "este es un mensaje secreto"
}
```
