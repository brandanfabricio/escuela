### Diagrama de Flujo de Datos (DFD) y Diccionario de Datos

---

#### **Diagrama de Flujo de Datos (DFD)**

##### **Nivel 0 (Diagrama Contextual)**

**Procesos Principales:**
1. Interactuar con el catálogo de productos.
2. Gestionar productos (administradores).

**Interacciones:**
- **Usuario Cliente**: Navega, filtra y selecciona productos del catálogo.
- **Administrador**: Gestiona el catálogo (agregar/editar productos).
- **WhatsApp**: Medio de contacto para concretar compras.

**Flujo Principal:**
Usuarios interactúan con el sistema para consultar productos y contactar al vendedor.

---

##### **Nivel 1 (Desglose de Procesos Principales)**

1. **Consultar Catálogo**:
    - Flujo desde el cliente al catálogo.
    - Proporciona una lista de productos.
2. **Gestionar Productos**:
    - Flujo desde administradores para agregar/editar productos.
3. **Contactar Vendedor**:
    - Flujo desde el cliente hacia WhatsApp para gestionar compras.

---

##### **Nivel 2 (Desglose Detallado de Procesos)**

1.1. **Ver Productos Destacados**:
    - Muestra los 5 productos destacados en el home.
1.2. **Buscar y Filtrar Productos**:
    - Permite buscar por nombre o filtrar por categoría.

2.1. **Agregar Producto**:
    - Permite a los administradores cargar nuevos productos.
2.2. **Editar Producto**:
    - Permite a los administradores modificar información de productos existentes.

3.1. **Enviar Mensaje por WhatsApp**:
    - Genera un enlace dinámico con los datos del producto seleccionado.

---

#### **Diccionario de Datos**

1. **Entidad: Producto**
    - **Descripción**: Representa un producto disponible en el catálogo.
    - **Atributos**:
        - `IDProducto`: Identificador único del producto.
        - `NombreProducto`: Nombre del producto.
        - `Precio`: Precio del producto.
        - `Categoría`: Categoría a la que pertenece.
        - `Descripción`: Información adicional del producto.
        - `Imagen`: URL o archivo de imagen del producto.

2. **Entidad: Usuario Cliente**
    - **Descripción**: Persona que interactúa con el catálogo.
    - **Atributos**:
        - `IDCliente`: Identificador único (opcional si no hay login).
        - `Acción`: Acción realizada por el cliente (ver, buscar, filtrar, etc.).

3. **Entidad: Administrador**
    - **Descripción**: Persona encargada de gestionar el catálogo.
    - **Atributos**:
        - `IDAdministrador`: Identificador único.
        - `Credenciales`: Usuario y contraseña para autenticación.

4. **Proceso: Consultar Catálogo**
    - **Descripción**: Proceso que permite a los clientes interactuar con el catálogo.
    - **Entradas**:
        - Solicitud de lista de productos.
    - **Salidas**:
        - Listado de productos (destacados o filtrados).

5. **Proceso: Gestionar Productos**
    - **Descripción**: Proceso para agregar y editar productos.
    - **Entradas**:
        - Datos del producto.
    - **Salidas**:
        - Producto agregado o actualizado.

6. **Proceso: Contactar Vendedor**
    - **Descripción**: Genera un enlace dinámico para abrir WhatsApp con un mensaje predefinido.
    - **Entradas**:
        - Selección del producto por el cliente.
    - **Salidas**:
        - Enlace de WhatsApp.

7. **Flujo de Datos: Solicitud de Producto**
    - **Origen**: Usuario Cliente.
    - **Destino**: Sistema.
    - **Descripción**: Solicitud para ver productos destacados o filtrados.

8. **Flujo de Datos: Gestión de Producto**
    - **Origen**: Administrador.
    - **Destino**: Sistema.
    - **Descripción**: Datos del producto para agregar o editar.

9. **Flujo de Datos: Enlace WhatsApp**
    - **Origen**: Sistema.
    - **Destino**: Cliente (WhatsApp).
    - **Descripción**: Enlace generado dinámicamente con los datos del producto seleccionado.

---

Este DFD y Diccionario de Datos está diseñado para modelar y comprender el flujo de información y las interacciones en el sistema estilo catálogo.

