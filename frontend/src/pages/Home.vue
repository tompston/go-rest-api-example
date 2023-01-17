<script setup lang="ts">
import MainLayout from '../layout/MainLayout.vue';
import { GetAuthToken } from '../assets/ts/auth';
import { clientWithAuth } from "../assets/ts/client"
import { useRouter } from 'vue-router';
import { Ref, onMounted, ref } from 'vue';
import * as F from "../../../backend/public/gomarvin.gen"

/** User API Client with auth headers */
const client = clientWithAuth(GetAuthToken())

/** State variables */
const API_transasctions: Ref<Array<any>> = ref([])
const API_user: Ref<any> = ref({})

async function FetchUserTransactions() {
    const res = await F.GetTransactionsForUser(client)
    /** Redirect to login if not authenticated */
    if (res.status === 403) {
        useRouter().push({ path: '/auth' })
    }
    if (res.status === 200) {
        const data = await res.json()
        console.log(data)
        API_transasctions.value = data.data
    }
}

async function FetchUserDetails() {
    const res = await F.GetUserDetailsWithAuth(client)
    /** Redirect to login if not authenticated */
    if (res.status === 403) {
        useRouter().push({ path: '/auth' })
    }
    if (res.status === 200) {
        const data = await res.json()
        console.log(data)
        API_user.value = data.data
    }
}

onMounted(() => {
    FetchUserTransactions()
    FetchUserDetails()
})

</script>

<template>
    <MainLayout>

        <div class="p-5 bg-black text-white mb-4 fw-700">
            <h3 class="">{{ API_user.username }}</h3>
            <div>Created on {{ new Date(API_user.created_at).toDateString() }}</div>
            <div class="flex gap-2">
                <div>ID for user </div>
                <div class="font-mono opacity-60">{{ API_user.user_id }}</div>
            </div>
        </div>

        <div class="grid">
            <div class="grid grid-cols-3 fw-700 fs-10 bg-gray-200 p-2">
                <div>SENT TO</div>
                <div>AMOUNT</div>
                <div>DATE</div>
            </div>
            <div v-for="transaction in API_transasctions" v-bind:key="transaction"
                class="grid grid-cols-3 font-mono fs-9 border-b py-2 bg-gray-50">
                <div>{{ transaction.username }}</div>
                <div>{{ transaction.amount }}</div>
                <div>{{ new Date(transaction.created_at).toLocaleDateString('en-us') }}</div>
            </div>
        </div>
    </MainLayout>
</template>
