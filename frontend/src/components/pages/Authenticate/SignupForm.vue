<script setup lang="ts">
import { ref } from 'vue'
import { StorageKey, Auth } from "../../../assets/ts"
import { useRouter } from 'vue-router';
import * as F from "../../../../../backend/public/gomarvin.gen"

// router variable
const router = useRouter()
// Form input
const username = ref<string>("")
const password = ref<string>("")
const email = ref<string>("")
// State Variables
const isFetching = ref<boolean>(false)
const apiResponseFailed = ref<boolean>(false)
const error_message = ref<string>("")

/** Full flow for loggingg in a user */
async function PostUserLoginDetails() {
    isFetching.value = true
    error_message.value = ""
    // error_message.value = ""

    const res = await F.UserEndpoints.RegisterUser(
        F.defaultClient, {
        username: username.value,
        password: password.value,
        email: email.value
    })

    /** 
     * If the response is ok:
     *      Save the returned access token in localstorage 
     *      and move to the default homepage which needs 
     *      a correct authentication token to return
     *      the date from an authenticated API endpoint
     * If response is not ok
     *      return error message
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
        <h4 class="fw-700">Sign Up</h4>
        <form @submit.prevent>

            <div class="grid gap-3">

                <!-- Email -->
                <div>
                    <label for="email-login" class="input__1-label">EMAIL</label>
                    <input required type="text" class="input__1" placeholder="username" id="email-login"
                        v-model="email">
                </div>
                <!-- Username -->
                <div>
                    <label for="username-login" class="input__1-label">USERNAME</label>
                    <input required type="text" class="input__1" placeholder="username" id="username-login"
                        v-model="username">
                </div>
                <!-- Password -->
                <div>
                    <label for="password-login" class="input__1-label">PASSWORD</label>
                    <input required type="password" class="input__1" placeholder="password" id="password-login"
                        v-model="password">
                </div>
                <div><button class="button__1" @click="PostUserLoginDetails()">LOG IN</button></div>
            </div>
        </form>
    </div>
</template>
