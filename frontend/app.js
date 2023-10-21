function actualizarDatos() {
    // Realizar la solicitud al servidor Go para obtener datos actualizados
    fetch('/prueba')
        .then(response => response.json())
        .then(data => {
            const door = document.getElementById('door');
            const light = document.getElementById('light');
            const person = document.getElementById('person');

            // Recoger los datos
            const puerta = data.puerta;
            const luz = data.luz;
            const personas = data.personas;

            // Mostrar condición de la puerta
            if (puerta === 'abierto') {
                door.innerHTML = `
                    <div class="door-frame">
                        <div class="door open">Puerta Abierta</div>
                    </div>
                `;
            } else {
                door.innerHTML = `
                    <div class="door-frame">
                        <div class="door">Puerta Cerrada</div>
                    </div>
                `;
            }

            // Mostrar condición de la luz
            if (luz === "prendido") {
                light.innerHTML = `
                    <div class="light-frame">
                        <div class="light open">Luz prendida</div>
                    </div>
                `;
            } else {
                light.innerHTML = `
                    <div class="light-frame">
                        <div class="light">Luz apagada</div>
                    </div>
                `;
            }

            // Mostrar condición de las personas
            if (personas === "si") {
                person.innerHTML = `
                    <div class="person-frame">
                        <div class="person open">Hay personas</div>
                    </div>
                `;
            } else {
                person.innerHTML = `
                    <div class="person-frame">
                        <div class="person">No hay personas</div>
                    </div>
                `;
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

// Actualizar los datos cada 5 segundos (5000 milisegundos)
setInterval(actualizarDatos, 2000);
