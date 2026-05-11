<template>
    <div class="app-background flex flex-col items-center justify-center min-h-screen text-white">

        <header
            class="relative backdrop-blur-md bg-black/30 dark:bg-pink-900/60 p-10 rounded-2xl shadow-2xl text-center space-y-6 border border-white/10 dark:border-gray-700 max-w-4xl w-full">

            <div class="absolute top-6 right-6 z-50 text-left">
                <button @click="isMenuOpen = !isMenuOpen"
                    class="relative group w-12 h-12 rounded-full bg-white/10 dark:bg-white hover:bg-white/20 dark:hover:bg-gray-700 shadow-xl ring-1 ring-white/10 dark:ring-gray-700 flex items-center justify-center transition-all duration-300 focus:outline-none">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16M4 16h16" />
                    </svg>
                    <div
                        class="flex flex-col justify-between w-20px h-20px transform transition-all duration-300 origin-center overflow-hidden">
                        <div class="bg-white dark:text-white h-2px w-7 transform transition-all duration-300 origin-left delay-150ms"
                            :class="{ 'rotate-42deg w-2/3': isMenuOpen }"></div>
                        <div class="bg-white dark:text-white h-2px w-7 rounded transform transition-all duration-300"
                            :class="{ 'translate-x-10': isMenuOpen }"></div>
                        <div class="bg-white dark:text-white h-2px w-7 transform transition-all duration-300 origin-left delay-150ms"
                            :class="{ '-rotate-42deg w-2/3': isMenuOpen }"></div>
                    </div>
                </button>

                <transition enter-active-class="transition duration-200 ease-out"
                    enter-from-class="opacity-0 -translate-y-2 scale-95"
                    enter-to-class="opacity-100 translate-y-0 scale-100"
                    leave-active-class="transition duration-150 ease-in"
                    leave-from-class="opacity-100 translate-y-0 scale-100"
                    leave-to-class="opacity-0 -translate-y-2 scale-95">
                    <div v-if="isMenuOpen"
                        class="absolute right-0 mt-3 w-64 origin-top-right rounded-xl bg-gray-900/95 backdrop-blur-xl shadow-2xl ring-1 ring-white/10 dark:ring-gray-700 py-2 focus:outline-none overflow-hidden">
                        <div
                            class="block px-4 py-3 text-sm text-gray-400 hover:bg-white/5 transition-colors cursor-default">
                            <span class="block text-xs uppercase tracking-wider opacity-50">Announcements</span>
                        </div>
                        <div
                            class="block px-4 py-3 text-sm text-gray-400 hover:bg-white/5 transition-colors cursor-default">
                            <span class="block text-xs uppercase tracking-wider opacity-50">Cart ID</span>
                            <span class="text-white font-mono">{{ cart_id }}</span>
                        </div>

                        <div
                            class="block px-4 py-3 text-sm text-gray-400 hover:bg-white/5 transition-colors cursor-default">
                            <span class="block text-xs uppercase tracking-wider opacity-50">Firmware</span>
                            <span class="text-white font-mono">{{ cart_version }}</span>
                        </div>

                        <button @click.prevent="triggerWeightCheck"
                            class="w-full text-left block px-4 py-3 text-sm text-gray-300 hover:bg-white/10 hover:text-white transition-colors">
                            Sensors Test
                        </button>

                        <a href="#"
                            class="block px-4 py-3 text-sm text-red-400 hover:bg-red-500/10 hover:text-red-300 transition-colors border-t border-white/5">
                            Logout
                        </a>

                    </div>
                </transition>
            </div>
            <h1 class="text-3xl font-bold tracking-wide mt-8">
                <p class="title dark:text-white">Hello {{ user }}, S.T.E.V.E Cart Management.</p>
                <p class="subtitle text-lg font-light opacity-80 dark:text-gray-300">{{ message }}</p>
            </h1>

            <!-- Flashlight Warning Banner -->
            <transition enter-active-class="transition duration-500 ease-out"
                enter-from-class="opacity-0 -translate-y-4" enter-to-class="opacity-100 translate-y-0"
                leave-active-class="transition duration-300 ease-in" leave-from-class="opacity-100 translate-y-0"
                leave-to-class="opacity-0 -translate-y-4">
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

            <transition enter-active-class="transition duration-500 ease-out"
                enter-from-class="opacity-0 -translate-y-4" enter-to-class="opacity-100 translate-y-0"
                leave-active-class="transition duration-300 ease-in" leave-from-class="opacity-100 translate-y-0"
                leave-to-class="opacity-0 -translate-y-4">
                <div v-if="assistanceCarts.length > 0"
                    class="w-full bg-red-500/90 text-white rounded-xl p-5 mt-6 shadow-red-500/50 shadow-2xl border border-red-400 animate-pulse">
                    <h2 class="text-xl font-bold mb-3">Assistance Requested!</h2>
                    <div class="space-y-2">
                        <div v-for="cart in assistanceCarts" :key="cart.cart_id"
                            class="flex justify-between items-center bg-white/10 rounded-lg px-4 py-2">
                            <span class="font-mono font-bold">{{ cart.cart_id }}</span>
                            <span class="text-sm opacity-80">{{ cart.gps }}</span>
                            <button @click="dismissAssistance(cart.cart_id)"
                                class="bg-white text-red-600 font-bold text-xs px-4 py-1.5 rounded-full hover:bg-red-100 transition whitespace-nowrap">
                                Dismiss
                            </button>
                        </div>
                    </div>
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
                class="bg-white/10 dark:bg-gray-800 rounded-xl p-6 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-full max-w-3xl mx-auto mt-8">
                <h3 class="text-2xl font-bold dark:text-white mb-4">Shelves Positions</h3>

                <!-- Form for Add/Update/Delete -->
                <div class="flex flex-wrap gap-3 mb-5">
                    <input type="text" v-model="shelveID" placeholder="Shelve ID"
                        class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500 flex-grow" />
                    <input type="text" v-model="xCoord" placeholder="X Coord"
                        class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500 w-24" />
                    <input type="text" v-model="yCoord" placeholder="Y Coord"
                        class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500 w-24" />
                    <input type="text" v-model="shelveDescription" placeholder="Description"
                        class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500 flex-grow" />
                </div>
                <div class="flex gap-3 mb-5 flex-wrap">
                    <button @click="addShelvePosition" class="bg-green-500/20 text-green-300 border border-green-500/30 rounded-xl px-4 py-2 hover:bg-green-500/30 transition font-semibold">+ Add</button>
                    <button @click="updateShelvePosition" class="bg-blue-500/20 text-blue-300 border border-blue-500/30 rounded-xl px-4 py-2 hover:bg-blue-500/30 transition font-semibold">✎ Update</button>
                    <button @click="deleteShelvePosition" class="bg-red-500/20 text-red-300 border border-red-500/30 rounded-xl px-4 py-2 hover:bg-red-500/30 transition font-semibold">✕ Delete</button>
                </div>

                <div class="overflow-x-auto">
                    <table class="w-full text-left text-gray-300 border-collapse">
                        <thead>
                            <tr class="border-b border-gray-600">
                                <th class="p-3">Shelve ID</th>
                                <th class="p-3">Products Available</th>
                                <th class="p-3">X Coord</th>
                                <th class="p-3">Y Coord</th>
                                <th class="p-3">Description</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="shelve in shelvePositions" :key="shelve.shelve_id"
                                class="border-b border-gray-700/50 hover:bg-white/5 transition cursor-pointer"
                                @click="shelveID = shelve.shelve_id; xCoord = shelve.x_coord; yCoord = shelve.y_coord; shelveDescription = shelve.description">
                                <td class="p-3 font-mono text-white">{{ shelve.shelve_id }}</td>
                                <td class="p-3">{{ shelve.product_count || '--' }}</td>
                                <td class="p-3 text-blue-400">{{ shelve.x_coord }}</td>
                                <td class="p-3 text-blue-400">{{ shelve.y_coord }}</td>
                                <td class="p-3">{{ shelve.description }}</td>
                            </tr>
                            <tr v-if="shelvePositions.length === 0">
                                <td colspan="5" class="p-6 text-center text-gray-500">No shelves added yet...</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>

            <div
                class="bg-white/10 dark:bg-gray-800 rounded-xl p-6 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-full max-w-3xl mx-auto mt-8">
                <h3 class="text-2xl font-bold dark:text-white mb-4">Online Carts</h3>
                <div class="overflow-x-auto">
                    <table class="w-full text-left text-gray-300 border-collapse">
                        <thead>
                            <tr class="border-b border-gray-600">
                                <th class="p-3">Cart ID</th>
                                <th class="p-3">Firmware</th>
                                <th class="p-3">Status</th>
                                <th class="p-3">Battery</th>
                                <th class="p-3">Location (GPS)</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="cart in onlineCarts" :key="cart.cart_id"
                                class="border-b border-gray-700/50 hover:bg-white/5 transition">
                                <td class="p-3 font-mono text-white">{{ cart.cart_id }}</td>
                                <td class="p-3">{{ cart.fw_version || 'v1.0' }}</td>
                                <td class="p-3">
                                    <span v-if="cart.is_active"
                                        class="px-2 py-1 bg-green-500/20 text-green-400 rounded-full text-xs font-bold">Active</span>
                                    <span v-else
                                        class="px-2 py-1 bg-red-500/20 text-red-400 rounded-full text-xs font-bold">Offline</span>
                                </td>
                                <td class="p-3 text-green-400">{{ cart.battery_level }}% ⚡</td>
                                <td class="p-3 text-blue-400">{{ cart.gps }}</td>
                            </tr>
                            <tr v-if="onlineCarts.length === 0">
                                <td colspan="5" class="p-6 text-center text-gray-500">No carts online...</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>

            <div
                class="bg-white/10 dark:bg-gray-800 rounded-xl p-6 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-full max-w-3xl mx-auto mt-8">
                <h3 class="text-2xl font-bold dark:text-white mb-4">User Statistics</h3>
                <div class="overflow-x-auto">
                    <table class="w-full text-left text-gray-300 border-collapse">
                        <thead>
                            <tr class="border-b border-gray-600">
                                <th class="p-3">User ID</th>
                                <th class="p-3">Loyalty Points</th>
                                <th class="p-3">Avg Total Cost</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="user in userStats" :key="user.id"
                                class="border-b border-gray-700/50 hover:bg-white/5 transition">
                                <td class="p-3 font-mono text-white">#{{ user.id }}</td>
                                <td class="p-3 text-yellow-400 font-bold"> {{ user.loyalty_points }}</td>
                                <td class="p-3 text-green-400 font-medium">{{ user.avg_cost.toFixed(2) }} €</td>
                            </tr>
                            <tr v-if="userStats.length === 0">
                                <td colspan="4" class="p-6 text-center text-gray-500">Loading stats...</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>

            <div class="flex flex-wrap gap-6 mt-10 text-center justify-center">

                <div
                    class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-fit">
                    <p class="text-xl font-semibold dark:text-white">Assistance Requests</p>
                    <ul class="bg-white/5 dark:bg-gray-900/30 rounded-xl p-4 mt-2 space-y-1 dark:text-gray-300 w-fit min-w-[180px]">
                        <li v-for="cart in assistanceCarts" :key="cart.cart_id"
                            class="flex items-center justify-between gap-3">
                            <span class="font-mono text-red-400 font-bold">{{ cart.cart_id }}</span>
                            <button @click="dismissAssistance(cart.cart_id)"
                                class="text-xs bg-red-500/20 text-red-300 border border-red-500/30 rounded-lg px-2 py-0.5 hover:bg-red-500/40 transition">
                                Dismiss
                            </button>
                        </li>
                        <li v-if="assistanceCarts.length === 0" class="text-gray-500 text-sm">No requests</li>
                    </ul>
                </div>

                <div
                    class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-fit text-left flex-grow">
                    <p class="text-xl font-semibold dark:text-white">Latest Sensor Scan</p>
                    <ul
                        class="bg-white/5 dark:bg-gray-900/30 rounded-xl p-4 mt-2 space-y-2 dark:text-gray-300 w-full min-w-[200px]">
                        <li class="flex justify-between gap-4">
                            <span class="opacity-70">Product:</span>
                            <span class="font-medium text-white">{{ scannedProduct || '--' }}</span>
                        </li>
                        <li class="flex justify-between gap-4">
                            <span class="opacity-70">Weight:</span>
                            <span class="font-medium text-white">{{ weight || '--' }}</span>
                        </li>
                        <li class="flex justify-between gap-4">
                            <span class="opacity-70">Price:</span>
                            <span class="font-medium text-white">{{ price || '--' }}</span>
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
                    class="bg-white/10 dark:bg-gray-800 rounded-xl p-4 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-fit">
                    <p class="text-xl font-semibold dark:text-white">Minimap(carts GPS status)</p>
                    <img src="../assets/minimap.png" alt="Supermarket Minimap" class="w-full h-full object-cover" />
                </div>
            </div>
        </header>
    </div>
