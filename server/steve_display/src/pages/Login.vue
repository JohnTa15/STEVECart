<script setup>
import { ref, onMounted } from 'vue'

const users = ref([])

const fetchUsersFromDB = async () => {
    try {
        const res = await fetch('http://localhost:8089/users')
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
        handleSubmitLogin() {
            // if (!this.form.username || !this.form.password) {
            //     this.errorMessage = "Please enter username and password!"
            //     return
            // }

            // const matched = users.value.find(
            //     user => this.form.username === user.username && this.form.password === user.password_hash
            // )

            // if (matched) {
            //     this.$router.push('/index')
            // } else {
            //     this.errorMessage = "Invalid username or password!"
            // }
            this.$router.push('/index')
        }
    }
};
</script>