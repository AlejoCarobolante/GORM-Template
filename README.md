#!/bin/bash

# 🏗️ Clean Architecture (Go + GORM + Gin)

Este documento describe la arquitectura de capas utilizada en el proyecto, siguiendo los principios de **Clean Architecture** (Arquitectura Limpia).

El principio fundamental es la **Regla de Dependencia**: las dependencias siempre deben apuntar hacia el interior. Las capas exteriores dependen de las capas interiores, pero nunca al revés.

---

## I. Capas del Sistema

### 1. 🧠 Capa de Dominio (\`domain\`)
| Aspecto | Descripción |
| :--- | :--- |
| **Rol** | Contiene las reglas de negocio, las estructuras de datos principales y los contratos (interfaces) que definen las interacciones. Es el núcleo puro de la aplicación. |
| **Contenido** | \`Entity1\` (Modelo de la entidad). \`Entity1Repository\` (Interfaz que define el CRUD: \`Create\`, \`Fetch\`, \`Update\`, \`Delete\`). |
| **Dependencias** | Ninguna dependencia de otras capas o frameworks. |

### 2. 💡 Capa de Casos de Uso (\`usecase\`)
| Aspecto | Descripción |
| :--- | :--- |
| **Rol** | Contiene la lógica de aplicación específica y orquesta el flujo de datos. Implementa los contratos definidos en el Dominio. |
| **Contenido** | \`Entity1UseCase\` (Implementación de la interfaz \`domain.Entity1Repository\`). |
| **Dependencias** | Depende del **Dominio** (para usar las entidades). **Nota Importante:** En esta implementación, también accede directamente al framework de persistencia (\`bootstrap.DB\`/GORM), lo cual es una simplificación común, pero técnicamente coloca la implementación del repositorio en esta capa. |

### 3. 🌐 Capa de Adaptadores (\`controller\` y \`route\`)
| Aspecto | Descripción |
| :--- | :--- |
| **Rol** | Adapta la entrada y salida de la aplicación al mundo exterior (HTTP/JSON). |
| **Contenido** | \`Entity1Controller\` (Maneja el binding de Gin, validación básica y formato de respuesta HTTP). \`route.go\` (Define los *endpoints* Gin e inyecta las dependencias). |
| **Dependencias** | Depende del **Dominio** (para la interfaz) y de la capa de **Casos de Uso** (para la implementación concreta a inyectar). |

---

![Diagrama de Capas de Clean Architecture](src/capas.jpg)

## II. Flujo Típico de una Petición (Ejemplo: \`Create\`)

Este es un resumen del recorrido de una petición HTTP a través de las capas:

| Paso | Capa de Origen | Acción Principal | Dependencia/Llamada |
| :--- | :--- | :--- | :--- |
| **1.** | **Route** | Mapea el método y el *path* (e.g., \`POST /entity1\`) al *handler* del Controller. | Inyección de \`Entity1UseCase\` en \`Entity1Controller\`. |
| **2.** | **Controller** | Recibe la petición. Hace *binding* del JSON a \`domain.Entity1\` y realiza validaciones de entrada. | Depende del **Dominio** (Estructuras/Tipos). |
| **3.** | **Controller** | Llama al método \`Create\` de la interfaz \`Entity1Repository\`. | Llama al método de la interfaz del **Dominio**. |
| **4.** | **UseCase** | Se ejecuta la implementación de \`Create\` dentro de \`Entity1UseCase\`. | Depende del **UseCase** (Lógica de Aplicación). |
| **5.** | **UseCase** | Ejecuta la operación de persistencia utilizando la conexión directa a la base de datos (\`bootstrap.DB\` y GORM). | Depende del **Framework/DB** (capa externa). |
| **6.** | **Controller** | Recibe el resultado/error del UseCase. Formatea y envía la respuesta HTTP (e.g., Status 200 OK o 500 Internal Server Error). | Retorno al cliente. |
EOF
