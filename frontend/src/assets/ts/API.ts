/**
 * data.data interface that is returned when 
 * payload validation fails
 */
export interface FieldValidationErrors {
    errors: FieldValidationError[]
}

interface FieldValidationError {
    failed_field: string
    message: string
}
