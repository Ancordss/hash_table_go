import hashlib

# Función para generar la clave hash
def generar_clave_hash(fila, columna):
    clave = str(fila) + ',' + str(columna)
    return hashlib.sha256(clave.encode()).hexdigest()

# Función para procesar el texto y generar la tabla hash
def procesar_texto(texto):
    tabla_hash = {}

    # Dividir el texto en líneas
    lineas = texto.split('\n')

    # Recorrer cada línea y obtener los tokens
    for fila, linea in enumerate(lineas):
        tokens = linea.split()

        # Recorrer cada token y generar la clave hash
        for columna, token in enumerate(tokens):
            clave_hash = generar_clave_hash(fila, columna)
            tabla_hash[clave_hash] = token

    return tabla_hash

# Ejemplo de uso
texto = """
int suma = 0;
suma = 54 + 20;
"""

tabla = procesar_texto(texto)

# Imprimir la tabla hash
for clave, valor in tabla.items():
    print(clave, ':', valor)
