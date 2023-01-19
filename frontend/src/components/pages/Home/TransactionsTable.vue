<script setup lang="ts">
defineProps<{
    transactions: any[]
}>()

</script>

<template>

    <div v-if="transactions.length === 0 || transactions == null">
        <div>
            <h4 class="fw-600">No transactions for the user found!</h4>
        </div>
    </div>

    <div v-if="transactions.length !== 0">
        <div class="grid">
            <div class="grid grid-cols-3 fw-700 fs-10 bg-gray-200 p-2">
                <div>SENT TO</div>
                <div>AMOUNT</div>
                <div>DATE</div>
            </div>
            <div v-for="transaction in transactions" v-bind:key="transaction"
                class="grid grid-cols-3 font-mono fs-9 border-b py-2 bg-gray-50">
                <div>{{ transaction.username }}</div>
                <div>{{ transaction.amount }}</div>
                <div>{{ new Date(transaction.created_at).toLocaleDateString('en-us') }}</div>
            </div>
        </div>
    </div>

</template>