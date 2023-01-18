<script setup lang="ts">
import { ref } from 'vue'
import { StorageKey, Auth, errorMessageForFieldName, FieldValidationError } from "../../../assets/ts"
import { useRouter } from 'vue-router';
import * as F from "../../../../../backend/public/gomarvin.gen"
import InputComponent from '../../global/InputComponent.vue';
import ApiValidationFailedErrors from '../../global/ApiValidationFailedErrors.vue';

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
const failedValidationFields = ref<FieldValidationError[]>([])

/** Full flow for loggingg in a user */
async function PostUserLoginDetails() {
    isFetching.value = true
    apiResponseFailed.value = false

    const res = await F.UserEndpoints.RegisterUser(
        F.defaultClient, {
        username: username.value,
        password: password.value,
        email: email.value
    })

    const data = await res.json()

    /** If everything goes ok */
    if (data.status === 200) {
        isFetching.value = false
        StorageKey.Set(Auth.ACCESS_TOKEN_KEY(), data.data.token.access_token)
        router.push({ path: '/' })
    }
    else {
        if (data.message === "Payload validation failed!") {
            failedValidationFields.value = data.data
        }
        isFetching.value = false
        error_message.value = data.message
        apiResponseFailed.value = true
    }
}

</script>

<template>
    <div>
        <h4 class="fw-700">Sign Up</h4>
        <form @submit.prevent>

            <div class="grid gap-3">
                <InputComponent labelClass="input__1-label" inputClass="input__1" labelName="EMAIL" inputType="text"
                    inputId="email-login" :inputValue="email" @update="(email = $event)" />
                <InputComponent labelClass="input__1-label" inputClass="input__1" labelName="USERNAME" inputType="text"
                    inputId="username-login" :inputValue="username" @update="(username = $event)" />
                <InputComponent labelClass="input__1-label" inputClass="input__1" labelName="PASSWORD" inputType="text"
                    inputId="password-login" :inputValue="password" @update="(password = $event)" />
                <div><button class="button__1" @click="PostUserLoginDetails()">SIGN UP</button></div>
            </div>
        </form>

        <ApiValidationFailedErrors :errors="failedValidationFields" />
    </div>

</template>
