<script setup lang="ts">

defineProps<{
    // Container 
    containerClass?: string
    // Label
    labelClass?: string
    labelName: string
    // Input
    inputType: string
    inputClass?: string
    inputId: string
}>()

</script>

<template>
    <div :class="containerClass">
        <label :for="inputId" :class="labelClass">{{ labelName.toUpperCase() }}</label>
        <input :type="inputType" :class="inputClass" :id="inputId">
    </div>
</template>