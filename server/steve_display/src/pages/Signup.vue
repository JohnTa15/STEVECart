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
            if (!this.form.username || !this.form.email || !this.form.password) {
                this.errorMessage = "Please fill in all fields!";
                return;
            }
            if (this.form.password !== this.form.password_confirmation) {
                this.errorMessage = "Passwords do not match!";
                return;
            }
            this.errorMessage = "";
            // Assuming successful registration for now, go to login
            alert("Registration successful! Please log in.");
            this.$router.push('/login');
        },
        login_btn() {
            this.$router.push('/');
        }
    }
};
</script>