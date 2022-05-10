<template>
<div class="body">
    <meu-header>
        <img class="img-header" src="../../assets/logo-black.svg" alt="">
    </meu-header>
    <ul>
        <li class="lista-logs">
            <div class="div-log">
                <h3 class="titulo">reservas:</h3>
            </div>
        </li>
    </ul>
    
</div> 
</template>

<script>
import Header from '../header/Header.vue'

export default {

    components: {
        'meu-header': Header
    }

}

</script>

<style>
.lista-logs {
    list-style: none;
}

.div-log {
    background-color: #8e9799 ;
    width: 50vw;
    height: 50vh;
    margin: 0px 100px;
    
}
.titulo {
    text-align: center;
}





</style>
