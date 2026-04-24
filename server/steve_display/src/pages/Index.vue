<template>
  <div class="app-background flex flex-col items-center justify-center min-h-screen text-white">

    <header
      class="relative backdrop-blur-md bg-black/30 dark:bg-gray-900/60 p-10 rounded-2xl shadow-2xl text-center space-y-6 border border-white/10 dark:border-gray-700 max-w-4xl w-full">

      <div class="absolute top-6 right-6 z-50 text-left">
        <button @click="isMenuOpen = !isMenuOpen"
          class="relative group w-12 h-12 rounded-full bg-white/10 dark:bg-white hover:bg-white/20 dark:hover:bg-gray-700 shadow-xl ring-1 ring-white/10 dark:ring-gray-700 flex items-center justify-center transition-all duration-300 focus:outline-none">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16M4 16h16" />
          </svg>
          <div
            class="flex flex-col justify-between w-20px h-20px transform transition-all duration-300 origin-center overflow-hidden">
            <div
              class="bg-white dark:text-white h-2px w-7 transform transition-all duration-300 origin-left delay-150ms"
              :class="{ 'rotate-42deg w-2/3': isMenuOpen }"></div>
            <div class="bg-white dark:text-white h-2px w-7 rounded transform transition-all duration-300"
              :class="{ 'translate-x-10': isMenuOpen }"></div>
            <div
              class="bg-white dark:text-white h-2px w-7 transform transition-all duration-300 origin-left delay-150ms"
              :class="{ '-rotate-42deg w-2/3': isMenuOpen }"></div>
          </div>
        </button>

        <transition enter-active-class="transition duration-200 ease-out"
          enter-from-class="opacity-0 -translate-y-2 scale-95" enter-to-class="opacity-100 translate-y-0 scale-100"
          leave-active-class="transition duration-150 ease-in" leave-from-class="opacity-100 translate-y-0 scale-100"
          leave-to-class="opacity-0 -translate-y-2 scale-95">
          <div v-if="isMenuOpen"
            class="absolute right-0 mt-3 w-64 origin-top-right rounded-xl bg-gray-900/95 backdrop-blur-xl shadow-2xl ring-1 ring-white/10 dark:ring-gray-700 py-2 focus:outline-none overflow-hidden">
            <div class="block px-4 py-3 text-sm text-gray-400 hover:bg-white/5 transition-colors cursor-default">
              <span class="block text-xs uppercase tracking-wider opacity-50">Announcements</span>
            </div>

            <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ" target="_blank" rel="noopener noreferrer"
              class="block px-4 py-3 text-sm text-gray-300 hover:bg-white/10 hover:text-white transition-colors border-b border-white/5">
              Need Help? Contact Support!
            </a>

            <div class="block px-4 py-3 text-sm text-gray-400 hover:bg-white/5 transition-colors cursor-default">
              <span class="block text-xs uppercase tracking-wider opacity-50">Cart ID</span>
              <span class="text-white font-mono">{{ cart_id }}</span>
            </div>

            <div class="block px-4 py-3 text-sm text-gray-400 hover:bg-white/5 transition-colors cursor-default">
              <span class="block text-xs uppercase tracking-wider opacity-50">Firmware</span>
              <span class="text-white font-mono">{{ cart_version }}</span>
            </div>

            <a href="#"
              class="block px-4 py-3 text-sm text-gray-300 hover:bg-white/10 hover:text-white transition-colors">
              Sensors Test
            </a>

            <a href="#"
              class="block px-4 py-3 text-sm text-red-400 hover:bg-red-500/10 hover:text-red-300 transition-colors border-t border-white/5">
              Logout
            </a>

          </div>
        </transition>
      </div>
      <h1 class="text-3xl font-bold tracking-wide mt-8">
        <p class="title dark:text-white">Hello {{ user }}, this is S.T.E.V.E Display.</p>
        <p class="subtitle text-lg font-light opacity-80 dark:text-gray-300">{{ message }}</p>
      </h1>

      <!-- Flashlight Warning Banner -->
      <transition enter-active-class="transition duration-500 ease-out"
        enter-from-class="opacity-0 -translate-y-4" enter-to-class="opacity-100 translate-y-0"
        leave-active-class="transition duration-300 ease-in" leave-from-class="opacity-100 translate-y-0" leave-to-class="opacity-0 -translate-y-4">
        <div v-if="isBlackout" class="w-full bg-yellow-500/90 text-gray-900 rounded-xl p-6 mt-6 flex flex-col md:flex-row justify-between items-center shadow-red-500/50 shadow-2xl border border-yellow-300 animate-pulse">
          <div class="text-left mb-4 md:mb-0">
            <h2 class="text-2xl font-bold">⚠️ BLACKOUT DETECTED!</h2>
            <p class="font-medium">Light levels dropped below safe limits. Do you need the flashlight?</p>
          </div>
          <button @click="flashlightOn = !flashlightOn" 
            class="bg-gray-900 text-white font-bold py-3 px-8 rounded-full hover:bg-black transition shadow-lg border border-gray-700 whitespace-nowrap">
            {{ flashlightOn ? 'TURN OFF FLASHLIGHT' : '🔦 TURN ON FLASHLIGHT' }}
          </button>
        </div>
      </transition>

      <div class="text-lg opacity-80 capitalize dark:text-gray-400">
        {{ date }}
      </div>
      <div class="text-2xl font-light tracking-wider dark:text-white">
        {{ time }}
      </div>

      <div class="relative">
        <input
          class="w-full bg-white/5 dark:bg-gray-800/50 placeholder:text-slate-500 rounded-4xl transition duration-300 border border-white/10 dark:border-gray-700 p-2"
          placeholder="Search here for products here" />
      </div>
      <div class="flex flex-wrap gap-6 mt-10 text-center justify-end">
        <button @click="logout"
          class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-fit mx-auto inline-block">Logout</button>
      </div>

      <div
        class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-fit mx-auto inline-block">
        <h3 class="text-xl font-semibold dark:text-white">Current scanned products list</h3>
        <ul class="text-lg font-light dark:text-gray-300 w-fit">
          <li>Product A</li>
          <li>Product B</li>
          <li>Product C</li>
        </ul>
      </div>

      <div class="flex flex-wrap gap-6 mt-10 text-center justify-center">

        <div
          class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-fit">
          <p class="text-xl font-semibold dark:text-white">Recent bought products</p>
          <ul class="bg-white/5 dark:bg-gray-900/30 rounded-xl p-4 mt-2 space-y-1 dark:text-gray-300 w-fit">
            <li>Product 1</li>
            <li>Product 2</li>
            <li>Product 3</li>
          </ul>
        </div>

        <div
          class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-fit">
          <p class="text-xl font-semibold dark:text-white">Minimap</p>
          <img src="../assets/minimap.png" alt="Supermarket Minimap" class="w-full h-full object-cover" />
        </div>

        <div
          class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-fit">
          <p class="text-xl font-semibold dark:text-white">Weather info</p>
          <div v-if="weatherData" class="mt-2 dark:text-gray-300">
            <p>Temp: {{ weatherData.temperature }}°C</p>
            <p>Wind: {{ weatherData.windspeed }} km/h</p>
          </div>
          <div v-else-if="!errMessage" class="animate-pulse">Loading..</div>
          <p v-if="errMessage" class="text-red-400">{{ errMessage }}</p>
        </div>

        <div
          class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition space-y-1 border border-transparent dark:border-gray-700 w-fit">
          <p class="text-xl font-semibold dark:text-white">Details</p>
          <p class="dark:text-gray-300">Weight: {{ weight }}</p>
          <p class="dark:text-gray-300">Price: {{ price }}</p>
        </div>
      </div>
    </header>
  </div>
