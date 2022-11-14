import Vue from 'vue'
import { BootstrapVue, BootstrapVueIcons } from 'bootstrap-vue'
import { library, config } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'
import feather from 'vue-icon';



config.autoAddCss = false

library.add(fas)


// Register the component globally
Vue.component('font-awesome-icon', FontAwesomeIcon)

Vue.use(feather, 'feather-icon');

Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)


import Multiselect from 'vue-multiselect'
Vue.component('multiselect', Multiselect)
import 'vue-multiselect/dist/vue-multiselect.min.css';


import DatePicker from "vue2-datepicker";
Vue.use(DatePicker)

import VueRangedatePicker from 'vue-rangedate-picker'
Vue.use(VueRangedatePicker)

import MonthPicker from 'vue-month-picker'
Vue.use(MonthPicker)
