<template>

  <Tutorial />
</template>

<script>


export default {
  name: 'PageSiswa',
  mounted: {

  },
  methods: {
    getData() {
      this.$store.dispatch("siswa/getListSiswa")
    }
  },
  mounted(){
    console.log('this mounted')
    this.getData()
  }
}
</script>
