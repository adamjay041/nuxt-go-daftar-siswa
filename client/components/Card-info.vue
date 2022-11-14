<template>
  <div class="card-outer" style="margin:10px">
    <div class="status">
      <span>{{data.status}}</span>
    </div>

    <div class="info-profile">
      <div class="img-profile" style="margin-right:10px">
        <img width="80px" height="80px" src="https://pbs.twimg.com/media/EePk2PzUcAEKIJh.jpg" alt="">
      </div>
      <div class="info-data">
        <span class="bold-font" >{{data.name}}</span>
        <span>{{data.npm}}</span>
      </div>
    </div>


    <div class="info-profile-wrap">
      <div class="info-data-mg">
        <span class="bold-font" >Major</span>
        <span>{{data.major}}</span>
      </div>
        <div class="info-data-mg">
        <span class="bold-font" >Class</span>
        <span>{{data.kelas}}</span>
      </div>
      <div class="info-data-mg">
        <span class="bold-font" >Generation</span>
        <span>{{data.generasi}}</span>
      </div>
      <div class="info-data-mg">
        <span class="bold-font">Period</span>
        <span>{{data.period}}</span>
      </div>
    </div>

    <hr class="solid">

    <div class="info-profile-footer">
      <button class="btn-edit">
        <span>Edit Profile</span>
        <span>
      <font-awesome-icon :icon="['fas', 'pen-to-square']" />
        </span>
      </button>
      <button class="btn-delete">
        <span>Delete Profile</span>
        <span>
          <font-awesome-icon :icon="['fas', 'trash-can']" />
        </span>
      </button>
    </div>
</div>

</template>

<script>


export default {
  props: ['data']
}
</script>

<style lang="scss">
  #__nuxt {
    height: 100%;
  }

  #__layout {
    height: 100%;
  }

  .info-profile {
    display: flex;
  }

  .btn-edit {
    border-radius: 5px;
    margin-right: 10px;
    border: 1px solid transparent;
    background-color: white;
    color: #4F84EA;
  }


  .btn-edit:hover {
    border-radius: 5px;
    margin-right: 10px;
    border: 1px solid transparent;
    background-color: rgb(244, 244, 244);
    color: #4F84EA;
  }


  .btn-delete {
    // margin-right: 10px;
    border: 1px solid transparent;
    background-color: white;
    color: #e94949;
    border-radius: 5px;
  }

  .btn-delete:hover {
    // margin-right: 10px;
    border-radius: 5px;
    border: 1px solid transparent;
    background-color: rgb(244, 244, 244);
    color: #e94949;
  }



  hr {
    margin-top: 0.5rem;
    margin-bottom: 0.5rem;
    border: 0;
    border-top: 3px solid #4F84EA;
  }


  .info-profile-wrap {
    margin-top: 5px;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
  }



  .info-profile-footer {
    display: flex;
    flex-wrap: wrap;
    justify-content:flex-end;
  }


  .info-data {
    display: grid;
    align-content: center;
    font-size: 12px;
  }

  .info-data-mg {
    display: grid;
    align-content: center;
    font-size: 12px;
    // margin: 9px;
    margin: 7px 20px 9px 9px
  }

  .bold-font {
    font-weight:bold;
  }


  .img-profile img {
    width: 70px;
    height: 70px;
    background-color: aqua;
    border-radius: 100%;
  }


  .status {
    // position: absolute;
    padding: 10;
    width: 65px;
    height: 25px;
    // float: right;
    background: rgba(0, 191, 142, 1);
    border-radius: 3px;
    display: flex;
    justify-content: center;
  }


  .status span {
    // font-family: 'Montserrat';
    font-style: normal;
    font-weight: 400;
    font-size: 12px;
    line-height: 10px;
    display: flex;
    align-items: center;
    text-align: center;
    color: #FFFFFF;
  }



  // .content-center {
  //   overflow: auto;
  //   height: inherit;
  //   width: 100%;
  //   display: flex;
  //   flex-wrap: wrap;
  //   align-content: center;
  //   align-items: center;
  //   justify-content: center;
  //   background-image: url("https://lms.unindra.ac.id/assets/template/login_colek/images/img-wayang.png");
  //   background-repeat: no-repeat;
  //   background-size: cover;
  //   background-position: center;
  // }
  .card-outer {
    // position: absolute;
    width: 400px;
    padding: 10px;
    // display: flex;
    background: #ffffff;
    border-radius: 10px;
    border-width: 10px;
    border-style: solid;
    border: 1px solid #4F84EA;
  }
</style>
