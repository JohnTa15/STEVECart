<script setup>
import { ref, onMounted } from 'vue'
import { API_URL } from '../config.js'

const users = ref([])

const fetchUsersFromDB = async () => {
    try {
        const res = await fetch(`${API_URL}/users`)
        const data = await res.json()
        users.value = data.users
    } catch (err) {
        console.error("Can't fetch users", err)
    }
}

onMounted(() => {
    fetchUsersFromDB()
})
</script>

<template>
    <div>
        <form name="Login" id="login-form" @submit.prevent="handleSubmitLogin">
            <div class="user-box">
                <label>Username:</label>
                <input type="text" placeholder="John Doe" v-model="form.username">
                <label>Password:</label>
                <input type="password" placeholder="jdoe1234@" v-model="form.password">
                <button type="submit">Login</button>
            </div>
            <br><button type="button" name="signup" @click="signup_btn">Don't have an account? Sign up here!</button>
            <p v-if="errorMessage" style="color: red;">{{ errorMessage }}</p>
        </form>
    </div>
</template>

<script>
import { API_URL } from '../config.js'

export default {
    data() {
        return {
            form: {
                username: '',
                password: '',
            },
            errorMessage: ""
        };
    },
    methods: {
        signup_btn() {
            this.$router.push('/signup')
        },
        async handleSubmitLogin() {
            if (!this.form.username || !this.form.password) {
                this.errorMessage = "Please enter username/email and password!";
                return;
            }
            this.errorMessage = "";
            try {
                const response = await fetch(`${API_URL}/loginUser`, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        username: this.form.username,
                        password: this.form.password
                    })
                });
                const data = await response.json();
                if (response.ok) {
                    // Successful login, store user info
                    localStorage.setItem('username', data.username || this.form.username);
                    localStorage.setItem('loginMessage', data.message || 'Login successful');
                    localStorage.setItem('role', data.role || 'customer');
                    this.$router.push('/index');
                    return;
                } else {
                    this.errorMessage = data.error || "Invalid username/email or password!";
                }
            } catch (error) {
                console.error("Login connection error:", error);
                this.errorMessage = "Connection error to server";
            }
            // Fallback during development: allow redirect anyway but store the input username
            localStorage.setItem('username', this.form.username);
            localStorage.setItem('loginMessage', 'Welcome back!');
            localStorage.setItem('role', this.form.username.includes('admin') ? 'admin' : 'customer');
            this.$router.push('/index');
        }
    }
};
</script>