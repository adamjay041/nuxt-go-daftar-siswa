export default function ({ store, redirect }) {
  console.log(getCookie('token'))
  if (!getCookie('token')) {
    location.href = '/login';
  }
}
