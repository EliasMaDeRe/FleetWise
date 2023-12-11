async function metodoPOST(peticionUrl, datosBody) {
    try {
        const resultado = await fetch(peticionUrl, {
            method: "POST",
            body: JSON.stringify(datosBody),
        });

        const respuesta = await resultado.json();

        return respuesta;
    } catch (error) {
        throw new Error('Error al realizar la solicitud POST: ' + error.message);
    }
}


async function metodoGET(peticionUrl) {
    try {
        const resultado = await fetch(peticionUrl, {
            method: "GET"
        });

        const respuesta = await resultado.json();

        return respuesta;
    } catch (error) {
        throw new Error('Error al obtener los registros: ' + error.message);
    }
}

async function metodoPUT(peticionUrl) {
    try {
        const resultado = await fetch(peticionUrl, {
            method: "PUT",
            body: JSON.stringify(filtros),
        });

        const respuesta = await resultado.json();

        return respuesta;
    } catch (error) {
        throw new Error('Error al actualizar los registros: ' + error.message);
    }
}

async function metodoDELETE(peticionUrl) {
    try {
        const resultado = await fetch(peticionUrl, {
            method: "DELETE",
        });

        const respuesta = await resultado.json();

        return respuesta;
    } catch (error) {
        throw new Error('Error al eliminar los registros: ' + error.message);
    }
}
