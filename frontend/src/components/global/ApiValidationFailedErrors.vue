<script setup lang="ts">
import { FieldValidationError } from "../../assets/ts"
defineProps<{
    errors: FieldValidationError[]
}>();
</script>

<template>
    <div v-if="errors.length !== 0">
        <div class="bg-red-700 p-4 fs-9 text-white fw-600 border-rad-4 mt-4">
            <div v-for="field in errors" v-bind:key="field.failed_field">
                <div>{{ field.message }}!</div>
            </div>
        </div>
    </div>
</template>