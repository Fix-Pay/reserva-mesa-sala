<template>
  <div class="corpo">
    <meu-header>
      <img class="img-header" src="./assets/logo-black.svg" alt="">
    </meu-header>
    <h2 class="sub-titulo">Reserve uma sala de reunião</h2>

    <main class="salas"> 
      <ul class="lista-fotos">
       <li class="lista-fotos-item" v-for="foto of fotos" v-bind:key="foto">
           <div class="painel">
              <img class="image-responsive" :src="foto.url" :alt="foto.titulo">
            <div>
            <h3 class="centralizado">{{foto.titulo}}</h3>
            <p class="centralizado">{{foto.desc}}</p>
          </div>
        </div> 
      </li>
    </ul>

    </main>
    
  </div>
</template>

<script>
import Header from './components/header/Header.vue'


export default {

  components: {
    'meu-header': Header
  },
  data () {
    return {
      //ainda tem que fazer uma requisiçao pra api pra buscar esses dados abaixo de forma dinamica
      fotos: [
        {
        url: 'https://cdn.eurekacoworking.com/wp-content/uploads/2020/08/09214622/Sala-reuniao-2-4-pessoas_2.jpg',
        titulo: 'sala 1707',
        desc: 'sala de reunião'
        },
        {
          url: 'https://www.deltabc.com.br/uploads/salas-de-reuniao-para-alugar.jpg',
          titulo: 'sala 1709',
          desc: 'sala de reunião'
        },
        {
          url: 'https://bclass.com.br/wp-content/uploads/2020/07/20170528_131819-1024x768.jpg',
          titulo: 'sala 1906',
          desc: 'sala de reunião'
        }
      ]
    }

  }
}
</script>

<style>
.corpo {
  margin: 0;
  padding: 0;
}
.sub-titulo {
  font-size: 40px;
  margin: 40px 45px;
  color: #606566;
}
.lista-fotos {
  list-style: none;
  display: inline-block;
}

.centralizado {
  text-align: center;
}
.lista-fotos-item {
  display: inline-block;
    align-items: center;
    width: 30%;
    vertical-align: top;
    border: 2px solid #000000;
    border-radius: 10px;
    margin: 0 1.5%;
    box-sizing: border-box;
}
.image-responsive {
  width: 100%;
}
</style>
