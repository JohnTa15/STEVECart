<template>
    <div class="app-background flex flex-col min-h-screen text-white p-4 md:p-6">

        <header
            class="relative backdrop-blur-md bg-black/30 dark:bg-pink-900/60 p-6 md:p-10 rounded-2xl shadow-2xl text-center space-y-6 border border-white/10 dark:border-gray-700 w-full flex-grow">

            <div class="absolute top-6 right-6 z-50 text-left">
                <button @click="logout"
                    class="bg-white/10 dark:bg-white/10 hover:bg-white/20 dark:hover:bg-gray-700 shadow-xl ring-1 ring-white/10 dark:ring-gray-700 px-6 py-2.5 rounded-xl transition-all duration-300 focus:outline-none text-sm font-semibold text-white">
                    Logout
                </button>
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
                        <h2 class="text-2xl font-bold">BLACKOUT DETECTED!</h2>
                        <p class="font-medium">Light levels dropped below safe limits. Do you need the flashlight?</p>
                    </div>
                    <button @click="flashlightOn = !flashlightOn"
                        class="bg-gray-900 text-white font-bold py-3 px-8 rounded-full hover:bg-black transition shadow-lg border border-gray-700 whitespace-nowrap">
                        {{ flashlightOn ? 'TURN OFF FLASHLIGHT' : 'TURN ON FLASHLIGHT' }}
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

            <div class="relative max-w-xl mx-auto my-6">
                <span class="absolute inset-y-0 left-0 flex items-center pl-4 text-slate-400">
                    <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                    </svg>
                </span>
                <input v-model="searchQuery"
                    class="w-full pl-12 pr-4 py-3 bg-white/5 dark:bg-gray-800/50 text-white placeholder:text-slate-500 rounded-2xl transition duration-300 border border-white/10 dark:border-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500"
                    placeholder="Search products by name, category, NFC tag, or description..." />
            </div>

            <div
                class="bg-white/10 dark:bg-gray-800 rounded-xl p-6 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-full mt-8">
                <div class="flex justify-between items-center select-none mb-4">
                    <h3 class="text-2xl font-bold dark:text-white">Products Management</h3>
                    <button type="button" @click="isProductsExpanded = !isProductsExpanded"
                        class="bg-white/10 hover:bg-white/20 px-3 py-1.5 rounded-lg text-sm transition-all duration-300 cursor-pointer">
                        {{ isProductsExpanded ? '▼ Collapse' : '▶ Expand' }}
                    </button>
                </div>

                <div v-if="isProductsExpanded">
                    <!-- Form for Add/Update/Delete -->
                    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3 mb-5">
                        <input type="text" v-model="productName" placeholder="Product Name"
                            class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500" />
                        <input type="text" v-model="productCategory" placeholder="Category"
                            class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500" />
                        <input type="text" v-model="nfcTag" placeholder="NFC Tag (ID)"
                            class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500" />
                        <input type="number" v-model="pcs" placeholder="Quantity (Pcs)"
                            class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500" />
                        <input type="text" v-model="weightVal" placeholder="Weight (kg)"
                            class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500" />
                        <input type="text" v-model="priceVal" placeholder="Price (€)"
                            class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500" />
                        <input type="text" v-model="productShelveID" placeholder="Shelve ID"
                            class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500" />
                        <input type="text" v-model="productDescription" placeholder="Description"
                            class="bg-white/5 border border-white/10 rounded-lg p-2 text-white placeholder:text-gray-500 flex-grow" />
                    </div>
                    <div class="flex gap-3 mb-5 flex-wrap">
                        <button @click="addProductPosition"
                            class="bg-green-500/20 text-green-300 border border-green-500/30 rounded-xl px-4 py-2 hover:bg-green-500/30 transition font-semibold">+
                            Add Product</button>
                        <button @click="updateProductPosition"
                            class="bg-blue-500/20 text-blue-300 border border-blue-500/30 rounded-xl px-4 py-2 hover:bg-blue-500/30 transition font-semibold">✎
                            Update Product</button>
                        <button @click="deleteProductPosition"
                            class="bg-red-500/20 text-red-300 border border-red-500/30 rounded-xl px-4 py-2 hover:bg-red-500/30 transition font-semibold">✕
                            Delete Product</button>
                    </div>

                    <div class="overflow-x-auto">
                        <table class="w-full text-left text-gray-300 border-collapse">
                            <thead>
                                <tr class="border-b border-gray-600">
                                    <th class="p-3">Product Name</th>
                                    <th class="p-3">Category</th>
                                    <th class="p-3">NFC Tag</th>
                                    <th class="p-3">Weight</th>
                                    <th class="p-3">Quantity</th>
                                    <th class="p-3">Price</th>
                                    <th class="p-3">Shelve ID</th>
                                    <th class="p-3">Description</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="product in filteredProducts" :key="product.ID"
                                    class="border-b border-gray-700/50 hover:bg-white/5 transition cursor-pointer"
                                    @click="productName = product.ProductName; productCategory = product.ProductCategory; nfcTag = product.NFCTag; productDescription = product.ProductDescription; weightVal = product.Weight; pcs = product.Pcs; priceVal = product.Price; productShelveID = product.shelve_id">
                                    <td class="p-3 font-semibold text-white">{{ product.ProductName }}</td>
                                    <td class="p-3">{{ product.ProductCategory }}</td>
                                    <td class="p-3 font-mono text-blue-400">{{ product.NFCTag }}</td>
                                    <td class="p-3">{{ product.Weight }} kg</td>
                                    <td class="p-3 text-yellow-400 font-bold">{{ product.Pcs }}</td>
                                    <td class="p-3 text-green-400 font-bold">{{ product.Price ? product.Price.toFixed(2)
                                        : '0.00' }} €</td>
                                    <td class="p-3 font-mono">{{ product.shelve_id }}</td>
                                    <td class="p-3 opacity-80">{{ product.ProductDescription }}</td>
                                </tr>
                                <tr v-if="filteredProducts.length === 0">
                                    <td colspan="8" class="p-6 text-center text-gray-500">No products added yet...</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div> <!-- End v-show="isProductsExpanded" -->
            </div>


            <div
                class="bg-white/10 dark:bg-gray-800 rounded-xl p-6 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-full mt-8">
                <div class="flex justify-between items-center select-none mb-4">
                    <h3 class="text-2xl font-bold dark:text-white">Shelves Positions</h3>
                    <button type="button" @click="isShelvesExpanded = !isShelvesExpanded"
                        class="bg-white/10 hover:bg-white/20 px-3 py-1.5 rounded-lg text-sm transition-all duration-300 cursor-pointer">
                        {{ isShelvesExpanded ? '▼ Collapse' : '▶ Expand' }}
                    </button>
                </div>

                <div v-if="isShelvesExpanded">
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
                        <button @click="addShelvePosition"
                            class="bg-green-500/20 text-green-300 border border-green-500/30 rounded-xl px-4 py-2 hover:bg-green-500/30 transition font-semibold">+
                            Add</button>
                        <button @click="updateShelvePosition"
                            class="bg-blue-500/20 text-blue-300 border border-blue-500/30 rounded-xl px-4 py-2 hover:bg-blue-500/30 transition font-semibold">✎
                            Update</button>
                        <button @click="deleteShelvePosition"
                            class="bg-red-500/20 text-red-300 border border-red-500/30 rounded-xl px-4 py-2 hover:bg-red-500/30 transition font-semibold">✕
                            Delete</button>
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
                </div> <!-- End v-show="isShelvesExpanded" -->
            </div>


            <div
                class="bg-white/10 dark:bg-gray-800 rounded-xl p-6 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-full mt-8">
                <div class="flex justify-between items-center select-none mb-4">
                    <h3 class="text-2xl font-bold dark:text-white">Users NEEDS FOR PROMOTING TO ADMIN...</h3>
                    <button type="button" @click="fetchUsers"
                        class="bg-white/10 hover:bg-white/20 px-3 py-1.5 rounded-lg text-sm transition-all duration-300 cursor-pointer flex items-center gap-1.5">
                        Refresh
                    </button>
                </div>
                <div class="overflow-x-auto">
                    <table class="w-full text-left text-gray-300 border-collapse">
                        <thead>
                            <tr class="border-b border-gray-600">
                                <th class="p-3">User ID</th>
                                <th class="p-3">Username</th>
                                <th class="p-3">Email</th>
                                <th class="p-3 text-center">Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="u in users" :key="u.id"
                                class="border-b border-gray-700/50 hover:bg-white/5 transition">
                                <td class="p-3 font-mono text-white">#{{ u.id }}</td>
                                <td class="p-3 font-semibold">{{ u.username }}</td>
                                <td class="p-3 font-mono text-blue-400">{{ u.email }}</td>
                                <td class="p-3 flex justify-center gap-3">
                                    <button @click="promoteUser(u.id)"
                                        class="bg-yellow-500/20 text-yellow-300 border border-yellow-500/30 rounded-xl px-3 py-1 hover:bg-yellow-500/40 transition text-xs font-semibold">
                                        Promote to Admin
                                    </button>
                                    <button @click="deleteUser(u.email)"
                                        class="bg-red-500/20 text-red-300 border border-red-500/30 rounded-xl px-3 py-1 hover:bg-red-500/40 transition text-xs font-semibold">
                                        Delete
                                    </button>
                                </td>
                            </tr>
                            <tr v-if="users.length === 0">
                                <td colspan="4" class="p-6 text-center text-gray-500">No users found...</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
            <div
                class="bg-white/10 dark:bg-gray-800 rounded-xl p-6 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-full mt-8">
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
                class="bg-white/10 dark:bg-gray-800 rounded-xl p-6 hover:bg-white/20 transition border border-transparent dark:border-gray-700 w-full mt-8">
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
                    <ul
                        class="bg-white/5 dark:bg-gray-900/30 rounded-xl p-4 mt-2 space-y-1 dark:text-gray-300 w-fit min-w-[180px]">
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
import { API_URL } from '../config.js';

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
            users: [], // Regular users directory
            weatherData: null,
            errMessage: "",
            // Shelve form state
            shelvePositions: [],
            shelveID: "",
            xCoord: "",
            yCoord: "",
            shelveDescription: "",
            role: localStorage.getItem('role') || "customer",
            user: localStorage.getItem('username') || "Admin",
            message: localStorage.getItem('loginMessage') || "Welcome to S.T.E.V.E Management",

            // Product form state
            products: [],
            productName: "",
            productCategory: "",
            nfcTag: "",
            productDescription: "",
            weightVal: "",
            pcs: "",
            priceVal: "",
            productShelveID: "",

            // Collapse/Expand state
            isProductsExpanded: true,
            isShelvesExpanded: true,

            // Search query state
            searchQuery: ""
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
        this.fetchProducts();
        this.fetchUsers();
        this.fetchWeather();
        setInterval(this.fetchWeather, 600000);
    },
    computed: {
        assistanceCarts() {
            return this.onlineCarts.filter(cart => cart.needs_assistance);
        },
        filteredProducts() {
            if (!this.searchQuery) {
                return this.products;
            }
            const query = this.searchQuery.toLowerCase();
            return this.products.filter(product => {
                const name = (product.ProductName || '').toLowerCase();
                const category = (product.ProductCategory || '').toLowerCase();
                const tag = (product.NFCTag || '').toLowerCase();
                const desc = (product.ProductDescription || '').toLowerCase();
                return name.includes(query) || 
                       category.includes(query) || 
                       tag.includes(query) || 
                       desc.includes(query);
            });
        }
    },
    methods: {
        logout() {
            this.$router.push('/login')
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
        async dismissAssistance(cartID) {
            try {
                await fetch(`${API_URL}/dismissAssistance?cartID=${cartID}`);
                // Immediately reflect in UI without waiting for next poll
                const cart = this.onlineCarts.find(c => c.cart_id === cartID);
                if (cart) cart.needs_assistance = false;
            } catch (error) {
                console.error("Failed to dismiss assistance:", error);
            }
        },
        async fetchCarts() {
            try {
                const response = await fetch(`${API_URL}/carts`);
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
                const response = await fetch(`${API_URL}/userStats`);
                if (!response.ok) throw new Error("Endpoint not available");
                const data = await response.json();
                if (data.status === 200) {
                    this.userStats = data.data;
                }
            } catch (error) {
                console.warn("API missing or failed, using mock data for User Stats.", error);
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
                const response = await fetch(`${API_URL}/shelves`);
                const data = await response.json();
                if (data.status === 200) {
                    this.shelvePositions = data.data;
                }
            } catch (error) {
                console.warn("Failed to fetch shelve positions:", error);
            }
        },
        async addShelvePosition() {
            if (!this.shelveID) {
                return alert("Shelve ID is required.");
            }
            try {
                const response = await fetch(`${API_URL}/addShelvePosition`, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        shelve_id: parseInt(this.shelveID, 10) || 0,
                        x_coord: parseFloat(this.xCoord) || 0.0,
                        y_coord: parseFloat(this.yCoord) || 0.0,
                        description: this.shelveDescription,
                        role: this.role
                    })
                });
                const data = await response.json();
                if (response.ok && data.status === 200) {
                    alert("Shelve position added successfully!");
                    await this.fetchShelvePositions();
                    this.shelveID = ""; this.xCoord = ""; this.yCoord = ""; this.shelveDescription = "";
                } else {
                    alert(data.error || "Failed to add shelve position.");
                }
            } catch (error) {
                console.error("Failed to add shelve position:", error);
                alert("Error communicating with backend");
            }
        },
        async deleteShelvePosition() {
            if (!this.shelveID) return alert("Select or enter a Shelve ID to delete.");
            if (!confirm(`Are you sure you want to delete shelve with ID ${this.shelveID}?`)) return;
            try {
                const response = await fetch(`${API_URL}/deleteShelvePosition?shelveID=${this.shelveID}&role=${this.role}`, {
                    method: 'DELETE'
                });
                const data = await response.json();
                if (response.ok && data.status === 200) {
                    alert("Shelve position deleted successfully!");
                    await this.fetchShelvePositions();
                    this.shelveID = ""; this.xCoord = ""; this.yCoord = ""; this.shelveDescription = "";
                } else {
                    alert(data.error || "Failed to delete shelve position.");
                }
            } catch (error) {
                console.error("Failed to delete shelve position:", error);
                alert("Error communicating with backend");
            }
        },
        async updateShelvePosition() {
            if (!this.shelveID) return alert("Select or enter a Shelve ID to update.");
            try {
                const response = await fetch(`${API_URL}/updateShelvePosition`, {
                    method: 'PUT',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        shelve_id: parseInt(this.shelveID, 10) || 0,
                        x_coord: parseFloat(this.xCoord) || 0.0,
                        y_coord: parseFloat(this.yCoord) || 0.0,
                        description: this.shelveDescription,
                        role: this.role
                    })
                });
                const data = await response.json();
                if (response.ok && data.status === 200) {
                    alert("Shelve position updated successfully!");
                    await this.fetchShelvePositions();
                } else {
                    alert(data.error || "Failed to update shelve position.");
                }
            } catch (error) {
                console.error("Failed to update shelve position:", error);
                alert("Error communicating with backend");
            }
        },
        // Product APIs
        async fetchProducts() {
            try {
                const response = await fetch(`${API_URL}/GetProducts`);
                const data = await response.json();
                if (data.status === 200) {
                    this.products = data.data;
                }
            } catch (error) {
                console.warn("Failed to fetch products:", error);
            }
        },
        async addProductPosition() {
            if (!this.productName || !this.nfcTag) {
                return alert("Product Name and NFC Tag are required.");
            }
            try {
                const response = await fetch(`${API_URL}/addProduct`, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        product_name: this.productName,
                        product_category: this.productCategory,
                        nfc_tag: this.nfcTag,
                        product_description: this.productDescription,
                        weight: parseFloat(this.weightVal) || 0.0,
                        pcs: parseInt(this.pcs) || 0,
                        price: parseFloat(this.priceVal) || 0.0,
                        shelve_id: parseInt(this.productShelveID, 10) || 0,
                        role: this.role
                    })
                });
                const data = await response.json();
                if (response.ok && data.status === 200) {
                    alert("Product added successfully!");
                    await this.fetchProducts();
                    this.clearProductForm();
                } else {
                    alert(data.error || "Failed to add product.");
                }
            } catch (error) {
                console.error("Failed to add product:", error);
                alert("Error communicating with backend");
            }
        },
        async updateProductPosition() {
            if (!this.nfcTag) return alert("NFC Tag is required to update a product.");
            try {
                const response = await fetch(`${API_URL}/updateProduct`, {
                    method: 'PUT',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        nfc_tag: this.nfcTag,
                        product_name: this.productName,
                        product_category: this.productCategory,
                        product_description: this.productDescription,
                        weight: parseFloat(this.weightVal) || 0.0,
                        pcs: parseInt(this.pcs) || 0,
                        price: parseFloat(this.priceVal) || 0.0,
                        shelve_id: parseInt(this.productShelveID, 10) || 0,
                        role: this.role
                    })
                });
                const data = await response.json();
                if (response.ok && data.status === 200) {
                    alert("Product updated successfully!");
                    await this.fetchProducts();
                } else {
                    alert(data.error || "Failed to update product.");
                }
            } catch (error) {
                console.error("Failed to update product:", error);
                alert("Error communicating with backend");
            }
        },
        async deleteProductPosition() {
            if (!this.nfcTag) return alert("Select or enter an NFC Tag to delete.");
            if (!confirm(`Are you sure you want to delete product with NFC Tag ${this.nfcTag}?`)) return;
            try {
                const response = await fetch(`${API_URL}/deleteProduct?nfcTag=${this.nfcTag}&role=${this.role}`, {
                    method: 'DELETE'
                });
                const data = await response.json();
                if (response.ok && data.status === 200) {
                    alert("Product deleted successfully!");
                    await this.fetchProducts();
                    this.clearProductForm();
                } else {
                    alert(data.error || "Failed to delete product.");
                }
            } catch (error) {
                console.error("Failed to delete product:", error);
                alert("Error communicating with backend");
            }
        },
        clearProductForm() {
            this.productName = "";
            this.productCategory = "";
            this.nfcTag = "";
            this.productDescription = "";
            this.weightVal = "";
            this.pcs = "";
            this.priceVal = "";
            this.productShelveID = "";
        },
        async fetchUsers() {
            try {
                const response = await fetch(`${API_URL}/users`);
                const data = await response.json();
                if (data.status === 200) {
                    this.users = data.data;
                }
            } catch (error) {
                console.error("Failed to fetch users:", error);
            }
        },

        async promoteUser(userID) {
            try {
                const response = await fetch(`${API_URL}/setAdmin`, {
                    method: 'PUT',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        user_id: String(userID),
                        role: this.role
                    })
                });
                const data = await response.json();
                if (response.ok) {
                    alert("User successfully promoted to Admin");
                    await this.fetchUsers();
                } else {
                    alert(data.error || "Failed to promote user");
                }
            } catch (error) {
                console.error("Failed to promote user:", error);
                alert("Error communicating with backend");
            }
        },
        async deleteUser(email) {
            if (!confirm(`Are you sure you want to delete user ${email}?`)) return;
            try {
                const response = await fetch(`${API_URL}/deleteUserByEmail?email=${encodeURIComponent(email)}`, {
                    method: 'DELETE'
                });
                const data = await response.json();
                if (response.ok) {
                    alert("User deleted successfully");
                    await this.fetchUsers();
                } else {
                    alert(data.error || "Failed to delete user");
                }
            } catch (error) {
                console.error("Failed to delete user:", error);
                alert("Error communicating with backend");
            }
        }
    }
};
</script>
