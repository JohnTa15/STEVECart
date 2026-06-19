<template>
    <div>
        <form name="SignUp" id="login-form" @submit.prevent="handleSubmitSignUp">
            <div class="user-box">
                <label>Username</label>
                <input type="text" id="username" placeholder="John Doe" v-model="form.username" required>
                <label>Email</label>
                <input type="text" id="email" placeholder="jdoe@gmail.com" v-model="form.email" required>
                <label>Password</label>
                <input type="password" id="password" placeholder="1234@" v-model="form.password" required>
                <label>Type your password again</label>
                <input type="password" id="password_confirmation" placeholder="1234@"
                    v-model="form.password_confirmation" required>
                <button type="submit">Complete the registration</button>
            </div>
            <br><button type="button" name="login" @click="login_btn">Already have an account? Login here!</button>
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
                email: '',
                password: '',
                password_confirmation: ''
            },
            errorMessage: ''
        };
    },
    methods: {
        async handleSubmitSignUp() {
            if (!this.form.username || !this.form.email || !this.form.password) {
                this.errorMessage = "Please fill in all fields!";
                return;
            }
            if (this.form.password !== this.form.password_confirmation) {
                this.errorMessage = "Passwords do not match!";
                return;
            }
            this.errorMessage = "";
            
            try {
                const response = await fetch(`${API_URL}/registerUser`, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        username: this.form.username,
                        email: this.form.email,
                        password: this.form.password
                    })
                });
                const data = await response.json();
                if (response.ok) {
                    alert("Registration successful! Please log in.");
                    this.$router.push('/login');
                } else {
                    this.errorMessage = data.error || "Failed to register user";
                }
            } catch (error) {
                console.error("Signup error:", error);
                this.errorMessage = "Connection error to server";
            }
        },
        login_btn() {
            this.$router.push('/');
        }
    }
};
</script>