# gambitUser_Lambda

Lambda hecha en Go para proyecto en AWS.

El objetivo de esta aplicación pensada para correr como una función Lambda en AWS, es recibir el evento de Cognito cuando un usuario autentica su correo electrónico ingresando el código de verificación.

En ese momento entra en funcionamiento esta aplicación hecha en Go. Recibimos el evento con la información del usuario y lo guardamos en una base de datos MySQL también alojada en AWS.

## Requisitos

-   Go 1.20
-   Cuenta en AWS
-   Configuración de un pool de usuarios en Cognito
-   Base de datos MySQL en AWS
-   Configurar un secret en Secrets Manager en AWS con la información de la base de datos
-   Crear una función Lambda en AWS
-   Crear una variable de entorno con el nombre del secreto creado en Secrets Manager
-   Crear un trigger en la función Lambda para que reciba los eventos de Cognito

## Instalación

-   Clonar el repositorio
-   Ejecutar `go get` para descargar las dependencias
-   Ejecutar `go build` para compilar el proyecto (recordar que el binario debe llamarse `main` y que lambda solo acepta binarios compilados para Linux, en el archivo update.sh se encuentra el comando para compilar el binario para Linux correctamente `env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o . main.go` )
-   Crear un archivo .zip con el binario generado
-   Subir el archivo .zip a la función Lambda en AWS

Después de esto, la función Lambda estará lista para recibir los eventos de Cognito y guardar los datos en la base de datos MySQL. Para probarlo, se puede crear un usuario en el pool de usuarios de Cognito, simular un registro con un correo electrónico válido y verificar el correo electrónico. En ese momento se ejecutará la función Lambda, para verificar, puedes ir a la base de datos para confirmar que el registro se agregó correctamente.