</template>
<script>
export default {
  data() {
    return {
      weight: "0.0 kg",
      price: "$0.00",
      weightStatus: "",
      scannedProduct: "",
      // Added these for dynamic API calls!
      cart_id: "DISPLAY_CART_01", 
      current_nfc_tag: "NFC_456",
      isBlackout: false,
      flashlightOn: false
    };
  },
  mounted() {
    // Poll the light sensor every 2 seconds for real-time blackout detection
    setInterval(this.checkEnvironmentLight, 2000);
  },
  methods: {
    logout() {
      this.$router.push('/login')
    },
    async checkEnvironmentLight() {
      try {
        const url = `http://localhost:8089/measureLight?cartID=${this.cart_id}`;
        const response = await fetch(url);
        const data = await response.json();
        
        if (data.flashlight_needed) {
          this.isBlackout = true;
          // You could automatically turn it on here if you wanted by setting this.flashlightOn = true;
        } else {
          this.isBlackout = false;
          // this.flashlightOn = false; // Optionally automatically turn it off if light is back
        }
      } catch (error) {
        console.error("Failed to read light sensor", error);
      }
    },
    //weight checking from the scale vs weight of the database
    async triggerWeightCheck() {
      try {
        const url = `http://localhost:8089/measureWeight?cartID=${this.cart_id}&tag=${this.current_nfc_tag}`;
        const response = await fetch(url);
        const data = await response.json();

        if (data.status === "correct") {
          this.weightStatus = "Weight Match!";
          this.weight = data.actual_weight + " kg";
          this.scannedProduct = data.product_name;
        } else {
          this.weightStatus = "Weight Mismatch!";
          this.weight = data.actual_weight + " kg (Expected: " + data.expected_weight + ")";
        }
      } catch (error) {
        console.error("Failed to fetch weight data:", error);
        this.weightStatus = "Error reaching Scale API";
      }
    }
  }
};
</script>