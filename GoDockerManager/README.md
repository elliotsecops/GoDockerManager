# Docker Container Manager

Docker Container Manager es una herramienta de línea de comandos diseñada para simplificar la gestión de contenedores Docker. Proporciona un conjunto de comandos para listar, iniciar, detener, ver registros, inspeccionar y monitorear estadísticas de contenedores. Esta herramienta es ideal para desarrolladores y administradores de sistemas que necesitan una manera rápida y fácil de gestionar contenedores Docker desde la línea de comandos.

## Tabla de Contenidos

- [Características](#características)
- [Instalación](#instalación)
  - [Prerrequisitos](#prerrequisitos)
  - [Usando el Script de Instalación](#usando-el-script-de-instalación)
- [Uso](#uso)
  - [Listar Contenedores](#listar-contenedores)
  - [Iniciar un Contenedor](#iniciar-un-contenedor)
  - [Detener un Contenedor](#detener-un-contenedor)
  - [Ver Registros de Contenedor](#ver-registros-de-contenedor)
  - [Inspeccionar Contenedor](#inspeccionar-contenedor)
  - [Monitorear Estadísticas de Contenedor](#monitorear-estadísticas-de-contenedor)
- [Desinstalación](#desinstalación)
- [Contribuir](#contribuir)
- [Licencia](#licencia)

## Características

- **Listar Contenedores**: Muestra una lista de todos los contenedores Docker.
- **Iniciar un Contenedor**: Inicia un contenedor Docker específico.
- **Detener un Contenedor**: Detiene un contenedor Docker específico.
- **Ver Registros de Contenedor**: Ve los registros de un contenedor Docker específico.
- **Inspeccionar Contenedor**: Obtiene información detallada sobre un contenedor Docker específico.
- **Monitorear Estadísticas de Contenedor**: Monitorea el uso de recursos de un contenedor Docker específico.

## Instalación

### Prerrequisitos

Antes de instalar Docker Container Manager, asegúrate de tener los siguientes prerrequisitos instalados en tu sistema:

- **Go (Golang)**: Versión 1.16 o superior.
- **Docker**: Asegúrate de que Docker esté instalado y en funcionamiento en tu sistema.

### Usando el Script de Instalación

Para instalar Docker Container Manager, sigue estos pasos:

1. **Clonar el Repositorio**:
   ```sh
   git clone https://github.com/elliotsecops/GoDockerManager.git
   cd GoDockerManager
   ```

2. **Ejecutar el Script de Instalación**:
   ```sh
   ./install.sh
   ```

   El script construirá el proyecto y moverá el ejecutable a `/usr/local/bin/`, haciéndolo disponible en todo el sistema.

## Uso

Una vez instalado, puedes usar el comando `docker-container-manager` para gestionar tus contenedores Docker. A continuación se presentan los comandos disponibles y sus descripciones.

### Listar Contenedores

Para listar todos los contenedores Docker, usa el siguiente comando:

```sh
docker-container-manager list
```

### Iniciar un Contenedor

Para iniciar un contenedor Docker específico, usa el siguiente comando:

```sh
docker-container-manager start <containerID>
```

Reemplaza `<containerID>` con el ID del contenedor que deseas iniciar.

### Detener un Contenedor

Para detener un contenedor Docker específico, usa el siguiente comando:

```sh
docker-container-manager stop <containerID>
```

Reemplaza `<containerID>` con el ID del contenedor que deseas detener.

### Ver Registros de Contenedor

Para ver los registros de un contenedor Docker específico, usa el siguiente comando:

```sh
docker-container-manager logs <containerID>
```

Reemplaza `<containerID>` con el ID del contenedor cuyos registros deseas ver.

### Inspeccionar Contenedor

Para obtener información detallada sobre un contenedor Docker específico, usa el siguiente comando:

```sh
docker-container-manager inspect <containerID>
```

Reemplaza `<containerID>` con el ID del contenedor que deseas inspeccionar.

### Monitorear Estadísticas de Contenedor

Para monitorear el uso de recursos de un contenedor Docker específico, usa el siguiente comando:

```sh
docker-container-manager stats <containerID>
```

Reemplaza `<containerID>` con el ID del contenedor que deseas monitorear.

## Desinstalación

Para desinstalar Docker Container Manager, sigue estos pasos:

1. **Ejecutar el Script de Desinstalación**:
   ```sh
   ./uninstall.sh
   ```

   El script eliminará el ejecutable de `/usr/local/bin/`.

## Contribuir

¡Las contribuciones son bienvenidas! Si tienes alguna idea, sugerencia o informe de errores, por favor abre un issue o envía un pull request. Para cambios importantes, por favor discute tus ideas en un issue antes de hacer cualquier cambio.

### Cómo Contribuir

1. Haz un fork del repositorio.
2. Crea una nueva rama (`git checkout -b feature/tu-nombre-de-característica`).
3. Realiza tus cambios y confírmalos (`git commit -m 'Agregar alguna característica'`).
4. Envía los cambios a la rama (`git push origin feature/tu-nombre-de-característica`).
5. Abre un pull request.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo [LICENSE](LICENSE) para más detalles.

¡Gracias por usar Docker Container Manager! Si tienes alguna pregunta o necesitas más ayuda, no dudes en ponerte en contacto.


--- 

# Docker Container Manager

Docker Container Manager is a CLI tool designed to simplify the management of Docker containers. It provides a set of commands to list, start, stop, view logs, inspect, and monitor container statistics. This tool is ideal for developers and system administrators who need a quick and easy way to manage Docker containers from the command line.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Using the Installation Script](#using-the-installation-script)
- [Usage](#usage)
  - [List Containers](#list-containers)
  - [Start a Container](#start-a-container)
  - [Stop a Container](#stop-a-container)
  - [View Container Logs](#view-container-logs)
  - [Inspect Container](#inspect-container)
  - [Monitor Container Statistics](#monitor-container-statistics)
- [Uninstallation](#uninstallation)
- [Contributing](#contributing)
- [License](#license)

## Features

- **List Containers**: Display a list of all Docker containers.
- **Start a Container**: Start a specific Docker container.
- **Stop a Container**: Stop a specific Docker container.
- **View Container Logs**: View the logs of a specific Docker container.
- **Inspect Container**: Get detailed information about a specific Docker container.
- **Monitor Container Statistics**: Monitor the resource usage of a specific Docker container.

## Installation

### Prerequisites

Before installing Docker Container Manager, ensure you have the following prerequisites installed on your system:

- **Go (Golang)**: Version 1.16 or higher.
- **Docker**: Ensure Docker is installed and running on your system.

### Using the Installation Script

To install Docker Container Manager, follow these steps:

1. **Clone the Repository**:
   ```sh
   git clone https://github.com/elliotsecops/GoDockerManager.git
   cd GoDockerManager
   ```

2. **Run the Installation Script**:
   ```sh
   ./install.sh
   ```

   The script will build the project and move the executable to `/usr/local/bin/`, making it available system-wide.

## Usage

Once installed, you can use the `docker-container-manager` command to manage your Docker containers. Below are the available commands and their descriptions.

### List Containers

To list all Docker containers, use the following command:

```sh
docker-container-manager list
```

### Start a Container

To start a specific Docker container, use the following command:

```sh
docker-container-manager start <containerID>
```

Replace `<containerID>` with the ID of the container you want to start.

### Stop a Container

To stop a specific Docker container, use the following command:

```sh
docker-container-manager stop <containerID>
```

Replace `<containerID>` with the ID of the container you want to stop.

### View Container Logs

To view the logs of a specific Docker container, use the following command:

```sh
docker-container-manager logs <containerID>
```

Replace `<containerID>` with the ID of the container you want to view logs for.

### Inspect Container

To get detailed information about a specific Docker container, use the following command:

```sh
docker-container-manager inspect <containerID>
```

Replace `<containerID>` with the ID of the container you want to inspect.

### Monitor Container Statistics

To monitor the resource usage of a specific Docker container, use the following command:

```sh
docker-container-manager stats <containerID>
```

Replace `<containerID>` with the ID of the container you want to monitor.

## Uninstallation

To uninstall Docker Container Manager, follow these steps:

1. **Run the Uninstallation Script**:
   ```sh
   ./uninstall.sh
   ```

   The script will remove the executable from `/usr/local/bin/`.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request. For major changes, please discuss your ideas in an issue before making any changes.

### How to Contribute

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature-name`).
3. Make your changes and commit them (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature-name`).
5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

Thank you for using Docker Container Manager! If you have any questions or need further assistance, feel free to reach out.
```

### Conclusion

This detailed `README.md` file provides users with clear instructions on how to install, use, and uninstall your Docker Container Manager tool. It also includes sections for contributing and licensing, making it easier for others to contribute to your project and understand its terms of use.