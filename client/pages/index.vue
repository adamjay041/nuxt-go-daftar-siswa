<template>
    <view-dashboard />
</template>

<script>
import Header from '../components/Header.vue'
import viewDashboard from '../views/view-dashboard.vue'
export default {
  middleware: 'auth',
  components: { viewDashboard, Header },
  name: 'PageSiswa'
}
</script>
