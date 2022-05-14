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
           <div class="painel" v-on:click="abriModalEPegaId(mesa.id)" v-bind:style=" mesa.status ? 'background-color: red; cursor: not-allowed; pointer-events: none;' : 'background-color: green;'">
             <div>
            <h3 class="centralizado">{{mesa.titulo}}</h3>
          </div>
        </div> 
      </li>
    </ul>
      <div v-if="modal" class="modal" >
        <div>
          <h2>Nos informe:</h2>
          <label for="horaI">O horario de inicio:</label>
          <input type="time" id="horaI" v-model="horarioI">
          <br>
          <label for="horaF">O horario final:</label>
          <input type="time" id="horaF" v-model="horarioF">
          <br>
          <label for="data">A data da reserva:</label>
          <input type="date" name="" id="data" v-model="dataInput">
          <br>
          <button v-on:click="modal = false">close</button>
          <button v-on:click="reservaMesa(selected, horarioI, horarioF, dataInput )">reservar</button>
        </div>
      </div>
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
      modal: false,
      horarioI : "",
      horarioF : "",
      dataInput : "",
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
          {
        id: '4',
        titulo: 'mesa 4',
        status: true,
        },
        {
        id: '5',
        titulo: 'mesa 5',
        status: false,
        },
         {
        id: '6',
        titulo: 'mesa 6',
        status: false,
        },
          {
        id: '7',
        titulo: 'mesa 7',
        status: true,
        },
        {
        id: '8',
        titulo: 'mesa 8',
        status: false,
        },
         {
        id: '9',
        titulo: 'mesa 9',
        status: false,
        },
          {
        id: '10',
        titulo: 'mesa 10',
        status: true,
        },
        {
        id: '11',
        titulo: 'mesa 11',
        status: false,
        },
         {
        id: '12',
        titulo: 'mesa 12',
        status: false,
        },
      ]
    }
  },
  
  // mounted(){
  //   Mesa.listar().then(resposta => {
  //     this.mesas = resposta.data
  //   })
  // },
  methods : {
    abriModalEPegaId(id) {
      this.modal = true
      this.selected = id
    },
    reservaMesa(selected, horarioI, horarioF, dataInput) {
      this.selected1 = selected
      this.horario1 = parseFloat(this.horarioI)
      this.horario2 = parseFloat(this.horarioF)
      this.dataI = dataInput
      
      if(window.confirm("voce quer reservar ?")) {
        //reservaMesa(selected1,horario1,horario2,dataI)
        alert(this.selected1)
        alert(this.horario1)
        alert(typeof(this.horario1))
        alert(this.horario2)
        alert(typeof(this.horario2))
        alert(this.dataI)
      }
    },
    // reservaMesa(id,horario1,horario2,dataI){
    //   const post:{
      //    id: this.id
      //    horario1: thishorario1
        //  horario2: this.horario2
        // dataI = this.dataI
        //}
        // axios.post(localhost:8000/,post)
        //.then((result) =>{
        // console.warn(result)  
        //})

    //}
    // })
    //
  
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
    width: 20%;
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
.modal {
    width: 300px;
    height: 200px;
    background-color: #dedede;
    position: absolute;
    top: 50%;
    left: 50%;
    margin-top: -150px;
    margin-left: -100px;
    box-shadow: 5px 5px 8px black;
    border: #000000 solid 1px;


}

</style>
