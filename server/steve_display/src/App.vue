<script setup>
import { errorMessages } from '@vue/compiler-core'
import { ref, onMounted, onUnmounted } from 'vue'
import { RouterView } from 'vue-router'
// import { getClientId } from './utils/clientId'
const user = ref('')
const user_id = ref('U-98765')
const cart_id = ref('')
const cart_version = ref('v1.0.3')
const message = ref('Your smart cart assistant!')
const time = ref('')
const date = ref('')
const weatherData = ref(null)
const errMessage = ref('')
const weight = ref('54.3 kg')
const price = ref('23.75 €')
const isMenuOpen = ref(false)
const isLogin = ref(false)
const getWeather = async () => {
  errMessage.value = ''
  weatherData.value = null
  const lat = ref(37.9838)  // Example latitude for Athens
  const lon = ref(23.7275)  // Example longitude for Athens

  try {
    if (!lat.value || !lon.value) {
      throw new Error('Latitude and Longitude must be provided!')
    }

    const url = `https://api.open-meteo.com/v1/forecast?latitude=${lat.value}&longitude=${lon.value}&current_weather=true`

    const response = await fetch(url)
    if (!response.ok) {
      throw new Error(`Error fetching weather data: ${response.statusText}`)
    }

    const data = await response.json()
    weatherData.value = data.current_weather
  } catch (error) {
    errMessage.value = error.message
  }
}

if (!user && !user_id.value) {
  throw new Error('User information is missing! Login or sign up to get more features.')
} else {
  isLogin.value = true
  user.value = localStorage.getItem('smartcart_user')
  user_id.value
}


const sen_bool_1 = ref('Active')
const sen_bool_2 = ref('Active')
const sen_bool_3 = ref('Inactive')
const sen_bool_4 = ref('Active')
const sen_bool_5 = ref('Inactive')


let timer = null

function updateTime() {
  const now = new Date()
  time.value = now.toLocaleTimeString('el-GR', {
    timeZone: 'Europe/Athens',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
  })
  date.value = now.toLocaleDateString('el-GR', {
    timeZone: 'Europe/Athens',
    weekday: 'long',
    day: '2-digit',
    month: 'long',
    year: 'numeric',
  })
}


onMounted(() => {
  updateTime()
  timer = setInterval(updateTime, 1000)
  getWeather()
  document.documentElement.classList.add('dark')
})

onUnmounted(() => {
  clearInterval(timer)
})

onMounted(() => {
  cart_id.value = getClientId()
})

</script>

<template>
  <router-view></router-view>
</template>
<style>
@font-face {
  font-family: 'Minimal';
  src: url('/steve_display/public/minimal.otf') format('opentype');
  font-weight: normal;
  font-style: normal;
}

.app-background {
  background-image: url('/steve_display/public/background.png');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
}

.support,
.cart-details {
  /* position: fixed;
  bottom: 0;
  left: 300; */
  border: 1px solid #ccc;
}
</style>