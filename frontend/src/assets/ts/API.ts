/**
 * data.data interface that is returned when 
 * payload validation fails
 */
export interface FieldValidationErrors {
    errors: FieldValidationError[]
}

export interface FieldValidationError {
    failed_field: string
    message: string
}

export function isValidationErrors(value: any): value is FieldValidationErrors {
    return value.hasOwnProperty('errors');
}

/** 
 * Filter an array of objects to return the message
 * based on the failed_field key
 */
export function errorMessageForFieldName(fields: any[], field_key: string) {
    const filteredArray = fields.filter(item => item.failed_field === field_key);
    const string = filteredArray.map(item => item.message).join(', ');
    return string
}