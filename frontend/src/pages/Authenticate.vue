<script setup lang="ts">
import MainLayout from '../layout/MainLayout.vue';
import LoginForm from '../components/pages/Authenticate/LoginForm.vue';
import SignupForm from '../components/pages/Authenticate/SignupForm.vue';
</script>

<template>
    <MainLayout>
        <div class="grid grid-cols-2 gap-3">
            <!-- Sign Up grid -->
            <!-- <div>
                <h4 class="fw-700">Sign Up</h4>
                <div class="flex-center bg-red-600 text-white p-4 fw-600 border-rad-4 m-4">Not Implemented</div>
            </div> -->
            <SignupForm/>
            <!-- Log In Grid  -->
            <LoginForm />
        </div>
    </MainLayout>
</template>