</template>
<script>
export default {
    data() {
        return {
            weight: "", //kg 
            price: "", //€
            weightStatus: "",
            scannedProduct: "",
            // Added these for dynamic API calls!
            cart_id: "DISPLAY_CART_01",
            current_nfc_tag: "NFC_456",
            isBlackout: false,
            flashlightOn: false,
            onlineCarts: [],
            userStats: [],
            // Shelve form state
            shelvePositions: [],
            shelveID: "",
            xCoord: "",
            yCoord: "",
            shelveDescription: "",
            role: "admin" // TODO: derive from logged-in user
        };
    },
    mounted() {
        // Poll the light sensor every 2 seconds for real-time blackout detection
        setInterval(this.checkEnvironmentLight, 2000);
        this.fetchCarts();
        setInterval(this.fetchCarts, 5000);
        this.fetchUserStats();
        setInterval(this.fetchUserStats, 10000);
        this.fetchShelvePositions();
    },
    computed: {
        assistanceCarts() {
            return this.onlineCarts.filter(cart => cart.needs_assistance);
        }
    },
    methods: {
        logout() {
            this.$router.push('/login')
        },
        async dismissAssistance(cartID) {
            try {
                await fetch(`http://localhost:8089/dismissAssistance?cartID=${cartID}`);
                // Immediately reflect in UI without waiting for next poll
                const cart = this.onlineCarts.find(c => c.cart_id === cartID);
                if (cart) cart.needs_assistance = false;
            } catch (error) {
                console.error("Failed to dismiss assistance:", error);
            }
        },
        async fetchCarts() {
            try {
                const response = await fetch('http://localhost:8089/carts');
                const data = await response.json();
                if (data.status === 200) {
                    this.onlineCarts = data.data;
                }
            } catch (error) {
                console.error("Failed to fetch carts:", error);
            }
        },
        async fetchUserStats() {
            try {
                const response = await fetch('http://localhost:8089/userStats');
                if (!response.ok) throw new Error("Endpoint not available");
                const data = await response.json();
                if (data.status === 200) {
                    this.userStats = data.data;
                }
            } catch (error) {
                console.warn("API missing or failed, using mock data for User Stats.", error);
                this.userStats = [
                    { id: 101, email: "john@example.com", loyalty_points: 1250, avg_cost: 65.40 },
                    { id: 102, email: "maria@example.com", loyalty_points: 340, avg_cost: 112.20 },
                    { id: 103, email: "alex@example.com", loyalty_points: 50, avg_cost: 14.50 }
                ];
            }
        },
        async checkEnvironmentLight() {
            try {
                const url = `http://localhost:8089/measureLight?cartID=${this.cart_id}`;
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
                const url = `http://localhost:8089/measureWeight?cartID=${this.cart_id}&tag=${this.current_nfc_tag}`;
                const response = await fetch(url);
                const data = await response.json();

                if (data.status === "correct") {
                    this.weightStatus = "Weight Match!";
                    this.weight = data.actual_weight + " kg";
                    this.price = data.price ? data.price.toFixed(2) + " €" : "0.00 €";
                    this.scannedProduct = data.product_name;
                } else {
                    this.weightStatus = "Weight Mismatch!";
                    this.weight = data.actual_weight + " kg (Expected: " + data.expected_weight + ")";
                    this.price = data.price ? data.price.toFixed(2) + " €" : "0.00 €";
                }
            } catch (error) {
                console.error("Failed to fetch weight data:", error);
                this.weightStatus = "Error reaching Scale API";
            }
        },
        async fetchShelvePositions() {
            try {
                const response = await fetch('http://localhost:8089/shelves');
                const data = await response.json();
                if (data.status === 200) {
                    this.shelvePositions = data.data;
                }
            } catch (error) {
                console.warn("Failed to fetch shelve positions:", error);
            }
        },
        async addShelvePosition() {
            try {
                const response = await fetch('http://localhost:8089/addShelvePosition', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        shelve_id: this.shelveID,
                        x_coord: parseFloat(this.xCoord),
                        y_coord: parseFloat(this.yCoord),
                        description: this.shelveDescription,
                        role: this.role
                    })
                });
                const data = await response.json();
                if (data.status === 200) {
                    await this.fetchShelvePositions();
                    this.shelveID = ""; this.xCoord = ""; this.yCoord = ""; this.shelveDescription = "";
                }
            } catch (error) {
                console.error("Failed to add shelve position:", error);
            }
        },
        async deleteShelvePosition() {
            if (!this.shelveID) return alert("Select or enter a Shelve ID to delete.");
            try {
                const response = await fetch(`http://localhost:8089/deleteShelvePosition?shelveID=${this.shelveID}&role=${this.role}`, {
                    method: 'DELETE'
                });
                const data = await response.json();
                if (data.status === 200) {
                    await this.fetchShelvePositions();
                    this.shelveID = ""; this.xCoord = ""; this.yCoord = ""; this.shelveDescription = "";
                }
            } catch (error) {
                console.error("Failed to delete shelve position:", error);
            }
        },
        async updateShelvePosition() {
            if (!this.shelveID) return alert("Select or enter a Shelve ID to update.");
            try {
                const response = await fetch('http://localhost:8089/updateShelvePosition', {
                    method: 'PUT',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        shelve_id: this.shelveID,
                        x_coord: parseFloat(this.xCoord),
                        y_coord: parseFloat(this.yCoord),
                        description: this.shelveDescription,
                        role: this.role
                    })
                });
                const data = await response.json();
                if (data.status === 200) {
                    await this.fetchShelvePositions();
                }
            } catch (error) {
                console.error("Failed to update shelve position:", error);
            }
        }
    }
};
</script>