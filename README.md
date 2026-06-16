# Patrones de Diseño (Gang of Four)

## Introducción

Los patrones de diseño documentados por el *Gang of Four* (GoF) no fueron inventados, sino **descubiertos**. A partir de la observación de código bien escrito en proyectos diversos, sus autores identificaron que ciertos problemas recurrentes se resolvían sistemáticamente de la misma manera.

El catálogo constituye, en esencia, un repertorio de soluciones probadas. Cada patrón posee un nombre, describe el problema que resuelve y propone una estructura de clases y objetos para abordarlo.

Los patrones se organizan en tres categorías.

## 1. Patrones creacionales

Resuelven el problema de **cómo se crean los objetos**. Cuando el código instancia objetos directamente mediante `new`, queda acoplado a clases concretas y resulta difícil de modificar. Estos patrones desacoplan la creación del uso.

| Patrón | Propósito |
| --- | --- |
| Abstract Factory | Crea familias de objetos relacionados sin especificar sus clases concretas. |
| Builder | Separa la construcción de un objeto complejo de su representación. |
| Factory Method | Delega en las subclases la decisión de qué clase instanciar. |
| Prototype | Crea nuevos objetos a partir de la clonación de una instancia existente. |
| Singleton | Garantiza que una clase tenga una única instancia en todo el sistema. |

## 2. Patrones estructurales

Resuelven **cómo se componen clases y objetos** para formar estructuras de mayor tamaño. La cuestión central es cómo ensamblar las distintas piezas sin que queden rígidamente acopladas entre sí.

| Patrón | Propósito |
| --- | --- |
| Adapter | Traduce una interfaz en otra compatible con lo que el cliente espera. |
| Bridge | Separa una abstracción de su implementación para que varíen de forma independiente. |
| Composite | Compone objetos en estructuras de árbol y los trata de manera uniforme. |
| Decorator | Añade responsabilidades a un objeto de forma dinámica, sin modificar su clase. |
| Facade | Ofrece una interfaz unificada que simplifica el acceso a un sistema complejo. |
| Flyweight | Comparte estado común entre múltiples objetos para reducir el consumo de memoria. |
| Proxy | Provee un sustituto que controla el acceso a otro objeto. |

## 3. Patrones de comportamiento

Son los más numerosos y resuelven **cómo se comunican los objetos** y cómo se distribuyen las responsabilidades. Su foco no es la estructura, sino el flujo de control y los algoritmos.

| Patrón | Propósito |
| --- | --- |
| Chain of Responsibility | Pasa una petición a lo largo de una cadena de receptores potenciales. |
| Command | Encapsula una petición como un objeto. |
| Interpreter | Define una gramática y un intérprete para un lenguaje determinado. |
| Iterator | Permite recorrer los elementos de una colección sin exponer su representación interna. |
| Mediator | Centraliza la comunicación entre objetos para reducir las dependencias directas. |
| Memento | Captura y restaura el estado interno de un objeto. |
| Observer | Notifica automáticamente a los objetos dependientes ante un cambio de estado. |
| State | Modifica el comportamiento de un objeto cuando cambia su estado interno. |
| Strategy | Permite intercambiar algoritmos en tiempo de ejecución. |
| Template Method | Define el esqueleto de un algoritmo y delega pasos a las subclases. |
| Visitor | Separa un algoritmo de la estructura de objetos sobre la que opera. |

## Referencia

Gamma, E., Helm, R., Johnson, R. y Vlissides, J. (1994). *Design Patterns: Elements of Reusable Object-Oriented Software*. Addison-Wesley.
