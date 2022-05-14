<template>
<div class="body">
<<<<<<< HEAD
  <meu-header>
    <navbar placement="top" type="default">
      <a slot="brand" href="/" title="Home" class="navbar-brand"> <img class="img-header" src="../../assets/logo-black.svg" alt=""> </a>
    </navbar>
  </meu-header>
=======
    <meu-header>
        <img class="img-header" src="../../assets/logo-black.svg" alt="">
    </meu-header>
>>>>>>> feature/front1
        <div class="container">
            <h2 class="titulo">Reservados:</h2>
            <ul class="ul-lista">
                <li class="lista-logs"></li>
<<<<<<< HEAD
            </ul>
        </div>
</div>
=======
            </ul>  
        </div> 
</div> 
>>>>>>> feature/front1
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
.container {
    width: 550px;
    height: 450px;
    background-color: #858687;
    position: absolute;
    top: 50%;
    left: 50%;
    margin-top: -225px;
    margin-left: -275px;
    box-shadow: 5px 5px 8px black; ;
}
.titulo {
    text-align: center;
}



</style>
