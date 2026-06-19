<template>
  <div class="app-background flex flex-col items-center justify-center min-h-screen text-white">

    <header
      class="relative backdrop-blur-md bg-black/30 dark:bg-gray-900/60 p-10 rounded-2xl shadow-2xl text-center space-y-6 border border-white/10 dark:border-gray-700 max-w-4xl w-full">

      <div class="absolute top-6 right-6 z-50 text-left">
        <button @click="logout"
          class="bg-white/10 dark:bg-white/10 hover:bg-white/20 dark:hover:bg-gray-700 shadow-xl ring-1 ring-white/10 dark:ring-gray-700 px-6 py-2.5 rounded-xl transition-all duration-300 focus:outline-none text-sm font-semibold text-white">
          Logout
        </button>
      </div>
      <h1 class="text-3xl font-bold tracking-wide mt-8">
        <p class="title dark:text-white">Hello {{ user }}, this is S.T.E.V.E Display.</p>
        <p class="subtitle text-lg font-light opacity-80 dark:text-gray-300">{{ message }}</p>
      </h1>

      <!-- Assistance Confirmation Banner -->
      <transition enter-active-class="transition duration-500 ease-out" enter-from-class="opacity-0 -translate-y-4"
        enter-to-class="opacity-100 translate-y-0" leave-active-class="transition duration-300 ease-in"
        leave-from-class="opacity-100 translate-y-0" leave-to-class="opacity-0 -translate-y-4">
        <div v-if="assistanceRequested"
          class="w-full bg-green-500/90 text-gray-900 rounded-xl p-6 mt-6 flex flex-col md:flex-row justify-between items-center shadow-green-500/50 shadow-2xl border border-green-300">
          <div class="text-left mb-4 md:mb-0">
            <h2 class="text-2xl font-bold">Help is on the way!</h2>
            <p class="font-medium">A staff member has been notified and will assist you shortly.</p>
          </div>
          <button @click="assistanceRequested = false"
            class="bg-gray-900 text-white font-bold py-3 px-8 rounded-full hover:bg-black transition shadow-lg border border-gray-700 whitespace-nowrap">
            Dismiss
          </button>
        </div>
      </transition>

      <!-- Flashlight Warning Banner -->
      <transition enter-active-class="transition duration-500 ease-out" enter-from-class="opacity-0 -translate-y-4"
        enter-to-class="opacity-100 translate-y-0" leave-active-class="transition duration-300 ease-in"
        leave-from-class="opacity-100 translate-y-0" leave-to-class="opacity-0 -translate-y-4">
        <div v-if="isBlackout"
          class="w-full bg-yellow-500/90 text-gray-900 rounded-xl p-6 mt-6 flex flex-col md:flex-row justify-between items-center shadow-red-500/50 shadow-2xl border border-yellow-300 animate-pulse">
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
        <button @click="assistance"
          class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-fit mx-auto inline-block">
          Need Help?</button>
      </div>

      <div
        class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-fit mx-auto inline-block text-left">
        <p class="text-xl font-semibold dark:text-white mb-2">Cart Summary</p>
        <ul class="bg-white/5 dark:bg-gray-900/30 rounded-xl p-4 space-y-2 dark:text-gray-300 w-full min-w-[300px]">
          <li class="flex justify-between gap-4">
            <span class="opacity-70">Total Weight:</span>
            <span class="font-medium text-white">{{ totalWeight }} kg</span>
          </li>
          <li class="flex justify-between gap-4">
            <span class="opacity-70">Total Price:</span>
            <span class="font-medium text-white">{{ totalPrice }} €</span>
          </li>
          <li class="flex justify-between gap-4">
            <span class="opacity-70">Status:</span>
            <span class="font-medium"
              :class="{ 'text-green-400': weightStatus === 'Weight Match!', 'text-red-400': weightStatus && weightStatus !== 'Weight Match!' }">
              {{ weightStatus || '--' }}
            </span>
          </li>
        </ul>
      </div>

      <div class="flex flex-wrap gap-6 mt-10 text-center justify-center">

        <div
          class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-full max-w-md text-left flex-grow">
          <p class="text-xl font-semibold dark:text-white mb-4">Scan Products</p>
          <div class="space-y-2 max-h-48 overflow-y-auto pr-2 custom-Productsar">
            <div v-for="(item, index) in scannedProducts" :key="index"
              class="bg-white/5 dark:bg-gray-900/30 rounded-lg p-3 flex justify-between items-center">
              <div>
                <p class="font-bold text-white">{{ item.name }}</p>
                <p class="text-sm opacity-70">{{ item.weight }} • {{ item.price }}</p>
              </div>
              <span class="text-sm font-bold px-2 py-1 rounded"
                :class="{ 'bg-green-500/20 text-green-400': item.status === 'Match', 'bg-red-500/20 text-red-400': item.status === 'Mismatch' }">
                {{ item.status }}
              </span>
            </div>
            <div v-if="scannedProducts.length === 0" class="text-center opacity-50 py-4">
              No products scanned yet.
            </div>
          </div>
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

      </div>
      <div
        class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-fit mx-auto mt-6">
        <p class="text-xl font-semibold dark:text-white mb-3">Minimap</p>
        <img id="source" src="../assets/minimap.png" alt="Supermarket Minimap" style="display: none;" />
        <canvas id="canvas" width="450" height="300" class="border border-white/10 rounded-xl shadow-inner"></canvas>
      </div>

    </header>
  </div>
