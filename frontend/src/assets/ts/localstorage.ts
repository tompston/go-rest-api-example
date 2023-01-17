
/**
 * Object that holds all of the util functions that are 
 * associated with localstorage manipulation
 */
export const StorageKey = {
    /** 
     * Set and store a value in localstorage.
     * Note that this overrides the value that was
     * stored for the key previously.
     */
    Set: function (key: string, value: string): void {
        localStorage.setItem(key, value)
    },
    /** 
     * If localstorage key exists, return true.
     * Else return false. 
     */
    Exists: function (key: string): boolean {
        return localStorage.getItem(key) ? true : false
    },
    /**
     * Delete Key from localstorage
     */
    Delete: function (key: string): void {
        localStorage.removeItem(key);
    },
    /**
     * 
     */
    Get: function (key: string): string | null {
        return localStorage.getItem(key)
    }
}
