<script setup lang="ts">
import MainLayout from '../layout/MainLayout.vue';
import { GetAuthToken } from '../assets/ts/auth';
import { clientWithAuth } from "../assets/ts/client"
import { useRouter } from 'vue-router';
import { Ref, onMounted, ref } from 'vue';
import * as F from "../../../backend/public/gomarvin.gen"
import TransactionsTable from '../components/pages/Home/TransactionsTable.vue';
import UserDetails from '../components/pages/Home/UserDetails.vue';

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

function secondsToMinutes(s: number) {
    let min = s / 60
    return min.toFixed(2)
}

onMounted(() => {
    FetchUserTransactions()
    FetchUserDetails()
})

</script>

<template>
    <MainLayout>
        <UserDetails :API_user="API_user" />
        <TransactionsTable :transactions="API_transasctions" />
    </MainLayout>
</template>
