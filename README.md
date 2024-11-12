# traffic_light
Comprender el uso de semáforos para la sincronización de procesos

# Laboratorio de Sistemas Operativos

## Objetivo

Comprender el uso de semáforos para la sincronización de procesos.

## Detalle

Usando el lenguaje de programación Go, se han creado dos procesos, P1 y P2. El proceso P1 imprime cien unos (1) en la pantalla y el proceso P2 imprime cien ceros (0) en la pantalla. La sincronización mediante semáforos asegura que en la pantalla aparezcan los números intercalados:

101010101010101010101010101010101010...

## Estructura del Proyecto

- `main.go`: Contiene el código principal del proyecto.

## Requisitos

- Go 1.16 o superior

## Ejecución

Para ejecutar el proyecto, sigue estos pasos:

1. Clona el repositorio:
   ```sh
   git https://github.com/DanielVelandia2407/traffic_light.git
   cd traffic_light

2. Ejecuta el programa:
   go run main.go

Explicación del Código
El código utiliza canales con buffer de tamaño 1 para actuar como semáforos y sincronizar dos goroutines que imprimen "1" y "0" alternativamente. Un WaitGroup asegura que el programa principal espere a que ambas goroutines terminen antes de finalizar.
Secciones del Código
Inicialización de Semáforos: Se crean dos canales con buffer de tamaño 1 (p1Chan y p2Chan) para actuar como semáforos.
Inicialización del Semáforo de P1: Se inicializa el semáforo de P1 para que comience primero.
WaitGroup: Se crea un WaitGroup para esperar a que ambos goroutines terminen.
Goroutine para P1: Imprime "1" y libera el semáforo de P2.
Goroutine para P2: Imprime "0" y libera el semáforo de P1.
Esperar a que Terminen las Goroutines: El WaitGroup espera a que ambas goroutines terminen antes de finalizar el programa.
