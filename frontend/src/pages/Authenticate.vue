<script setup lang="ts">
import { ref } from 'vue'
import MainLayout from '../layout/MainLayout.vue';
import { placeholderUser } from "../../../backend/public/placeholder"
import * as F from "../../../backend/public/gomarvin.gen"
import { StorageKey, Auth } from "../assets/ts"
import { useRouter } from 'vue-router';

const router = useRouter()
// Form input
const username = ref<string>(placeholderUser.username)
const password = ref<string>(placeholderUser.password)
// State Variables
const isFetching = ref<boolean>(false)
const apiResponseFailed = ref<boolean>(false)
const error_message = ref<string>("")

/**
 * Full flow for loggingg in a user
 */
async function PostUserLoginDetails() {
    isFetching.value = true
    error_message.value = ""

    const res = await F.UserEndpoints.LoginUser(
        F.defaultClient, {
        username: username.value,
        password: password.value
    })

    /** 
     * If the response is ok:
     *  Save the returned access token in localstorage 
     *  and move to the default homepage which needs 
     *  a correct authentication token to return
     *  the date from an authenticated API endpoint
     * If response is not ok
     *  
     */
    if (res.status === 200) {
        isFetching.value = false
        const data = await res.json()
        StorageKey.Set(Auth.ACCESS_TOKEN_KEY(), data.data.token.access_token)
        router.push({ path: '/' })
    } else {
        isFetching.value = false
        apiResponseFailed.value = true
        // error_message.value = 
    }
}

</script>

<template>
    <MainLayout>
        <div class="grid grid-cols-2 gap-3">
            <!-- Sign Up grid -->
            <div>
                <h4 class="fw-700">Sign Up</h4>
                <div class="flex-center bg-red-600 text-white p-4 fw-600 border-rad-4 m-4">Not Implemented</div>
            </div>
            <!-- Log In Grid  -->
            <div>
                <h4 class="fw-700">Log In</h4>
                <div class="grid gap-3">
                    <div>
                        <label for="username" class="input__1-label">USERNAME</label>
                        <input type="text" class="input__1" placeholder="username" id="username" v-model="username">
                    </div>
                    <div>
                        <label for="password" class="input__1-label">PASSWORD</label>
                        <input type="password" class="input__1" placeholder="password" id="password" v-model="password">
                    </div>
                    <div><button class="button__1" @click="PostUserLoginDetails()">LOG IN</button></div>
                </div>
            </div>
        </div>
    </MainLayout>
</template>
