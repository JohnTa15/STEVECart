<template>
    <div>
        <form name="SignUp" id="login-form" @submit.prevent="handleSubmitSignUp">
            <!-- ID HAS LOGIN FORM CAUSE THIS IS THE SAME WITH LOGIN -->
            <div class="user-box">
                <label>Username</label>
                <input type="text" id="username" placeholder="Admin" v-model="form.username" required>
                <label>Email</label>
                <input type="text" id="email" placeholder="admin@gmail.com" v-model="form.email" required>
                <label>Password</label>
                <input type="password" id="password" placeholder="1234@" v-model="form.password" required>
                <label>Type your password again</label>
                <input type="password" id="password_confirmation" placeholder="1234@" v-model="form.password_confirmation" required>
                <button type="submit" @click="signup_btn">Complete the registration</button>
            </div>
            <br><button type="button" name="login" @click="login_btn">Already have an account? Login here!</button>
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
                email: '',
                password: '',
                password_confirmation: ''
            },
            errorMessage: ''
        };
    },
    methods: {
        handleSubmitSignUp() {
            if (this.form.password !== this.form.password_confirmation) {
                this.errorMessage = "Passwords do not match!";
                return;
            }
            this.errorMessage = "";
        },
        login_btn() {
            this.$router.push('/');
        }
    }
};
</script>