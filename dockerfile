# Usar una imagen base de Go
# docker run -p 8080:8080 my-go-app-image
FROM golang:latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el código fuente de tu aplicación y los archivos necesarios
COPY . .

# Compilar la aplicación
RUN go build -o Application

# Exponer el puerto en el que la aplicación se ejecutará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./Application"]