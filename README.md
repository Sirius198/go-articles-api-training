# Golang API Test

El alcance de esta API es crear un grupo de endpoints usando el codigo base disponible 
que satisfagan los requerimientos mencionados a continuacion. 

Debera contar con los siguientes conocimientos:
- [Golang](https://go.dev/)
- [Go Swagger](https://goswagger.io/)
- [PostgresSQL](https://www.postgresql.org/)
- Restfull APIs


## Alcance del Proyecto
El objectivo principal es crear una API basica que permita dar soporte a un carrito de compras, a continuacion se describen los requerimientos y las iteracciones con la API:

### Entidades minimas que deben existir
*(Hay libertad de extenderlo a preferencia del candidato)*
- **Articulos**: Debe como minimo tener los campos de `ID`, `Disponible(FALSE|TRUE)`, `Precio`, `Nombre`
- **Promociones**: Debe como minimo tener los campos de `ID`, `Usada(FALSE|TRUE)`, `Codigo`, `Nombre`
- **Usuarios**: Debe como minimo tener los campos de `ID`, `Nombre`, `Activo(TRUE|FALSE)`, `Tipo(CLIENTE|ADMIN)`
- **Ordenes**: Debe como minimo tener los campos de `ID`, `NRO_ORDEN`, `ID_USUARIO`, `ID_PROMOCION`, `STATUS(REVISION|COMPLETADA)`, `TOTAL`, `TOTAL_DESCUENTO`
- **Articulos_Ordenes**: Debe como minimo tener los campos de `ID`, `ID_ORDEN`, `ID_ARTICULO`, `PRICE`

Puede crear mas Entidades y mas cmapos a cualquiera de las entidades mencionadas como minimo, la definiciones anterior es lo minimo que se espera

### Requerimientos de usuarios
El usuario cuenta con dos tipos de usuarios `Clientes` y `Administrados`

**La iteraciones a continuacion es de usuario tipo Cliente**:
1. El Usuario debe poder LISTAR los Articulos disponibles, y filtrarlos por nombre y disponibilidad de articulos
2. El Usuario debe poder LEER un Articulo Individual
3. El Usuario NO debe tener accesso a CREAR, MODIFICAR o BORRAR Articulos
4. El Usuario debe poder CREAR ordenes
5. El Usuario debe poder CREAR, MODIFICAR y BORRAR solo las ordenes que le pertencen a su usuario
6. El Usuario no debe tener accesso a CREAR, LEER, MODIFICAR or BORRAR ordenes de otros usuarios
7. El Usuario debe poder agregar promociones a las ordenes
8. El Usuario no debe tener accesso a CREAR, LEER, MODIFICAR or BORRAR ordenes de otros usuarios, solo puede usar las promociones al crear ordenes usando su codigo
9. El Usuario debe poder MODIFICAR su USUARIO unicamente
10. El Usuario no debe tener accesso a CREAR o BORRAR usuarios y no debe tener accesso a MODIFICAR o LEER otros Usuarios
11. 

**La iteraciones a continuacion is de usuario tipo Administrativo**:
1. El Usuario debe tener accesso a CREAR, LEER, MODIFICAR or BORRAR Usuarios
2. El Usuario debe tener accesso a CREAR, LEER, MODIFICAR or BORRAR Articulos
3. El Usuario debe tener accesso a CREAR, LEER, MODIFICAR or BORRAR Promociones
4. El Usuario debe tener accesso a LEER, MODIFICAR or BORRAR Ordenes
5. El Usuario NO debe tener accesso a CREAR Ordenes


## Consideraciones de la API para proteger su integridad
Al Crear un Orden:
- Si el Articulo no esta disponible(`Disponible=false`), la API debe retornar error con el mensaje indicado
- Si la Orden incluye promocion y la promocion esta usada(`Usada=true`), la API debe retornar error con el mensaje indicado
- Si el Usuario esta esta inactivo(`Activo=false`), la API debe retornar error con el mensaje indicado
- Cuando la orden se crea, el Status por defecto es `REVISION`, la Orden solo puede tener `Status=COMPLETADA` a travez de una actualizado

Al Interactuar con Articulos:
- Por defecto el API debe retornar Articulos Activos (debe ignorar `Disponible=false`)
- Si el usuario intenta borrar un Articulo asociado a una Orden, la API debe retornar error con el mensaje indicado
- Si el usuario intenta borrar una Promocion asociada a una Orden, la API debe retornar error con el mensaje indicado
- Si el usuario intenta borrar un Usuario asociado a una Orden, la API debe retornar error con el mensaje indicado

### Requerimientos tecnicos
1. Cada endpoint creado debe leer el ID del usuario de los headers, el header esperado se llamara `X-USERID`
2. Cada endpoint debe validar si el usuario tienen acceso, por cuestiones de tiempo limitaremos la validacion al tipo de usuario
3. Cada endpoint debe hacer uso correcto de los verbos HTTP, and del StatusCode para la respuesta al usuario.
4. El Header se espera que venga codificado en Base64, y los endpoints debe decodificarlo
5. Los endpoints debe ofrecer filtrado basico, e.g.: usuarios pueden filtrar Articulos por nombre y disponibilidad


### Punto Extras
- Hacer uso de middleware donde se pueda
- Hacer uso de [Golang Context](https://pkg.go.dev/context) package
- Hacer uso de Database connection pool


## Expectativas de la Solucion
- Instruciones para compilar el Binary
- Instruciones para correr pruebas (e.g. go test ...)
- Varios endpoints que den soporte a operaciones CRUD (Create, Retrieve, Update, Delete)
- Un estructura de Modelo para schema de base de datos
- Uso appropiados de librerias de Golang para conectarse a una PostgresSQL Base de Dato
- Uso appropiados de variables de entorno que permita configurar el codigo a diferentes ambientes
- Codigo limpio y bien indexado
- Uso correcto de taxonomia de API 
- Estructura de projecto que incluya propias librerias que facilite reusabilidad
