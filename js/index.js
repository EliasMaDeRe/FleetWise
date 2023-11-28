(function () {

    const headingBienvenida = document.querySelector(".menu__contenedor .menu__heading");
    if(headingBienvenida){
        headingBienvenida.textContent = "Bienvenido";
    }

    const logoHeader = document.querySelector(".header__logo");

    if(logoHeader){
        crearImagenLogo();
    }

    function crearImagenLogo(){
        const imagen = document.createElement("img");
        imagen.classList.add("header__logo-image");
        imagen.src = "img/Logo.png"
        imagen.addEventListener("click", (e) =>{
            e.preventDefault();
            solicitarWeb("/")
        })
        logoHeader.appendChild(imagen);
    }

    const barraNavegacion = document.querySelector(".navbar");

    const textoBotones = {
        Vehiculos: 'Capturar Vehiculos',
        Servicios: 'Capturar Conceptos',
        Registros: 'Agregar Registro',
    }

    const urlBotones = {
        Vehiculos: "/AgregarVehiculo",
        Servicios: "/AgregarServicioVehicular",
        Registros: "/SeleccionarVehiculoParaNuevoRegistro",
    }

    if(barraNavegacion){
        crearBoton("Vehiculos");
        crearBoton("Servicios");
        crearBoton("Registros")
        crearBotonCerrarSesion();
    }

    function crearBoton(servicio){
        const boton = document.createElement("a");
        boton.classList.add("navbar__enlace");
        boton.textContent = textoBotones[servicio];
        boton.addEventListener("click", (e) =>{
            e.preventDefault();
            solicitarWeb(urlBotones[servicio])
        })
        barraNavegacion.appendChild(boton);
    }

    function crearBotonCerrarSesion(){
        const boton = document.createElement("a");
        boton.classList.add("navbar__enlace--logout");
        boton.textContent = "Cerrar Sesión";
        boton.addEventListener("click", (e) =>{
            e.preventDefault();
            solicitarWeb("/Logout")
        })
        barraNavegacion.appendChild(boton);
    }

    function solicitarWeb(url){
        window.location.href = url;
    }

  })();
  