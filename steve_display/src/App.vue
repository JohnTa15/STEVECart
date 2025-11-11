<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const message = ref('Hello {{user}}, this is S.T.E.V.E Display. The smart cart assistant!')
const time = ref('')
const date = ref('')

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
})

onUnmounted(() => {
  clearInterval(timer)
})
</script>

<template>
  <div class="app-background flex flex-col items-center justify-center min-h-screen text-white">
    <header class="backdrop-blur-md bg-black/30 p-10 rounded-2xl shadow-2xl text-center space-y-6">
      <h1 class="text-3xl font-bold tracking-wide">
        {{ message }}
      </h1>

      <div class="text-6xl font-light tracking-wider">
        {{ time }}
      </div>
      <div class="text-lg opacity-80 capitalize">
        {{ date }}
      </div>

      <div class="grid grid-cols-4 gap-6 mt-10 text-center">
        <div class="bg-white/10 rounded-xl p-4 hover:bg-white/20 transition">
          <div class="relative">Search Bar
            <input class="w-full bg-transparent placeholder:text-slate-500 rounded-4xl transition duration-300"
            placeholder="Search here for products here"/>
          </div>
          <p class="text-xl font-semibold">Recent bought products</p>
        </div>
        <div class="bg-white/10 rounded-xl p-4 hover:bg-white/20 transition">
          <p class="text-xl font-semibold">Minimap
            <!-- <img src="/src/assets/supermarket_map.png" alt="Minimap" class="mt-2 rounded-lg"/> -->
          </p>
        </div>
        <div class="bg-white/10 rounded-xl p-4 hover:bg-white/20 transition">
          <p class="text-xl font-semibold">Weather Widget</p>
        </div>
        <div class="bg-white/10 rounded-xl p-4 hover:bg-white/20 transition space-y-1">
          <p class="text-xl font-semibold">Total Weight</p>
          <p class="text-xl font-semibold">Total Price</p>
        </div>
      </div>
    </header>
  </div>
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
</style>
