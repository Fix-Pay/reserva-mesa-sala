<template>
  <div>
    <meu-header>
        <img class="img-header" src="../../assets/logo-black.svg" alt="">  
    </meu-header>
    <h2 class="sub-titulo">Locais Reservaveis:</h2>
    <main class="salas"> 
        <div class="titulo-salas">
            <h2>Locais:</h2>
        </div>
      <ul class="lista-fotos">
       <li class="lista-fotos-item" v-for="mesa of mesas" v-bind:key="mesa">
           <div class="painel" v-on:click="pegaId(mesa.id)" v-bind:style=" mesa.status ? 'background-color: red;' : 'background-color: green;'">
             <div>
            <h3 class="centralizado">{{mesa.titulo}}</h3>
          </div>
        </div> 
      </li>
    </ul>
    <!-- <div class="titulo-salas">
        <h2>locais:</h2>
    </div>
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
    </ul> -->

    </main>
    
  </div>
</template>

<script>
import Header from '../header/Header.vue'
import Mesa from '../../services/mesas'


export default {

  components: {
    'meu-header': Header
  },
  data () {
    return {
      //ainda tem que fazer uma requisiçao pra api pra buscar esses dados abaixo de forma dinamica
      mesas: [
        {
        id: '1',
        titulo: 'mesa 1',
        status: true,
        },
        {
        id: '2',
        titulo: 'mesa 2',
        status: false,
        },
         {
        id: '3',
        titulo: 'mesa 3',
        status: false,
        },
      ]
    }
  },
  //get para pegar os locais

  // mounted(){
  //   Mesa.listar().then(resposta => {
  //     this.mesas = resposta.data
  //   })
  // },
  methods : {
    pegaId(id) {
      this.selected = id
      this.time = this.printTime();
      if(window.confirm("voce quer reservar ?")) {
        //reservaMesa(this.selected)
        alert(this.selected)
        alert(this.time)
      }
    },
     printTime: function () {
            return new Date().toLocaleTimeString();
          },
    // reservaMesa(id){
    //   Mesa.salvar(id).then(resposta => {
    //     alert('reservado')
    // })
    //}
  
  // if( resultado == 200){
    //alguma coisa pra reservar ou methodo de atualizar pagina
  //}
  //}
 
  } 
}
</script>

<style>

.sub-titulo {
  font-size: 40px;
  margin: 40px 45px 70px;
  color: #606566;
  text-align: center;
}
.titulo-salas {
    background: #ffff;
    border-bottom: #606566 solid 2px;


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
.painel {
  cursor: pointer;
}

</style>
