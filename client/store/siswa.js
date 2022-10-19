

export const state = () => ({
  listSiswa : [],
  formSubmit: {
    First_Name: "a",
    Last_Name: "b",
    Full_Name: "c",
    Class: "sn2",
    NPM : "2020202",
    Email: "a@m.com",
    Password: "12345678"
  },
  isBusy: false,
  titleModal: ""
})

export const mutations = {
  setListSiswa(state, value) {
    if(value['id'] == 1) {
      console.log('ini masuk arr', value)
      state.listSiswa = value['data']
    } else {
      value.id = state.listSiswa.length +1
      state.listSiswa.push(value['data']);
    }
  },
  setFormSubmit (state,value) {
    state.formSubmit = value
  },
  setFormSubmitDefault(state, value) {
    state.formSubmit = {
      first_name: "",
      last_name: "",
      full_name: "",
      kelas: "",
      npm :"",
      id: ""
    }
    console.log("masuk sini harusnysa", state.formSubmit)
  },
  setTitleModal(state, val) {
    state.titleModal = val
  },
  setDeleteSiswa(state,value) {
    state.listSiswa.splice(value,1)
  },
  setIsBusy(state, value) {
    state.isBusy = value;
  }
}


export const actions ={
  getListSiswa: function (context) {

    // token = getCookie("token")
    // console.log(token, "ININ TOKEKN")
    get("http://localhost:8080/",{
      headers: {
        Authorization: 'Bearer ' + getCookie("token") //the token is a variable which holds the token
      }
    })
    .then(r => {
      var arr = []
      r.data.data.forEach((e,index) => {
        console.log('ini masuk arr', e)
        var obj = {
          First_Name: e.first_name,
          Last_Name: e.last_name,
          Full_Name: e.full_name,
          Class: e.class,
          NPM : e.npm.replace(/ /g,''),
          Email: e.email,
          Id: e.id,
          Role: e.role
        }
        arr.push(obj)
      })
      context.commit('setListSiswa', {data:arr, id: 1})
    })
    .catch(err => {
      console.log(err)
    })
  },
  postSiswa: function (context, value) {
    context.commit("setIsBusy", true)
    console.log(context )
    post("http://localhost:8080/student/post", value,{
      headers: {
        Authorization: 'Bearer ' + getCookie("token") //the token is a variable which holds the token
      }
    })
    .then(r => {
      toastSuccess("Siswa Berhasil di tambahkan")
      context.dispatch('getListSiswa')
      context.commit("setIsBusy", false)
    })
    .catch(err => {
      console.log(err)
      context.commit("setIsBusy", false)
    })
    // .finally(e => )
  },
  deleteSiswa: function(context,id) {
    context.commit("setIsBusy", false)
    deleted(`http://localhost:8080/student/${id}`,{
      headers: {
        Authorization: 'Bearer ' + getCookie("token") //the token is a variable which holds the token
      }
    })
    .then(r => {
      toastSuccess("Siswa Berhasil di hapus")
      context.dispatch('getListSiswa')
      context.commit("setIsBusy", false)
    })
    .catch(e => {
      context.commit("setIsBusy", false)
    })
  },
  putSiswa: function(context, value) {
    context.commit("setIsBusy", true)
    put(`http://localhost:8080/student/${value['id']}`, value['data'],{
      headers: {
        Authorization: 'Bearer ' + getCookie("token") //the token is a variable which holds the token
      }
    })
    .then(r => {
      toastSuccess("Siswa Berhasil di update")
      context.dispatch('getListSiswa')
      context.commit("setIsBusy", false)
    })
    .catch(e => {
      context.commit("setIsBusy", false)
    })
  },
  loginSiswa: function (context, value) {
    context.commit("setIsBusy", true)
    // console.log(context )
    post("http://localhost:8080/login", value)
    .then(r => {
      toastSuccess("Login Berhasil")
      // localStorage.setItem("authToken", token)
      console.log(getCookie('token'))
      setCookie("token", r.data.data, 60 * 4 * 1000)
      // context.dispatch('getListSiswa')
      this.$router.push("/");
      toastSuccess()
      context.commit("setIsBusy", false)
    })
    .catch(err => {
      context.commit("setIsBusy", false)
    })
  },
  createSiswa: function (context, value) {
    context.commit("setIsBusy", true)
    console.log(context )
    post("http://localhost:8080/create", value)
    .then(r => {
      // context.dispatch('getListSiswa')
      context.commit("setIsBusy", false)
    })
    .catch(err => {
      console.log(err)
      context.commit("setIsBusy", false)
    })
    // .finally(e => )
  },
}
