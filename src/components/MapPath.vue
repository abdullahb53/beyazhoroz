<template>
<div class="container">

      <div class="row">
        <div class="col"></div>
        <div class="col-10">

             <h1>{{ secilen_sehir }}</h1>

            <div v-for="item in 10" :key="item" class="card" style="width: 13rem;" >
              <img class="card-img-top" src="img/noPhoto.svg" alt="Card image cap">
              <div class="card-body">
                <h5 class="card-title">Uretici İsmi</h5>
                <p class="card-text"> elma,armut,patates,marul,portakal </p>
                <a href="#" class="btn btn-primary">İletişim</a>
              </div>
            </div>

        </div>
        <div class="col"></div>
      </div>
</div>
</template>

<script>
export default {
  name: 'MapPath',
  data (){
    return {
      secilen_sehir
    }

  }
}
var secilen_sehir="";

secilen_sehir = window.location.pathname.split("/")[1].toUpperCase();
console.log(secilen_sehir);




</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.card{
float:left;
margin-top:10px;


}

</style>
