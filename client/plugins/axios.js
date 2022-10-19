import axios from 'axios'
import configs from '@/app.config.json'


window.toastSuccess = function(description) {
  window.$nuxt.$bvToast.toast(description, {
    title: 'Success',
    autoHideDelay: 5000,
    // toaster: 'b-toaster-bottom-right',
    appendToast: false,
    variant: 'success'
  })
}
window.toastDanger = function(description) {
  window.$nuxt.$bvToast.toast(description, {
    title: 'Error',
    autoHideDelay: 5000,
    // toaster: 'b-toaster-bottom-right',
    appendToast: false,
    variant: 'danger'
  })
}

window.post = function (...args) {
  return new Promise((resolve, reject) => {
    axios.post(...args)
    .then(resolve)
      .catch(err => {
        if(err.response) {

          if(err.response.status == 500) {
            toastDanger("Internal error. please contact it support! [Code: 500]")
          }

          if(err.response.status == 404) {
            toastDanger("Internal error. please contact it support! [Code: 404]")
          }

          if(err.response.status == 400) {
            toastDanger(`${err.response.data.message}. please contact it support! [Code: 400]`)
          }

          if(err.response.status == 203) {
            toastDanger(`${err.response.data.message}. please contact it support! [Code: 400]`)
          }

          if(err.response.status == 401) {
            clearCookies();
            toastDanger(`${err.response.data.message}. please contact it support! [Code: 400]`)
            // location.href = '/login'
          }
        }

        if(err.message == 'Network Error')
          toastDanger("Network Error. can't connect to the server, please contact it support!")

        reject(err)
      })
  })
}

window.get = function (...args) {
  return new Promise((resolve, reject) => {
    axios.get(...args)
    .then(resolve)
      .catch(err => {
        if(err.response) {

          if(err.response.status == 500) {
            toastDanger("Internal error. please contact it support!")
          }

          if(err.response.status == 404) {
            toastDanger("Internal error. please contact it support!")
          }

          if(err.response.status == 401) {
            location.href = '/login'
          }
        }

        if(err.message == 'Network Error')
          toastDanger("Network Error. can't connect to the server, please contact it support!")

        reject(err)
      })
  })
}

window.put = function (...args) {
  return new Promise((resolve, reject) => {
    axios.put(...args)
    .then(resolve)
      .catch(err => {
        if(err.response) {

          if(err.response.status == 500) {
            toastDanger("Internal error. please contact it support!")
          }

          if(err.response.status == 404) {
            toastDanger("Internal error. please contact it support!")
          }

          if(err.response.status == 401) {
            location.href = '/login'
          }
          if(err.response.status == 400) {
            toastDanger(`${err.response.data.message}. please contact it support! [Code: 400]`)
          }
        }

        if(err.message == 'Network Error')
          toastDanger("Network Error. can't connect to the server, please contact it support!")

        reject(err)
      })
  })
}

window.deleted = function (...args) {
  return new Promise((resolve, reject) => {
    axios.delete(...args)
    .then(resolve)
      .catch(err => {
        if(err.response) {

          if(err.response.status == 500) {
            toastDanger("Internal error. please contact it support!")
          }

          if(err.response.status == 404) {
            toastDanger("Internal error. please contact it support!")
          }

          if(err.response.status == 401) {
            location.href = '/login'
          }
          if(err.response.status == 400) {
            toastDanger(`${err.response.data.message}. please contact it support! [Code: 400]`)
          }
        }

        if(err.message == 'Network Error')
          toastDanger("Network Error. can't connect to the server, please contact it support!")

        reject(err)
      })
  })
}


window.setCookie = function (cname, cvalue, ticks = 1) {
  var d = new Date();
  d.setTime(d.getTime() + ticks);
  var expires = 'expires=' + d;
  document.cookie = cname + '=' + cvalue + ';' + expires + ';path=/';
};

window.getCookie = function (cname) {
  var name = cname + "=";
  var decodedCookie = decodeURIComponent(document.cookie);
  var ca = decodedCookie.split(';');
  for(var i = 0; i <ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return null;
}

window.clearCookies = function() {
  document.cookie.split(";").forEach(function(c) { document.cookie = c.replace(/^ +/, "").replace(/=.*/, "=;expires=" + new Date().toUTCString() + ";path=/"); });
}

