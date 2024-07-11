# Etapa de construcción
FROM golang:1.22.5 AS builder

# Instalar dependencias para Oracle Instant Client 
RUN apt-get update && apt-get install -y libaio1 wget unzip

# Descargar e instalar Oracle Instant Client
RUN wget -O /tmp/instantclient-basic-linux-x64.zip https://download.oracle.com/otn_software/linux/instantclient/193000/instantclient-basic-linux.x64-19.3.0.0.0dbru.zip \
    && mkdir -p /usr/lib/oracle \
    && unzip /tmp/instantclient-basic-linux-x64.zip -d /usr/lib/oracle \
    && rm /tmp/instantclient-basic-linux-x64.zip

# Configurar Oracle Instant Client
RUN echo "/usr/lib/oracle/instantclient_19_3" > /etc/ld.so.conf.d/oracle-instantclient.conf \
    && ldconfig

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Aprovechar el chache copiando primero los archivos go.mod y go.sum
COPY go.mod go.sum ./

# Descargar dependencias 
RUN go mod download && go mod verify 

COPY . .

# Eliminar y agregar dependencias y actualizar los archivos go.mod y go.sum
RUN go mod tidy

# Compilar la aplicación
RUN go build -o storeproc

# Etapa de ejecución
FROM ubuntu:22.04

WORKDIR /app

# Instalar las dependencias necesarias para ejecutar el binario Go y Oracle Instant Client
RUN apt-get update && apt-get install -y libaio1 && rm -rf /var/lib/apt/lists/*

# Copiar Oracle Instant Client desde la etapa de construcción
COPY --from=builder /usr/lib/oracle /usr/lib/oracle

# Configurar las variables de entorno necesarias para Oracle Instant Client
ENV LD_LIBRARY_PATH /usr/lib/oracle/instantclient_19_3
ENV PATH /usr/lib/oracle/instantclient_19_3:$PATH

# Copiar el binario compilado desde la imagen de compilación
COPY --from=builder /app/storeproc .

# Copiar el archivo .env
COPY .env .env

EXPOSE 8080

# Ejecutar la aplicación
CMD ["./storeproc"]