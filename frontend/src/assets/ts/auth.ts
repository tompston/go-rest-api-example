import { StorageKey } from "./localstorage"
import * as F from "../../../../backend/public/gomarvin.gen"

/**
 * Object that holds all of the util functions
 * that are used for authentication.
 */
export const Auth = {
    /**
     * Name of the localstorage key for the value that
     * holds the auth token
     */
    ACCESS_TOKEN_KEY: (): string => "access_token"
}


/** 
 * Flow that validates if the user is authenticated 
 *  - if localstorage auth token exists and is valid, return true
 *  - Else return false
*/
export async function isAuthenticated(): Promise<boolean> {
    let key = StorageKey.Get(Auth.ACCESS_TOKEN_KEY())
    if (key == null) return false

    const res = await F.AuthEndpoints.TokenIsValid(F.defaultClient, { access_token: key })

    if (res.status === 401) return false

    if (res.status === 200) {
        const data = await res.json()
        // console.log(data)
        if (data.data === true) return true
    }

    return false
}

export function GetAuthToken(): string {
    const key = StorageKey.Get(Auth.ACCESS_TOKEN_KEY())
    return key ? key : ""
}