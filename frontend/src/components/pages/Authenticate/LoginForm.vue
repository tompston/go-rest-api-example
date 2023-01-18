<script setup lang="ts">
import { ref } from 'vue'
import { placeholderUser } from "../../../../../backend/public/placeholder"
import * as F from "../../../../../backend/public/gomarvin.gen"
import { StorageKey, Auth } from "../../../assets/ts"
import { useRouter } from 'vue-router';
import InputComponent from '../../global/InputComponent.vue';

// router variable
const router = useRouter()
// Form input
const username = ref<string>(placeholderUser.username)
const password = ref<string>(placeholderUser.password)
// State Variables
const isFetching = ref<boolean>(false)
const apiResponseFailed = ref<boolean>(false)
const error_message = ref<string>("")

/** Full flow for loggingg in a user */
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
     *      Save the returned access token in localstorage 
     *      and move to the default homepage which needs 
     *      a correct authentication token to return
     *      the date from an authenticated API endpoint
     * If response is not ok
     *      return error message
     */
    const data = await res.json()

    if (data.status === 200) {
        isFetching.value = false
        StorageKey.Set(Auth.ACCESS_TOKEN_KEY(), data.data.token.access_token)
        router.push({ path: '/' })
    } else {
        isFetching.value = false
        apiResponseFailed.value = true
    }
}

</script>

<template>
    <div>
        <h4 class="fw-700">Log In</h4>
        <form @submit.prevent>
            <div class="grid gap-3">
                <InputComponent labelClass="input__1-label" inputClass="input__1" labelName="USERNAME" inputType="text"
                    inputId="username" :inputValue="username" @update="(username = $event)" />
                <InputComponent labelClass="input__1-label" inputClass="input__1" labelName="PASSWORD" inputType="text"
                    inputId="password" :inputValue="password" @update="(password = $event)" />
                <div><button class="button__1" @click="PostUserLoginDetails()">LOG IN</button></div>
            </div>
        </form>
    </div>
</template>