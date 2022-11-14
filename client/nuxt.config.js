export default {
  ssr: false,

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: "Daftar Siswa",
    htmlAttrs: {
      lang: "en",
    },
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "" },
      { name: "format-detection", content: "telephone=no" },
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "https://lms.unindra.ac.id/media_public/get_gambar/d2llSWJaazRBRUtSZ2lFQm15V3RiZi9GV3oyd09zSGdTdXlOVjVyR0Zsc20yWkxNc3J4dUNFdi9MbENKd2ZYSg==" }],
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ["@/assets/app.scss",'@fortawesome/fontawesome-svg-core/styles.css'],
  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: ["@/plugins/axios", "@/plugins/vendor"],

  target: "static",

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: [
    // Equivalent to { path: '~/components' }

    "~/components",
    { path: "~/components/", extensions: ["vue"] },
  ],

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/bootstrap
    "bootstrap-vue/nuxt",
    // '@nuxtjs/fontawesome',
  ],

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},
};
