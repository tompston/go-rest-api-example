<script setup lang="ts">
import { ref } from 'vue'
import { placeholderUser } from "../../../../../backend/public/placeholder"
import * as F from "../../../../../backend/public/gomarvin.gen"
import { StorageKey, Auth } from "../../../assets/ts"
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
    <div>
        <h4 class="fw-700">Log In</h4>
        <form @submit.prevent>
            <div class="grid gap-3">
                <div>
                    <label for="username" class="input__1-label">USERNAME</label>
                    <input required type="text" class="input__1" placeholder="username" id="username"
                        v-model="username">
                </div>
                <div>
                    <label for="password" class="input__1-label">PASSWORD</label>
                    <input required type="password" class="input__1" placeholder="password" id="password"
                        v-model="password">
                </div>
                <div><button class="button__1" @click="PostUserLoginDetails()">LOG IN</button></div>
            </div>
        </form>
    </div>
</template>