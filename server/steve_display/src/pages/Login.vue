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

    onMounted(() => {
        fetchUsersFromDB()
    })
}
</script>
<template>
    <div>
        <form name="Login" id="login-form" @submit.prevent="handleSubmitLogin">
            <div class="user-box">
                <label>Username:</label>
                <input type="text" placeholder="Admin" v-model="form.username" required>
                <label>Password:</label>
                <input type="password" placeholder="1234@" v-model="form.password" required>
                <button type="submit" @click="login_btn">Login</button>
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
        login_btn() {
            this.$router.push('/index')
        },
        signup_btn() {
            this.$router.push('/signup')
        },
        handleSubmitLogin() {
            if (this.form.username != null && this.form.password != null) {
                for (const user of users.value) {
                    if (this.form.username === user.username && this.form.password === user.password_hash) {
                        this.$router.push('/index')
                    } else {
                        this.errorMessage = "Invalid username or password!"
                    }
                }
            }
        }
    }
};
</script>