</template>
<script>
import { API_URL, WS_URL } from '../config.js';

export default {
  data() {
    return {
      user: localStorage.getItem('username') || "Guest",
      message: localStorage.getItem('loginMessage') || "Welcome to S.T.E.V.E Display",
      date: "",
      time: "",
      weight: "", //kg 
      price: "", //€
      weightStatus: "",
      scannedProduct: "",
      scannedProducts: [],
      // Added these for dynamic API calls!
      cart_id: "DISPLAY_CART_01",
      current_nfc_tag: "NFC_456",
      isBlackout: false,
      flashlightOn: false,
      assistanceRequested: false,
      weatherData: null,
      errMessage: "",
    };
  },
  computed: {
    totalWeight() {
      return this.scannedProducts.reduce((acc, item) => {
        const val = parseFloat(item.weight) || 0;
        return acc + val;
      }, 0).toFixed(2);
    },
    totalPrice() {
      return this.scannedProducts.reduce((acc, item) => {
        const val = parseFloat(item.price) || 0;
        return acc + val;
      }, 0).toFixed(2);
    },
    paymentTime() {
      let totalWeight = this.totalWeight;
      let totalPrice = this.totalPrice;
      let payment_method = ["credit card", "debit card", "cash"];
      console.log("The user chose to pay with " + payment_method)
    }

  },
  mounted() {
    this.updateDateTime();
    setInterval(this.updateDateTime, 1000);
    // Poll the light sensor every 2 seconds for real-time blackout detection
    setInterval(this.checkEnvironmentLight, 2000);
    this.initMinimap();
    this.fetchWeather();
    setInterval(this.fetchWeather, 600000);
  },
  methods: {
    logout() {
      localStorage.removeItem('username');
      localStorage.removeItem('loginMessage');
      this.$router.push('/login')
    },
    updateDateTime() {
      const now = new Date();
      this.date = now.toLocaleDateString('en-US', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' });
      this.time = now.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit', second: '2-digit', hour12: false });
    },
    async fetchWeather() {
      try {
        const response = await fetch('https://api.open-meteo.com/v1/forecast?latitude=37.9838&longitude=23.7275&current_weather=true');
        if (!response.ok) throw new Error("Weather service unavailable");
        const data = await response.json();
        if (data && data.current_weather) {
          this.weatherData = {
            temperature: data.current_weather.temperature,
            windspeed: data.current_weather.windspeed
          };
        }
      } catch (error) {
        console.error("Failed to fetch weather data:", error);
        this.errMessage = "Failed to load weather info";
      }
    },
    async assistance() {
      try {
        const url = `${API_URL}/assistance?cartID=${this.cart_id}`;
        await fetch(url);
        this.assistanceRequested = true;
        // Auto-clear confirmation after 30 seconds
        setTimeout(() => { this.assistanceRequested = false; }, 30000);
      } catch (error) {
        console.error("Failed to request assistance", error);
      }
    },
    async checkEnvironmentLight() {
      try {
        const url = `${API_URL}/measureLight?cartID=${this.cart_id}`;
        const response = await fetch(url);
        const data = await response.json();

        if (data.flashlight_needed) {
          this.isBlackout = true;
        } else {
          this.isBlackout = false;
        }
      } catch (error) {
        console.error("Failed to read light sensor", error);
      }
    },
    //weight checking from the scale vs weight of the database
    async triggerWeightCheck() {
      try {
        const url = `${API_URL}/measureWeight?cartID=${this.cart_id}&tag=${this.current_nfc_tag}`;
        const response = await fetch(url);
        const data = await response.json();

        if (data.status === "correct") {
          this.weightStatus = "Weight Match!";
          this.weight = data.actual_weight + " kg";
          this.price = data.price ? data.price.toFixed(2) + " €" : "0.00 €";
          this.scannedProduct = data.product_name;
          this.scannedProducts.unshift({
            name: data.product_name,
            weight: data.actual_weight + " kg",
            price: data.price ? data.price.toFixed(2) + " €" : "0.00 €",
            status: "Match"
          });
        } else {
          this.weightStatus = "Weight Mismatch!";
          this.weight = data.actual_weight + " kg (Expected: " + data.expected_weight + ")";
          this.price = data.price ? data.price.toFixed(2) + " €" : "0.00 €";
          this.scannedProducts.unshift({
            name: data.product_name || "Unknown",
            weight: data.actual_weight + " kg",
            price: data.price ? data.price.toFixed(2) + " €" : "0.00 €",
            status: "Mismatch"
          });
        }
      } catch (error) {
        console.error("Failed to fetch weight data:", error);
        this.weightStatus = "Error reaching Scale API";
      }
    },
    initMinimap() {
      const canvas = document.getElementById("canvas");
      if (!canvas) return;
      const ctx = canvas.getContext("2d");
      const img = document.getElementById("source");

      this.cartPath = [];
      this.currentPos = null;

      const draw = () => {
        ctx.clearRect(0, 0, canvas.width, canvas.height);

        // Draw background map
        if (img && img.complete) {
          ctx.drawImage(img, 0, 0, canvas.width, canvas.height);
        } else {
          // Fallback dark grid
          ctx.fillStyle = "#0f172a";
          ctx.fillRect(0, 0, canvas.width, canvas.height);
          ctx.strokeStyle = "rgba(255, 255, 255, 0.05)";
          ctx.lineWidth = 1;
          for (let x = 0; x < canvas.width; x += 30) {
            ctx.beginPath();
            ctx.moveTo(x, 0);
            ctx.lineTo(x, canvas.height);
            ctx.stroke();
          }
          for (let y = 0; y < canvas.height; y += 30) {
            ctx.beginPath();
            ctx.moveTo(0, y);
            ctx.lineTo(canvas.width, y);
            ctx.stroke();
          }
        }

        // Draw trace path line
        if (this.cartPath.length > 1) {
          ctx.beginPath();
          ctx.strokeStyle = "rgba(59, 130, 246, 0.7)"; // Vibrant light blue trace line
          ctx.lineWidth = 4;
          ctx.lineCap = "round";
          ctx.lineJoin = "round";
          ctx.setLineDash([6, 6]);

          for (let i = 0; i < this.cartPath.length; i++) {
            const pt = this.cartPath[i];
            const px = (pt.x / 10) * canvas.width;
            const py = (pt.y / 10) * canvas.height;

            if (i === 0) {
              ctx.moveTo(px, py);
            } else {
              ctx.lineTo(px, py);
            }
          }
          ctx.stroke();
          ctx.setLineDash([]);

          // Draw individual trace dots
          for (let i = 0; i < this.cartPath.length - 1; i++) {
            const pt = this.cartPath[i];
            const px = (pt.x / 10) * canvas.width;
            const py = (pt.y / 10) * canvas.height;
            ctx.beginPath();
            ctx.arc(px, py, 3.5, 0, Math.PI * 2);
            ctx.fillStyle = "rgba(96, 165, 250, 0.85)";
            ctx.fill();
          }
        }

        // Draw current position pulsing marker
        if (this.currentPos) {
          const px = (this.currentPos.x / 10) * canvas.width;
          const py = (this.currentPos.y / 10) * canvas.height;

          const time = Date.now() * 0.005;
          const pulseRadius = 12 + Math.sin(time) * 6;

          // Pulsing circle
          ctx.beginPath();
          ctx.arc(px, py, pulseRadius, 0, Math.PI * 2);
          ctx.strokeStyle = "rgba(59, 130, 246, 0.4)";
          ctx.lineWidth = 2.5;
          ctx.stroke();

          // Solid glowing blue core
          ctx.beginPath();
          ctx.arc(px, py, 7, 0, Math.PI * 2);
          ctx.fillStyle = "#3b82f6";
          ctx.shadowBlur = 12;
          ctx.shadowColor = "#3b82f6";
          ctx.fill();
          ctx.shadowBlur = 0;

          ctx.beginPath();
          ctx.arc(px, py, 2.5, 0, Math.PI * 2);
          ctx.fillStyle = "#ffffff";
          ctx.fill();

          // Label block
          ctx.fillStyle = "rgba(15, 23, 42, 0.9)";
          ctx.strokeStyle = "rgba(59, 130, 246, 0.6)";
          ctx.lineWidth = 1;
          const label = "CART_01";
          ctx.font = "bold 10px monospace";
          const textWidth = ctx.measureText(label).width;
          const rectW = textWidth + 12;
          const rectH = 18;
          const rectX = px - rectW / 2;
          const rectY = py - 28;

          ctx.beginPath();
          if (ctx.roundRect) {
            ctx.roundRect(rectX, rectY, rectW, rectH, 4);
          } else {
            ctx.rect(rectX, rectY, rectW, rectH);
          }
          ctx.fill();
          ctx.stroke();

          ctx.fillStyle = "#60a5fa";
          ctx.textAlign = "center";
          ctx.textBaseline = "middle";
          ctx.fillText(label, px, py - 19);
        }
      };

      this.ws = new WebSocket(`${WS_URL}/ws/minimap`);

      this.ws.onopen = () => {
        console.log("Minimap WebSocket connected successfully!");
      };

      this.ws.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data);
          if (data && typeof data.x_coordinate === "number" && typeof data.y_coordinate === "number") {
            const x = data.x_coordinate;
            const y = data.y_coordinate;

            this.cartPath.push({ x, y });
            if (this.cartPath.length > 50) {
              this.cartPath.shift();
            }

            this.currentPos = { x, y };
            draw();
          }
        } catch (err) {
          console.error("Error reading WebSocket position data:", err);
        }
      };

      this.ws.onerror = (err) => {
        console.error("WebSocket error:", err);
      };

      this.ws.onclose = () => {
        console.log("WebSocket disconnected, retrying in 5 seconds...");
        setTimeout(() => this.initMinimap(), 5000);
      };

      const animate = () => {
        if (this.ws && this.ws.readyState === WebSocket.OPEN) {
          draw();
          requestAnimationFrame(animate);
        }
      };

      if (img) {
        img.onload = () => draw();
      }
      draw();
      requestAnimationFrame(animate);
    }
  }
};
</script>